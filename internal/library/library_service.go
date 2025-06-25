package library

import (
	"context"
	"github.com/chiquitav2/journalful/pkg/library/v1"
)

type LibraryService struct {
	library.UnimplementedLibraryServiceServer
}

func (l LibraryService) SaveArticleToLibrary(ctx context.Context, request *library.SaveArticleToLibraryRequest) (*library.SaveArticleToLibraryResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (l LibraryService) GetUserLibrary(ctx context.Context, request *library.GetUserLibraryRequest) (*library.GetUserLibraryResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (l LibraryService) mustEmbedUnimplementedLibraryServiceServer() {
	//TODO implement me
	panic("implement me")
}
