package profile

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/chiquitav2/journalful/internal/db"
	"github.com/chiquitav2/journalful/pkg/profile/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	// ErrAuthorNotFound is returned when the author is not found
	ErrAuthorNotFound = fmt.Errorf("author not found")
)

type Repository interface {
	// Profile methods
	GetProfile(ctx context.Context, id int64) (*profile.Profile, error)
	CreateProfile(ctx context.Context, arg db.CreateProfileParams) (sql.Result, error)
	UpdateProfile(ctx context.Context, arg db.UpdateProfileParams) error
	DeleteProfile(ctx context.Context, id int64) error
	ListProfiles(ctx context.Context) ([]*profile.Profile, error)

	// Author methods
	GetAuthor(ctx context.Context, id int64) (*profile.Author, error)
	GetAuthorByName(ctx context.Context, name string) (*profile.Author, error)
	GetAuthorByProfileID(ctx context.Context, profileID sql.NullInt64) (*profile.Author, error)
	FindOrCreateAuthors(ctx context.Context, names []string) ([]*profile.Author, error)
	CreateAuthor(ctx context.Context, name string) (sql.Result, error)
	CreateAuthorWithProfile(ctx context.Context, arg db.CreateAuthorWithProfileParams) (sql.Result, error)
	UpdateAuthor(ctx context.Context, arg db.UpdateAuthorParams) error
	DeleteAuthor(ctx context.Context, id int64) error
	ListAuthors(ctx context.Context) ([]*profile.Author, error)
	ListArticleAuthorsByAuthorID(ctx context.Context, authorID int64) ([]db.ListArticleAuthorsByAuthorIDRow, error)
}

type profileRepository struct {
	queries *db.Queries
}

// NewProfileRepository creates a new instance of profileRepository
func NewProfileRepository(conn *sql.DB) Repository {
	return &profileRepository{
		queries: db.New(conn),
	}
}

// GetProfile retrieves a profile by its ID
func (r *profileRepository) GetProfile(ctx context.Context, id int64) (*profile.Profile, error) {
	profileData, err := r.queries.GetProfile(ctx, id)
	if err != nil {
		return nil, err
	}
	return profileToGrpcProfile(&profileData), nil
}

func (r *profileRepository) CreateProfile(ctx context.Context, arg db.CreateProfileParams) (sql.Result, error) {
	return r.queries.CreateProfile(ctx, arg)
}

func (r *profileRepository) UpdateProfile(ctx context.Context, arg db.UpdateProfileParams) error {
	return r.queries.UpdateProfile(ctx, arg)
}

func (r *profileRepository) DeleteProfile(ctx context.Context, id int64) error {
	return r.queries.DeleteProfile(ctx, id)
}

func (r *profileRepository) ListProfiles(ctx context.Context) ([]*profile.Profile, error) {
	dbProfiles, err := r.queries.ListProfiles(ctx)
	if err != nil {
		return nil, err
	}
	var profiles []*profile.Profile
	for i := range dbProfiles {
		profiles = append(profiles, profileToGrpcProfile(&dbProfiles[i]))
	}
	return profiles, nil
}

// GetAuthor retrieves an author by their ID
func (r *profileRepository) GetAuthor(ctx context.Context, id int64) (*profile.Author, error) {
	author, err := r.queries.GetAuthor(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("author not found with ID: %d", id)
		}
		return nil, fmt.Errorf("failed to get author: %w", err)
	}
	return authorToGrpcAuthor(&author), nil
}

func (r *profileRepository) GetAuthorByName(ctx context.Context, name string) (*profile.Author, error) {
	if name == "" {
		return nil, fmt.Errorf("author name cannot be empty")
	}
	author, err := r.queries.GetAuthorByName(ctx, name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrAuthorNotFound
		}
		return nil, fmt.Errorf("failed to get author by name: %w", err)
	}
	return authorToGrpcAuthor(&author), nil
}

func (r *profileRepository) GetAuthorByProfileID(ctx context.Context, profileID sql.NullInt64) (*profile.Author, error) {
	author, err := r.queries.GetAuthorByProfileID(ctx, profileID)
	if err != nil {
		return nil, err
	}
	return authorToGrpcAuthor(&author), nil
}

func (r *profileRepository) FindOrCreateAuthors(ctx context.Context, names []string) ([]*profile.Author, error) {
	if len(names) == 0 {
		return nil, fmt.Errorf("no author names provided")
	}
	var grpcAuthors []*profile.Author
	// Insert or update authors in the database
	for _, name := range names {
		author, err := r.GetAuthorByName(ctx, name)
		if err == ErrAuthorNotFound {
			// If the author does not exist, create a new one
			authorRow, err := r.queries.CreateAuthor(ctx, name)
			if err != nil {
				return nil, fmt.Errorf("failed to create author with name %s: %w", name, err)
			}

			id, err := authorRow.LastInsertId()
			if err != nil {
				return nil, fmt.Errorf("failed to get last insert ID for author %s: %w", name, err)
			}
			// Convert the db.Author to profile.Author
			grpcAuthors = append(grpcAuthors, &profile.Author{
				Id:   id,
				Name: name,
			})
			continue
		} else if err != nil {
			return nil, fmt.Errorf("failed to get author by name %s: %w", name, err)
		}
		grpcAuthors = append(grpcAuthors, author)
	}

	return grpcAuthors, nil
}

func (r *profileRepository) CreateAuthor(ctx context.Context, name string) (sql.Result, error) {
	return r.queries.CreateAuthor(ctx, name)
}

func (r *profileRepository) CreateAuthorWithProfile(ctx context.Context, arg db.CreateAuthorWithProfileParams) (sql.Result, error) {
	return r.queries.CreateAuthorWithProfile(ctx, arg)
}

func (r *profileRepository) UpdateAuthor(ctx context.Context, arg db.UpdateAuthorParams) error {
	return r.queries.UpdateAuthor(ctx, arg)
}

func (r *profileRepository) DeleteAuthor(ctx context.Context, id int64) error {
	return r.queries.DeleteAuthor(ctx, id)
}

func (r *profileRepository) ListAuthors(ctx context.Context) ([]*profile.Author, error) {
	dbAuthors, err := r.queries.ListAuthors(ctx)
	if err != nil {
		return nil, err
	}
	var authors []*profile.Author
	for i := range dbAuthors {
		authors = append(authors, authorToGrpcAuthor(&dbAuthors[i]))
	}
	return authors, nil
}

func (r *profileRepository) ListArticleAuthorsByAuthorID(ctx context.Context, authorID int64) ([]db.ListArticleAuthorsByAuthorIDRow, error) {
	return r.queries.ListArticleAuthorsByAuthorID(ctx, authorID)
}

func profileToGrpcProfile(p *db.Profile) *profile.Profile {
	if p == nil {
		return nil
	}
	return &profile.Profile{
		Id:          p.ID,
		Name:        p.Name,
		Bio:         p.Bio.String,
		Institution: p.Institution.String,
		CreatedAt:   timestamppb.New(p.CreatedAt.Time),
		UpdatedAt:   timestamppb.New(p.UpdatedAt.Time),
	}
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
