package article

import (
	context "context"
	"database/sql"
	"fmt"

	"github.com/chiquitav2/journalful/internal/db"

	"log/slog"

	article "github.com/chiquitav2/journalful/pkg/articles/v1"
	v1 "github.com/chiquitav2/journalful/pkg/profile/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
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

func (s *ArticleSerivceImp) getArticleWithAuthors(ctx context.Context, fetchArticle func() (db.Article, error)) (*article.Article, error) {
	articleData, err := fetchArticle()
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "article not found")
		}
		slog.Error("failed to get article", "error", err)
		return nil, status.Error(codes.Internal, "failed to get article")
	}

	authors, err := s.queries.ListArticleAuthorsByArticleID(ctx, articleData.ID)
	if err != nil {
		slog.Error("failed to get article authors", "error", err)
		return nil, status.Error(codes.Internal, "failed to get article authors")
	}

	return dbToGrpcArticle(articleData, authors), nil
}

func (s *ArticleSerivceImp) GetArticle(ctx context.Context, id int64) (*article.GetArticleResponse, error) {
	grpcArticle, err := s.getArticleWithAuthors(ctx, func() (db.Article, error) {
		return s.queries.GetArticle(ctx, id)
	})
	if err != nil {
		return nil, err
	}

	return &article.GetArticleResponse{
		Article: grpcArticle,
	}, nil
}

func (s *ArticleSerivceImp) GetArticleByDOI(ctx context.Context, doi string) (*article.GetArticleByDOIResponse, error) {
	grpcArticle, err := s.getArticleWithAuthors(ctx, func() (db.Article, error) {
		return s.queries.GetArticleByDOI(ctx, doi)
	})
	if err != nil {
		return nil, err
	}

	return &article.GetArticleByDOIResponse{
		Article: grpcArticle,
	}, nil
}

func (s *ArticleSerivceImp) ListArticles(ctx context.Context, request *article.ListArticlesRequest) (*article.ListArticlesResponse, error) {
	if request == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	// Fetch the articles from the database
	articlesData, err := s.queries.ListArticlesWithAuthors(ctx)
	if err != nil {
		slog.Error("failed to list articles", "error", err)
		return nil, status.Error(codes.Internal, "failed to list articles")
	}

	// Group authors by article
	articlesMap := make(map[int64]*article.Article)
	for _, row := range articlesData {
		if _, ok := articlesMap[row.Article.ID]; !ok {
			articlesMap[row.Article.ID] = dbToGrpcArticle(row.Article, nil)
		}
		author := &v1.Author{
			Id:   row.Author.ID,
			Name: row.Author.Name,
		}
		articlesMap[row.Article.ID].Authors = append(articlesMap[row.Article.ID].Authors, author)
	}

	// Convert map to slice
	var articles []*article.Article
	for _, a := range articlesMap {
		articles = append(articles, a)
	}

	return &article.ListArticlesResponse{
		Articles: articles,
	}, nil
}

func (s *ArticleSerivceImp) CreateArticle(ctx context.Context, request *article.CreateArticleRequest) (*article.CreateArticleResponse, error) {
	// Validate the request
	if request == nil || request.Doi == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid request: DOI is required")
	}

	var meta *db.CreateArticleParams
	var authorNames []string
	var err error

	// Try to fetch metadata from external sources first
	meta, authorNames, err = s.metadataSvc.FetchAndPrepareArticle(request.Doi)
	if err != nil || meta == nil {
		// If external metadata fetch fails, use the provided request data
		slog.Info("external metadata fetch failed, using provided data", "doi", request.Doi)

		// Validate that we have at least basic required fields when metadata fetch fails
		if request.Title == "" {
			return nil, status.Error(codes.InvalidArgument, "title is required when metadata cannot be fetched")
		}

		// Create metadata from request
		meta = &db.CreateArticleParams{
			Doi:   request.Doi,
			Title: request.Title,
			Abstract: sql.NullString{
				String: "",
				Valid:  false,
			},
			PublicationYear: sql.NullInt32{
				Int32: 0,
				Valid: false,
			},
			JournalName: sql.NullString{
				String: "",
				Valid:  false,
			},
		}

		// Set optional fields if provided
		if request.Abstract != nil {
			meta.Abstract = sql.NullString{
				String: *request.Abstract,
				Valid:  true,
			}
		}
		if request.PublicationYear != nil {
			meta.PublicationYear = sql.NullInt32{
				Int32: *request.PublicationYear,
				Valid: true,
			}
		}
		if request.JournalName != nil {
			meta.JournalName = sql.NullString{
				String: *request.JournalName,
				Valid:  true,
			}
		}

		// Convert request authors to string slice
		authorNames = make([]string, len(request.Authors))
		for i, author := range request.Authors {
			authorNames[i] = author.Name
		}
	}

	// Create article
	dbArticle, err := s.queries.CreateArticle(ctx, *meta)
	if err != nil {
		slog.Error("failed to create article", "error", err)
		return nil, status.Error(codes.Internal, "failed to create article")
	}

	articleID, err := dbArticle.LastInsertId()
	if err != nil {
		slog.Error("failed to get last insert ID", "error", err)
		return nil, status.Error(codes.Internal, "failed to get last insert ID")
	}

	// Create article authors
	authors, err := s.FindOrCreateAuthors(ctx, authorNames)
	if err != nil {
		slog.Error("failed to find or create authors", "error", err)
		return nil, status.Error(codes.Internal, "failed to find or create authors")
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
			return nil, status.Error(codes.Internal, "failed to create article author")
		}
	}

	slog.Info("article created successfully", "id", articleID, "doi", request.Doi, "title", meta.Title)
	return &article.CreateArticleResponse{
		Id: articleID,
	}, nil
}

