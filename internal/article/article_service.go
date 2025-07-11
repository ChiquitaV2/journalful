package article

import (
	context "context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/chiquitav2/journalful/internal/db"
	"github.com/chiquitav2/journalful/internal/profile"
	"github.com/chiquitav2/journalful/pkg/articles/v1"
	v1 "github.com/chiquitav2/journalful/pkg/profile/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io"
	"log/slog"
	"net/http"
)

type ArticleGrpcService struct {
	article.UnimplementedArticlesServiceServer
	authorRepo  profile.Repository
	articleRepo *db.Queries
}

func NewArticleGrpcService(conn *sql.DB) *ArticleGrpcService {
	return &ArticleGrpcService{
		authorRepo:  profile.NewProfileRepository(conn),
		articleRepo: db.New(conn),
	}
}

func (s *ArticleGrpcService) GetArticle(ctx context.Context, request *article.GetArticleRequest) (*article.GetArticleResponse, error) {
	if request == nil || request.Id == 0 {
		return nil, fmt.Errorf("invalid request: %v", request)
	}
	// Fetch the article from the database
	articleData, err := s.articleRepo.GetArticle(ctx, request.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			slog.Warn("article not found", "id", request.Id)
			return nil, fmt.Errorf("article not found with ID: %d", request.Id)
		}
		slog.Error("failed to get article", "error", err)
		return nil, fmt.Errorf("failed to get article: %w", err)
	}
	// Fetch the authors for the article
	authors, err := s.articleRepo.ListArticleAuthorsByArticleID(ctx, request.Id)

	if err != nil {
		if err == sql.ErrNoRows {
			slog.Warn("no authors found for article", "id", request.Id)
			return nil, fmt.Errorf("no authors found for article with ID: %d", request.Id)
		}
		slog.Error("failed to get article authors", "error", err)
		return nil, fmt.Errorf("failed to get article authors: %w", err)
	}
	return &article.GetArticleResponse{
		Article: dbToGrpcArticle(articleData, authors),
	}, nil
}

func (s *ArticleGrpcService) GetArticleByDOI(ctx context.Context, request *article.GetArticleByDOIRequest) (*article.GetArticleByDOIResponse, error) {
	if request == nil || request.Doi == "" {
		return nil, fmt.Errorf("invalid request: %v", request)
	}
	// Fetch the article from the database
	articleData, err := s.articleRepo.GetArticleByDOI(ctx, request.Doi)
	if err != nil {
		if err == sql.ErrNoRows {
			slog.Warn("article not found", "doi", request.Doi)
			return nil, fmt.Errorf("article not found with DOI: %s", request.Doi)
		}
		slog.Error("failed to get article by DOI", "error", err)
		return nil, fmt.Errorf("failed to get article by DOI: %w", err)
	}
	// Fetch the authors for the article
	authors, err := s.articleRepo.ListArticleAuthorsByArticleID(ctx, articleData.ID)

	if err != nil {
		if err == sql.ErrNoRows {
			slog.Warn("no authors found for article", "id", articleData.ID)
			return nil, fmt.Errorf("no authors found for article with ID: %d", articleData.ID)
		}
		slog.Error("failed to get article authors", "error", err)
		return nil, fmt.Errorf("failed to get article authors: %w", err)
	}
	return &article.GetArticleByDOIResponse{
		Article: dbToGrpcArticle(articleData, authors),
	}, nil
}

func (s *ArticleGrpcService) ListArticles(ctx context.Context, request *article.ListArticlesRequest) (*article.ListArticlesResponse, error) {
	if request == nil {
		return nil, fmt.Errorf("invalid request: %v", request)
	}
	// Fetch the articles from the database
	articlesData, err := s.articleRepo.ListArticles(ctx)
	if err != nil {
		slog.Error("failed to list articles", "error", err)
		return nil, fmt.Errorf("failed to list articles: %w", err)
	}

	var articles []*article.Article
	for _, dbArticle := range articlesData {
		// Fetch the authors for each article
		authors, err := s.articleRepo.ListArticleAuthorsByArticleID(ctx, dbArticle.ID)
		if err != nil {
			slog.Error("failed to get article authors", "error", err)
			return nil, fmt.Errorf("failed to get article authors: %w", err)
		}
		articles = append(articles, dbToGrpcArticle(dbArticle, authors))
	}

	return &article.ListArticlesResponse{
		Articles: articles,
	}, nil
}

