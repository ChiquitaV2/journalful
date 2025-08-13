package article

import (
	context "context"
	"database/sql"
	"fmt"
	"github.com/chiquitav2/journalful/internal/db"

	"github.com/chiquitav2/journalful/pkg/articles/v1"
	v1 "github.com/chiquitav2/journalful/pkg/profile/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
)

type ArticleService interface {
	GetArticle(ctx context.Context, id int64) (*article.GetArticleResponse, error)
	GetArticleByDOI(ctx context.Context, doi string) (*article.GetArticleByDOIResponse, error)
	ListArticles(ctx context.Context, request *article.ListArticlesRequest) (*article.ListArticlesResponse, error)
	CreateArticle(ctx context.Context, request *article.CreateArticleRequest) (*article.CreateArticleResponse, error)
	UpdateArticle(ctx context.Context, request *article.UpdateArticleRequest) (*article.UpdateArticleResponse, error)
	DeleteArticle(ctx context.Context, request *article.DeleteArticleRequest) (*article.DeleteArticleResponse, error)
}

type ArticleSerivceImp struct {
	queries     *db.Queries
	metadataSvc *MetadataService
}

func NewArticleSerivce(conn *sql.DB) ArticleService {
	return &ArticleSerivceImp{
		queries:     db.New(conn),
		metadataSvc: NewMetadataService(),
	}
}

func (s *ArticleSerivceImp) GetArticle(ctx context.Context, id int64) (*article.GetArticleResponse, error) {
	// Fetch the article from the database
	articleData, err := s.queries.GetArticle(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			slog.Warn("article not found", "id", id)
			return nil, fmt.Errorf("article not found with ID: %d", id)
		}
		slog.Error("failed to get article", "error", err)
		return nil, fmt.Errorf("failed to get article: %w", err)
	}
	// Fetch the authors for the article
	authors, err := s.queries.ListArticleAuthorsByArticleID(ctx, id)

	if err != nil {
		if err == sql.ErrNoRows {
			slog.Warn("no authors found for article", "id", id)
			return nil, fmt.Errorf("no authors found for article with ID: %d", id)
		}
		slog.Error("failed to get article authors", "error", err)
		return nil, fmt.Errorf("failed to get article authors: %w", err)
	}
	return &article.GetArticleResponse{
		Article: dbToGrpcArticle(articleData, authors),
	}, nil
}

