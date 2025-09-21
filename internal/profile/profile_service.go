package profile

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/chiquitav2/journalful/internal/db"
	"github.com/chiquitav2/journalful/pkg/profile/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	// ErrInvalidRequest is returned when the request is invalid
	ErrInvalidRequest = errors.New("invalid request")
	// ErrProfileNotFound is returned when the profile is not found
	ErrProfileNotFound = errors.New("profile not found")
	// ErrAuthorNotFound is returned when the author is not found
	ErrAuthorNotFound = fmt.Errorf("author not found")
)

type ProfileService struct {
	queries *db.Queries
}

func NewProfileService(conn *sql.DB) *ProfileService {
	return &ProfileService{
		queries: db.New(conn),
	}
}

func (p *ProfileService) GetProfile(ctx context.Context) (*profile.GetProfileResponse, error) {
	userID := ctx.Value("userID").(string)

	// Fetch the profile from the database
	profileData, err := p.queries.GetProfileByUserID(ctx, userID)
	if err != nil {
		return nil, ErrProfileNotFound
	}

	return &profile.GetProfileResponse{
		Profile: profileToGrpcProfile(&profileData),
	}, nil
}

func (p *ProfileService) CreateProfile(ctx context.Context, request *profile.CreateProfileRequest) (*profile.CreateProfileResponse, error) {
	userID := ctx.Value("userID").(string)

	result, err := p.queries.CreateProfile(ctx, db.CreateProfileParams{
		UserID:      userID,
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

	err := p.queries.UpdateProfile(ctx, db.UpdateProfileParams{
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
	err := p.queries.DeleteProfile(ctx, id)
	if err != nil {
		return nil, err
	}

	return &profile.DeleteProfileResponse{}, nil
}

func (p *ProfileService) ListProfiles(ctx context.Context) (*profile.ListProfilesResponse, error) {
	profiles, err := p.queries.ListProfiles(ctx)
	if err != nil {
		return nil, err
	}
	grpcProfiles := make([]*profile.Profile, len(profiles))
	for i, profile := range profiles {
		grpcProfiles[i] = profileToGrpcProfile(&profile)
	}
	return &profile.ListProfilesResponse{Profiles: grpcProfiles}, nil
}

func (p *ProfileService) mustEmbedUnimplementedProfileServiceServer() {
	//TODO implement me
	panic("implement me")
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
