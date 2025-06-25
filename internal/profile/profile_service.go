package profile

import (
	"context"
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
	profile.UnimplementedProfileServiceServer
	query *db.Queries
}

func (p ProfileService) GetProfile(ctx context.Context, request *profile.GetProfileRequest) (*profile.GetProfileResponse, error) {
	if request == nil || request.Id == 0 {
		return nil, ErrInvalidRequest
	}

	// Fetch the profile from the database
	profileData, err := p.query.GetProfile(ctx, request.Id)
	if err != nil {
		return nil, ErrProfileNotFound
	}

	return &profile.GetProfileResponse{
		Profile: profileToGrpcProfile(&profileData),
	}, nil
}

func (p ProfileService) mustEmbedUnimplementedProfileServiceServer() {
	//TODO implement me
	panic("implement me")
}