func (s *ArticleSerivceImp) GetArticleByDOI(ctx context.Context, doi string) (*article.GetArticleByDOIResponse, error) {
	// Fetch the article from the database
	articleData, err := s.queries.GetArticleByDOI(ctx, doi)
	if err != nil {
		if err == sql.ErrNoRows {
			slog.Warn("article not found", "doi", doi)
			return nil, fmt.Errorf("article not found with DOI: %s", doi)
		}
		slog.Error("failed to get article by DOI", "error", err)
		return nil, fmt.Errorf("failed to get article by DOI: %w", err)
	}
	// Fetch the authors for the article
	authors, err := s.queries.ListArticleAuthorsByArticleID(ctx, articleData.ID)

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

func (s *ArticleSerivceImp) ListArticles(ctx context.Context, request *article.ListArticlesRequest) (*article.ListArticlesResponse, error) {
	if request == nil {
		return nil, fmt.Errorf("invalid request: %v", request)
	}
	// Fetch the articles from the database
	articlesData, err := s.queries.ListArticles(ctx)
	if err != nil {
		slog.Error("failed to list articles", "error", err)
		return nil, fmt.Errorf("failed to list articles: %w", err)
	}

	var articles []*article.Article
	for _, dbArticle := range articlesData {
		// Fetch the authors for each article
		authors, err := s.queries.ListArticleAuthorsByArticleID(ctx, dbArticle.ID)
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

func (s *ArticleSerivceImp) CreateArticle(ctx context.Context, request *article.CreateArticleRequest) (*article.CreateArticleResponse, error) {
	// Validate the request
	if request == nil || request.Doi == "" {
		return nil, fmt.Errorf("invalid request: DOI is required")
	}
	meta, normalizedAuthors, err := s.metadataSvc.FetchAndPrepareArticle(request.Doi)
	if err != nil {
		slog.Error("failed to fetch article metadata from DOI", "doi", request.Doi, "error", err)
		return nil, fmt.Errorf("failed to fetch article metadata from DOI: %w", err)
	}
	if meta == nil {
		slog.Warn("no metadata found for article", "doi", request.Doi)
		return nil, fmt.Errorf("no metadata found for article with DOI: %s", request.Doi)
	}

	// Create article
	dbArticle, err := s.queries.CreateArticle(ctx, *meta)
	if err != nil {
		slog.Error("failed to create article", "error", err)
		return nil, fmt.Errorf("failed to create article: %w", err)
	}

	articleID, err := dbArticle.LastInsertId()
	if err != nil {
		slog.Error("failed to get last insert ID", "error", err)
		return nil, fmt.Errorf("failed to get last insert ID: %w", err)
	}

	// Create article authors
	authors, err := s.FindOrCreateAuthors(ctx, normalizedAuthors)
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
		_, err = s.queries.AddArticleAuthor(ctx, db.AddArticleAuthorParams{
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

	slog.Info("article created successfully", "id", articleID, "doi", request.Doi, "title", meta.Title)
	return &article.CreateArticleResponse{
		Id: articleID,
	}, nil
}

func (s *ArticleSerivceImp) UpdateArticle(ctx context.Context, request *article.UpdateArticleRequest) (*article.UpdateArticleResponse, error) {
	if request == nil || request.Id == 0 {
		return nil, fmt.Errorf("invalid request: %v", request)
	}
	// Fetch the existing article from the database
	existingArticle, err := s.queries.GetArticle(ctx, request.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			slog.Warn("article not found for update", "id", request.Id)
			return nil, fmt.Errorf("article not found with ID: %d", request.Id)
		}
		slog.Error("failed to get article for update", "error", err)
		return nil, fmt.Errorf("failed to get article for update: %w", err)
	}

	err = s.queries.UpdateArticle(
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

func (s *ArticleSerivceImp) DeleteArticle(ctx context.Context, request *article.DeleteArticleRequest) (*article.DeleteArticleResponse, error) {
	if request == nil || request.Id == 0 {
		return nil, fmt.Errorf("invalid request: %v", request)
	}
	// Delete the article from the database
	err := s.queries.DeleteArticle(ctx, request.Id)
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

func (s *ArticleSerivceImp) mustEmbedUnimplementedArticlesServiceServer() {
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

func (s *ArticleSerivceImp) FindOrCreateAuthors(ctx context.Context, names []string) ([]*v1.Author, error) {
	if len(names) == 0 {
		return nil, fmt.Errorf("no author names provided")
	}
	var grpcAuthors []*v1.Author
	// Insert or update authors in the database
	for _, name := range names {
		author, err := s.queries.GetAuthorByName(ctx, name)
		if err != nil && err != sql.ErrNoRows {
			return nil, fmt.Errorf("failed to get author by name %s: %w", name, err)
		}

		if err == sql.ErrNoRows {
			// If the author does not exist, create a new one
			authorRow, err := s.queries.CreateAuthor(ctx,
				db.CreateAuthorParams{
					Name: name,
				},
			)
			if err != nil {
				return nil, fmt.Errorf("failed to create author with name %s: %w", name, err)
			}

			id, err := authorRow.LastInsertId()
			if err != nil {
				return nil, fmt.Errorf("failed to get last insert ID for author %s: %w", name, err)
			}
			// Convert the db.Author to profile.Author
			grpcAuthors = append(grpcAuthors, &v1.Author{
				Id:   id,
				Name: name,
			})
			continue
		} else if err != nil {
			return nil, fmt.Errorf("failed to get author by name %s: %w", name, err)
		}
		grpcAuthors = append(grpcAuthors, &v1.Author{
			Id:        author.ID,
			Name:      author.Name,
			ProfileId: author.ProfileID.Int64,
			CreatedAt: timestamppb.New(author.CreatedAt.Time),
			UpdatedAt: timestamppb.New(author.UpdatedAt.Time),
		})
	}

	return grpcAuthors, nil
}
