package profile

import (
	"context"
	"database/sql"
	"github.com/chiquitav2/journalful/internal/db"
	"github.com/chiquitav2/journalful/pkg/profile/v1"
)

type AuthorService struct {
	profile.UnimplementedAuthorServiceServer
	repo Repository
}

func NewAuthorService(conn *sql.DB) *AuthorService {
	return &AuthorService{
		repo: NewProfileRepository(conn),
	}
}

func (a *AuthorService) GetAuthor(ctx context.Context, request *profile.GetAuthorRequest) (*profile.GetAuthorResponse, error) {
	author, err := a.repo.GetAuthor(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	return &profile.GetAuthorResponse{Author: author}, nil
}

func (a *AuthorService) GetAuthorByProfileID(ctx context.Context, request *profile.GetAuthorByProfileIDRequest) (*profile.GetAuthorByProfileIDResponse, error) {
	author, err := a.repo.GetAuthorByProfileID(ctx, sql.NullInt64{Int64: request.ProfileId, Valid: true})
	if err != nil {
		return nil, err
	}
	return &profile.GetAuthorByProfileIDResponse{Author: author}, nil
}

func (a *AuthorService) ListAuthors(ctx context.Context, request *profile.ListAuthorsRequest) (*profile.ListAuthorsResponse, error) {
	authors, err := a.repo.ListAuthors(ctx)
	if err != nil {
		return nil, err
	}
	return &profile.ListAuthorsResponse{Authors: authors}, nil
}

func (a *AuthorService) CreateAuthor(ctx context.Context, request *profile.CreateAuthorRequest) (*profile.CreateAuthorResponse, error) {
	var result sql.Result
	var err error
	if request.ProfileId != nil {
		result, err = a.repo.CreateAuthorWithProfile(ctx, db.CreateAuthorWithProfileParams{
			Name:      request.Name,
			ProfileID: sql.NullInt64{Int64: *request.ProfileId, Valid: true},
		})
	} else {
		result, err = a.repo.CreateAuthor(ctx, request.Name)
	}
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &profile.CreateAuthorResponse{Id: id}, nil
}

func (a *AuthorService) UpdateAuthor(ctx context.Context, request *profile.UpdateAuthorRequest) (*profile.UpdateAuthorResponse, error) {
	err := a.repo.UpdateAuthor(ctx, db.UpdateAuthorParams{
		ID:        request.Id,
		Name:      request.Name,
		ProfileID: sql.NullInt64{Int64: *request.ProfileId, Valid: request.ProfileId != nil},
	})
	if err != nil {
		return nil, err
	}
	return &profile.UpdateAuthorResponse{}, nil
}

func (a *AuthorService) DeleteAuthor(ctx context.Context, request *profile.DeleteAuthorRequest) (*profile.DeleteAuthorResponse, error) {
	err := a.repo.DeleteAuthor(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	return &profile.DeleteAuthorResponse{}, nil
}

func (a *AuthorService) mustEmbedUnimplementedAuthorServiceServer() {
	//TODO implement me
	panic("implement me")
}
