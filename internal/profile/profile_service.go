package profile

import (
	"context"
	"database/sql"
	"errors"
	"github.com/chiquitav2/journalful/internal/db"
	"github.com/chiquitav2/journalful/pkg/profile/v1"
)

var (
	// ErrInvalidRequest is returned when the request is invalid
	ErrInvalidRequest = errors.New("invalid request")
	// ErrProfileNotFound is returned when the profile is not found
	ErrProfileNotFound = errors.New("profile not found")
)

type ProfileService struct {
	repo Repository
}

func NewProfileService(conn *sql.DB) *ProfileService {
	return &ProfileService{
		repo: NewProfileRepository(conn),
	}
}

func (p *ProfileService) GetProfile(ctx context.Context, id int64) (*profile.GetProfileResponse, error) {

	// Fetch the profile from the database
	profileData, err := p.repo.GetProfile(ctx, id)
	if err != nil {
		return nil, ErrProfileNotFound
	}

	return &profile.GetProfileResponse{
		Profile: profileData,
	}, nil
}

func (p *ProfileService) CreateProfile(ctx context.Context, request *profile.CreateProfileRequest) (*profile.CreateProfileResponse, error) {

	result, err := p.repo.CreateProfile(ctx, db.CreateProfileParams{
		Name:        request.Name,
		Bio:         sql.NullString{String: *request.Bio, Valid: request.Bio != nil},
		Institution: sql.NullString{String: *request.Institution, Valid: request.Institution != nil},
	})
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &profile.CreateProfileResponse{Id: id}, nil
}

func (p *ProfileService) UpdateProfile(ctx context.Context, request *profile.UpdateProfileRequest) (*profile.UpdateProfileResponse, error) {
	if request == nil || request.Id == 0 {
		return nil, ErrInvalidRequest
	}

	err := p.repo.UpdateProfile(ctx, db.UpdateProfileParams{
		ID:          request.Id,
		Name:        request.Name,
		Bio:         sql.NullString{String: *request.Bio, Valid: request.Bio != nil},
		Institution: sql.NullString{String: *request.Institution, Valid: request.Institution != nil},
	})
	if err != nil {
		return nil, err
	}

	return &profile.UpdateProfileResponse{}, nil
}

func (p *ProfileService) DeleteProfile(ctx context.Context, id int64) (*profile.DeleteProfileResponse, error) {
	err := p.repo.DeleteProfile(ctx, id)
	if err != nil {
		return nil, err
	}

	return &profile.DeleteProfileResponse{}, nil
}

func (p *ProfileService) ListProfiles(ctx context.Context) (*profile.ListProfilesResponse, error) {
	profiles, err := p.repo.ListProfiles(ctx)
	if err != nil {
		return nil, err
	}

	return &profile.ListProfilesResponse{Profiles: profiles}, nil
}

func (p *ProfileService) mustEmbedUnimplementedProfileServiceServer() {
	//TODO implement me
	panic("implement me")
}
