package library

import (
	"context"
	"database/sql"

	"github.com/chiquitav2/journalful/pkg/library/v1"
)

type GrpcHandler struct {
	library.UnimplementedLibraryServiceServer
	service LibraryServiceInterface
}

func NewLibraryGrpcHandler(conn *sql.DB) *GrpcHandler {
	return &GrpcHandler{
		service: NewLibraryService(conn),
	}
}

func (h *GrpcHandler) SaveArticleToLibrary(ctx context.Context, request *library.SaveArticleToLibraryRequest) (*library.SaveArticleToLibraryResponse, error) {
	return h.service.SaveArticleToLibrary(ctx, request)
}

func (h *GrpcHandler) GetUserLibrary(ctx context.Context, request *library.GetUserLibraryRequest) (*library.GetUserLibraryResponse, error) {
	return h.service.GetUserLibrary(ctx, request)
}
