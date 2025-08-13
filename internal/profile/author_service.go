package profile

import (
	"context"
	"database/sql"
	"github.com/chiquitav2/journalful/internal/db"
	"github.com/chiquitav2/journalful/pkg/profile/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AuthorService struct {
	queries *db.Queries
}

func NewAuthorService(conn *sql.DB) *AuthorService {
	return &AuthorService{
		queries: db.New(conn),
	}
}

func (a *AuthorService) GetAuthor(ctx context.Context, id int64) (*profile.GetAuthorResponse, error) {
	author, err := a.queries.GetAuthor(ctx, id)
	if err != nil {
		return nil, err
	}
	return &profile.GetAuthorResponse{Author: authorToGrpcAuthor(&author)}, nil
}

func (a *AuthorService) GetAuthorByProfileID(ctx context.Context, id int64) (*profile.GetAuthorByProfileIDResponse, error) {
	author, err := a.queries.GetAuthorByProfileID(ctx, sql.NullInt64{Int64: id, Valid: true})
	if err != nil {
		return nil, err
	}
	return &profile.GetAuthorByProfileIDResponse{Author: authorToGrpcAuthor(&author)}, nil
}

func (a *AuthorService) ListAuthors(ctx context.Context) (*profile.ListAuthorsResponse, error) {
	authors, err := a.queries.ListAuthors(ctx)
	if err != nil {
		return nil, err
	}
	grpcAuthors := make([]*profile.Author, len(authors))
	for i, author := range authors {
		grpcAuthors[i] = authorToGrpcAuthor(&author)
	}
	return &profile.ListAuthorsResponse{Authors: grpcAuthors}, nil
}

func (a *AuthorService) CreateAuthor(ctx context.Context, request *profile.CreateAuthorRequest) (*profile.CreateAuthorResponse, error) {
	var result sql.Result
	var err error
	if request.ProfileId != nil {
		result, err = a.queries.CreateAuthor(ctx, db.CreateAuthorParams{
			Name:      request.Name,
			ProfileID: sql.NullInt64{Int64: *request.ProfileId, Valid: true},
		})
	} else {
		result, err = a.queries.CreateAuthor(ctx, db.CreateAuthorParams{
			Name:      request.Name,
			ProfileID: sql.NullInt64{Valid: false},
		})
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
	err := a.queries.UpdateAuthor(ctx, db.UpdateAuthorParams{
		ID:        request.Id,
		Name:      request.Name,
		ProfileID: sql.NullInt64{Int64: *request.ProfileId, Valid: request.ProfileId != nil},
	})
	if err != nil {
		return nil, err
	}
	return &profile.UpdateAuthorResponse{}, nil
}

func (a *AuthorService) DeleteAuthor(ctx context.Context, id int64) (*profile.DeleteAuthorResponse, error) {
	err := a.queries.DeleteAuthor(ctx, id)
	if err != nil {
		return nil, err
	}
	return &profile.DeleteAuthorResponse{}, nil
}

func (a *AuthorService) mustEmbedUnimplementedAuthorServiceServer() {
	//TODO implement me
	panic("implement me")
}

func authorToGrpcAuthor(a *db.Author) *profile.Author {
	if a == nil {
		return nil
	}
	return &profile.Author{
		Id:        a.ID,
		Name:      a.Name,
		ProfileId: a.ProfileID.Int64,
		CreatedAt: timestamppb.New(a.CreatedAt.Time),
		UpdatedAt: timestamppb.New(a.UpdatedAt.Time),
	}
}