func (s *ArticleGrpcService) CreateArticle(ctx context.Context, request *article.CreateArticleRequest) (*article.CreateArticleResponse, error) {
	// Validate the request
	if request == nil || request.Doi == "" {
		return nil, fmt.Errorf("invalid request: DOI is required")
	}

	var (
		articleTitle           string
		articleAbstract        sql.NullString
		articlePublicationYear sql.NullInt32
		articleJournalName     sql.NullString
		normalizedAuthors      []string
	)

	// Attempt to fetch metadata from DOI if provided
	if request.Doi != "" {
		meta, err := s.fetchArticleMetadataFromDOI(request.Doi)
		if err != nil {
			slog.Error("failed to fetch article metadata from DOI", "doi", request.Doi, "error", err)
			// Continue with provided request data if DOI fetch fails
		} else if meta != nil {
			if len(meta.Title) > 0 {
				articleTitle = meta.Title[0]
			}
			articleAbstract = sql.NullString{String: meta.Abstract, Valid: meta.Abstract != ""}
			if len(meta.PublishedOnline.DateParts) > 0 && len(meta.PublishedOnline.DateParts[0]) > 0 {
				articlePublicationYear = sql.NullInt32{Int32: int32(meta.PublishedOnline.DateParts[0][0]), Valid: true}
			}
			if len(meta.ContainerTitle) > 0 {
				articleJournalName = sql.NullString{String: meta.ContainerTitle[0], Valid: meta.ContainerTitle[0] != ""}
			}

			for _, author := range meta.Author {
				fullName := fmt.Sprintf("%s %s", author.Given, author.Family)
				normalizedAuthors = append(normalizedAuthors, fullName)
			}
		}
	}

	// Override with request data if provided
	if request.Title != "" {
		articleTitle = request.Title
	}
	if request.Abstract != nil && *request.Abstract != "" {
		articleAbstract = sql.NullString{String: *request.Abstract, Valid: true}
	}
	if request.PublicationYear != nil && *request.PublicationYear != 0 {
		articlePublicationYear = sql.NullInt32{Int32: *request.PublicationYear, Valid: true}
	}
	if request.JournalName != nil && *request.JournalName != "" {
		articleJournalName = sql.NullString{String: *request.JournalName, Valid: true}
	}

	if len(request.Authors) > 0 {
		normalizedAuthors = make([]string, 0, len(request.Authors))
		for _, author := range request.Authors {
			if author != nil && author.Name != "" {
				normalizedAuthors = append(normalizedAuthors, author.Name)
			} else {
				slog.Warn("invalid author data in request", "author", author)
			}
		}
	}

	// Final validation after attempting to fetch/override
	if articleTitle == "" || len(normalizedAuthors) == 0 || !articlePublicationYear.Valid {
		return nil, fmt.Errorf("missing required article fields after DOI lookup and request override: title, authors, or publication year")
	}

	dbArticle, err := s.articleRepo.CreateArticle(
		ctx, db.CreateArticleParams{
			Title:           articleTitle,
			Doi:             request.Doi,
			Abstract:        articleAbstract,
			PublicationYear: articlePublicationYear,
			JournalName:     articleJournalName,
		})
	if err != nil {
		slog.Error("failed to create article", "error", err)
		return nil, fmt.Errorf("failed to create article: %w", err)
	}
	articleID, err := dbArticle.LastInsertId()
	if err != nil {
		slog.Error("failed to get last insert ID", "error", err)
		return nil, fmt.Errorf("failed to get last insert ID: %w", err)
	}
	slog.Info("article created", "id", articleID, "doi", request.Doi, "title", articleTitle)

	authors, err := s.authorRepo.FindOrCreateAuthors(ctx, normalizedAuthors)
	if err != nil {
		slog.Error("failed to find or create authors", "error", err)
		return nil, fmt.Errorf("failed to find or create authors: %w", err)
	}
	for i, author := range authors {
		if author == nil {
			slog.Warn("skipping nil author", "index", i)
			continue // Skip nil authors
		}
		// Insert each author into the database
		_, err = s.articleRepo.AddArticleAuthor(ctx, db.AddArticleAuthorParams{
			ArticleID: articleID,
			AuthorID:  author.Id,
			AuthorOrder: sql.NullInt32{
				Int32: int32(i + 1), // Use 1-based index for order
				Valid: true,
			},
		})
		if err != nil {
			slog.Error("failed to create article author", "error", err, "author", author.Name)
			return nil, fmt.Errorf("failed to create article author: %w", err)
		}
	}

	slog.Info("article created successfully", "id", articleID, "doi", request.Doi, "title", articleTitle)
	return &article.CreateArticleResponse{
		Id: articleID,
	}, nil
}