func (s *ArticleSerivceImp) UpdateArticle(ctx context.Context, request *article.UpdateArticleRequest) (*article.UpdateArticleResponse, error) {
	if request == nil || request.Id == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	// Fetch the existing article from the database
	_, err := s.queries.GetArticle(ctx, request.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			slog.Warn("article not found for update", "id", request.Id)
			return nil, status.Error(codes.NotFound, "article not found")
		}
		slog.Error("failed to get article for update", "error", err)
		return nil, status.Error(codes.Internal, "failed to get article for update")
	}

	err = s.queries.UpdateArticle(
		ctx,
		db.UpdateArticleParams{
			ID:              request.Id,
			Doi:             request.Doi,
			Title:           request.Title,
			Abstract:        sql.NullString{String: *request.Abstract, Valid: *request.Abstract != ""},
			PublicationYear: sql.NullInt32{Int32: *request.PublicationYear, Valid: *request.PublicationYear != 0},
			JournalName:     sql.NullString{String: *request.JournalName, Valid: *request.JournalName != ""},
		})
	if err != nil {
		slog.Error("failed to update article", "error", err)
		return nil, status.Error(codes.Internal, "failed to update article")
	}

	return &article.UpdateArticleResponse{}, nil
}

func (s *ArticleSerivceImp) DeleteArticle(ctx context.Context, request *article.DeleteArticleRequest) (*article.DeleteArticleResponse, error) {
	if request == nil || request.Id == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	// Delete the article from the database
	err := s.queries.DeleteArticle(ctx, request.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			slog.Warn("article not found for deletion", "id", request.Id)
			return nil, status.Error(codes.NotFound, "article not found")
		}
		slog.Error("failed to delete article", "error", err)
		return nil, status.Error(codes.Internal, "failed to delete article")
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
		return nil, status.Error(codes.InvalidArgument, "no author names provided")
	}

	var grpcAuthors []*v1.Author
	for _, name := range names {
		author, err := s.queries.GetAuthorByName(ctx, name)
		if err != nil && err != sql.ErrNoRows {
			slog.Error("failed to get author by name", "name", name, "error", err)
			return nil, status.Error(codes.Internal, fmt.Sprintf("failed to get author by name %s", name))
		}

		if err == sql.ErrNoRows {
			// If the author does not exist, create a new one
			authorRow, err := s.queries.CreateAuthor(ctx,
				db.CreateAuthorParams{
					Name: name,
				},
			)
			if err != nil {
				slog.Error("failed to create author", "name", name, "error", err)
				return nil, status.Error(codes.Internal, fmt.Sprintf("failed to create author with name %s", name))
			}

			id, err := authorRow.LastInsertId()
			if err != nil {
				slog.Error("failed to get last insert ID for author", "name", name, "error", err)
				return nil, status.Error(codes.Internal, fmt.Sprintf("failed to get last insert ID for author %s", name))
			}
			grpcAuthors = append(grpcAuthors, &v1.Author{Id: id, Name: name})
			continue
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
