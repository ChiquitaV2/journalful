package profile

import (
	"context"
	"github.com/chiquitav2/journalful/pkg/profile/v1"
)

type AuthorService struct {
	profile.UnimplementedAuthorServiceServer
	profileRepository *ProfileRepository
}

func (a AuthorService) GetAuthor(ctx context.Context, request *profile.GetAuthorRequest) (*profile.GetAuthorResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthorService) GetAuthorByProfileID(ctx context.Context, request *profile.GetAuthorByProfileIDRequest) (*profile.GetAuthorByProfileIDResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthorService) ListAuthors(ctx context.Context, request *profile.ListAuthorsRequest) (*profile.ListAuthorsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthorService) CreateAuthor(ctx context.Context, request *profile.CreateAuthorRequest) (*profile.CreateAuthorResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthorService) UpdateAuthor(ctx context.Context, request *profile.UpdateAuthorRequest) (*profile.UpdateAuthorResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthorService) DeleteAuthor(ctx context.Context, request *profile.DeleteAuthorRequest) (*profile.DeleteAuthorResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthorService) mustEmbedUnimplementedAuthorServiceServer() {
	//TODO implement me
	panic("implement me")
}
