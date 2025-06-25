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

type ProfileRepository struct {
	queries *db.Queries
}

// NewProfileRepository creates a new instance of profileRepository
func NewProfileRepository(queries *db.Queries) *ProfileRepository {
	return &ProfileRepository{
		queries: queries,
	}
}

// GetProfile retrieves a profile by its ID
func (r *ProfileRepository) GetProfile(ctx context.Context, id int64) (*profile.Profile, error) {
	profileID, err := r.queries.GetProfile(ctx, id)
	if err != nil {
		return nil, err
	}
	return profileToGrpcProfile(&profileID), nil
}

// GetAuthor retrieves an author by their ID
func (r *ProfileRepository) GetAuthor(ctx context.Context, id int64) (*profile.Author, error) {
	author, err := r.queries.GetAuthor(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("author not found with ID: %d", id)
		}
		return nil, fmt.Errorf("failed to get author: %w", err)
	}
	return authorToGrpcAuthor(&author), nil
}

func (r *ProfileRepository) GetAuthorByName(ctx context.Context, name string) (*profile.Author, error) {
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

func (r *ProfileRepository) FindOrCreateAuthors(ctx context.Context, names []string) ([]*profile.Author, error) {
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

func profileToGrpcProfile(p *db.Profile) *profile.Profile {
	if p == nil {
		return nil
	}
	return &profile.Profile{
		Id:        p.ID,
		Name:      p.Name,
		CreatedAt: timestamppb.New(p.CreatedAt.Time),
		UpdatedAt: timestamppb.New(p.UpdatedAt.Time),
	}
}

func authorToGrpcAuthor(a *db.Author) *profile.Author {
	if a == nil {
		return nil
	}
	return &profile.Author{
		Id:        a.ID,
		Name:      a.Name,
		CreatedAt: timestamppb.New(a.CreatedAt.Time),
		UpdatedAt: timestamppb.New(a.UpdatedAt.Time),
	}
}
