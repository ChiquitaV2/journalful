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

func (h *GrpcHandler) GetLibrary(ctx context.Context, request *library.GetLibraryRequest) (*library.GetLibraryResponse, error) {
	return h.service.GetLibrary(ctx, request)
}

func (h *GrpcHandler) CreateLibrary(ctx context.Context, request *library.CreateLibraryRequest) (*library.CreateLibraryResponse, error) {
	return h.service.CreateLibrary(ctx, request)
}

func (h *GrpcHandler) UpdateLibrary(ctx context.Context, request *library.UpdateLibraryRequest) (*library.UpdateLibraryResponse, error) {
	return h.service.UpdateLibrary(ctx, request)
}

func (h *GrpcHandler) DeleteLibrary(ctx context.Context, request *library.DeleteLibraryRequest) (*library.DeleteLibraryResponse, error) {
	return h.service.DeleteLibrary(ctx, request)
}
