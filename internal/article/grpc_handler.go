package article

import (
	"context"
	"database/sql"
	article "github.com/chiquitav2/journalful/pkg/articles/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ArticleGrpcHandler struct {
	article.UnimplementedArticlesServiceServer
	service ArticleService
}

func (h *ArticleGrpcHandler) GetArticle(ctx context.Context, request *article.GetArticleRequest) (*article.GetArticleResponse, error) {
	if request.Id == 0 {
		return nil, status.Error(codes.InvalidArgument, "article ID cannot be empty")
	}
	articleData, err := h.service.GetArticle(ctx, request.Id)
	if err != nil {
		return nil, err
	}

	return articleData, nil
}

func (h *ArticleGrpcHandler) GetArticleByDOI(ctx context.Context, request *article.GetArticleByDOIRequest) (*article.GetArticleByDOIResponse, error) {
	if request.Doi == "" {
		return nil, status.Error(codes.InvalidArgument, "DOI cannot be empty")
	}

	articleData, err := h.service.GetArticleByDOI(ctx, request.GetDoi())
	if err != nil {
		return nil, err
	}

	return articleData, nil
}

func (h *ArticleGrpcHandler) ListArticles(ctx context.Context, request *article.ListArticlesRequest) (*article.ListArticlesResponse, error) {

	articles, err := h.service.ListArticles(ctx, request)
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (h *ArticleGrpcHandler) CreateArticle(ctx context.Context, request *article.CreateArticleRequest) (*article.CreateArticleResponse, error) {
	if request.Doi == "" {
		return nil, status.Error(codes.InvalidArgument, "DOI cannot be empty")
	}
	articleData, err := h.service.CreateArticle(ctx, request)
	if err != nil {
		return nil, err
	}
	return articleData, nil
}

func (h *ArticleGrpcHandler) UpdateArticle(ctx context.Context, request *article.UpdateArticleRequest) (*article.UpdateArticleResponse, error) {
	if request.Id == 0 {
		return nil, status.Error(codes.InvalidArgument, "article ID cannot be empty")
	}
	articleData, err := h.service.UpdateArticle(ctx, request)
	if err != nil {
		return nil, err
	}
	return articleData, nil
}

func (h *ArticleGrpcHandler) DeleteArticle(ctx context.Context, request *article.DeleteArticleRequest) (*article.DeleteArticleResponse, error) {
	if request.Id == 0 {
		return nil, status.Error(codes.InvalidArgument, "article ID cannot be empty")
	}
	deletedArticle, err := h.service.DeleteArticle(ctx, request)
	if err != nil {
		return nil, err
	}
	return deletedArticle, nil
}

func NewArticleGrpcHandler(db *sql.DB) *ArticleGrpcHandler {
	return &ArticleGrpcHandler{
		service: NewArticleSerivce(db),
	}
}