func (s *ArticleGrpcService) UpdateArticle(ctx context.Context, request *article.UpdateArticleRequest) (*article.UpdateArticleResponse, error) {
	if request == nil || request.Id == 0 {
		return nil, fmt.Errorf("invalid request: %v", request)
	}
	// Fetch the existing article from the database
	existingArticle, err := s.articleRepo.GetArticle(ctx, request.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			slog.Warn("article not found for update", "id", request.Id)
			return nil, fmt.Errorf("article not found with ID: %d", request.Id)
		}
		slog.Error("failed to get article for update", "error", err)
		return nil, fmt.Errorf("failed to get article for update: %w", err)
	}

	err = s.articleRepo.UpdateArticle(
		ctx,
		db.UpdateArticleParams{
			ID:              request.Id,
			Doi:             request.Doi,
			Title:           request.Title,
			Abstract:        sql.NullString{String: *request.Abstract, Valid: *request.Abstract != ""},
			PublicationYear: sql.NullInt32{Int32: *request.PublicationYear, Valid: *request.PublicationYear != 0},
			JournalName:     existingArticle.JournalName,
		})
	if err != nil {
		slog.Error("failed to update article", "error", err)
		return nil, fmt.Errorf("failed to update article: %w", err)
	}

	return &article.UpdateArticleResponse{}, nil
}

func (s *ArticleGrpcService) DeleteArticle(ctx context.Context, request *article.DeleteArticleRequest) (*article.DeleteArticleResponse, error) {
	if request == nil || request.Id == 0 {
		return nil, fmt.Errorf("invalid request: %v", request)
	}
	// Delete the article from the database
	err := s.articleRepo.DeleteArticle(ctx, request.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			slog.Warn("article not found for deletion", "id", request.Id)
			return nil, fmt.Errorf("article not found with ID: %d", request.Id)
		}
		slog.Error("failed to delete article", "error", err)
		return nil, fmt.Errorf("failed to delete article: %w", err)
	}
	slog.Info("article deleted successfully", "id", request.Id)
	return &article.DeleteArticleResponse{}, nil
}

func (s *ArticleGrpcService) mustEmbedUnimplementedArticlesServiceServer() {
	// This method is required to satisfy the interface, but it does not need to do anything.
	// It is used to ensure that the service implements all methods of the ArticlesServiceServer interface.
}

func dbToGrpcArticle(dbArticle db.Article, dbAuthors []db.ListArticleAuthorsByArticleIDRow) *article.Article {

	convertedAuthors := make([]*v1.Author, len(dbAuthors))
	for i, dbAuthor := range dbAuthors {
		convertedAuthors[i] = &v1.Author{
			Id:   dbAuthor.AuthorID,
			Name: dbAuthor.AuthorName,
		}
	}

	return &article.Article{
		Id:              dbArticle.ID,
		Doi:             dbArticle.Doi,
		Title:           dbArticle.Title,
		Authors:         convertedAuthors,
		Abstract:        &dbArticle.Abstract.String,
		PublicationYear: &dbArticle.PublicationYear.Int32,
		JournalName:     &dbArticle.JournalName.String,
		CreatedAt:       timestamppb.New(dbArticle.CreatedAt.Time),
		UpdatedAt:       timestamppb.New(dbArticle.UpdatedAt.Time),
	}
}

// CrossRef API response structs
type CrossRefResponse struct {
	Message CrossRefMessage `json:"message"`
}

type CrossRefMessage struct {
	DOI             string           `json:"DOI"`
	Title           []string         `json:"title"`
	Author          []CrossRefAuthor `json:"author"`
	PublishedOnline CrossRefDate     `json:"published-online"`
	Abstract        string           `json:"abstract"`
	ContainerTitle  []string         `json:"container-title"`
}

type CrossRefAuthor struct {
	Given  string `json:"given"`
	Family string `json:"family"`
}

type CrossRefDate struct {
	DateParts [][]int `json:"date-parts"`
}

func (s *ArticleGrpcService) fetchArticleMetadataFromDOI(doi string) (*CrossRefMessage, error) {
	url := fmt.Sprintf("https://api.crossref.org/v1/works/%s", doi)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make HTTP request to CrossRef API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("CrossRef API returned non-OK status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read CrossRef API response body: %w", err)
	}

	var crossRefResponse CrossRefResponse
	err = json.Unmarshal(body, &crossRefResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal CrossRef API response: %w", err)
	}

	return &crossRefResponse.Message, nil
}
