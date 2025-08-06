package profile

import (
	"context"
	"database/sql"

	"github.com/chiquitav2/journalful/pkg/profile/v1"
)

type ProfileGrpcHandler struct {
	profile.UnimplementedAuthorServiceServer
	profile.UnimplementedProfileServiceServer
	profileService *ProfileService
	authorService  *AuthorService
}

func (p ProfileGrpcHandler) GetProfile(ctx context.Context, request *profile.GetProfileRequest) (*profile.GetProfileResponse, error) {
	if request == nil || request.Id == 0 {
		return nil, ErrInvalidRequest
	}
	profileData, err := p.profileService.GetProfile(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	return profileData, nil
}

func (p ProfileGrpcHandler) ListProfiles(ctx context.Context, request *profile.ListProfilesRequest) (*profile.ListProfilesResponse, error) {
	return p.profileService.ListProfiles(ctx)
}

func (p ProfileGrpcHandler) CreateProfile(ctx context.Context, request *profile.CreateProfileRequest) (*profile.CreateProfileResponse, error) {
	if request == nil || request.Name == "" {
		return nil, ErrInvalidRequest
	}

	result, err := p.profileService.CreateProfile(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p ProfileGrpcHandler) UpdateProfile(ctx context.Context, request *profile.UpdateProfileRequest) (*profile.UpdateProfileResponse, error) {
	if request == nil || request.Id == 0 {
		return nil, ErrInvalidRequest
	}

	result, err := p.profileService.UpdateProfile(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p ProfileGrpcHandler) DeleteProfile(ctx context.Context, request *profile.DeleteProfileRequest) (*profile.DeleteProfileResponse, error) {
	if request == nil || request.Id == 0 {
		return nil, ErrInvalidRequest
	}

	result, err := p.profileService.DeleteProfile(ctx, request.Id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p ProfileGrpcHandler) mustEmbedUnimplementedProfileServiceServer() {
	//TODO implement me
	panic("implement me")
}

func (p ProfileGrpcHandler) GetAuthor(ctx context.Context, request *profile.GetAuthorRequest) (*profile.GetAuthorResponse, error) {
	if request == nil || request.Id == 0 {
		return nil, ErrInvalidRequest
	}
	authorData, err := p.authorService.GetAuthor(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	return authorData, nil
}

func (p ProfileGrpcHandler) GetAuthorByProfileID(ctx context.Context, request *profile.GetAuthorByProfileIDRequest) (*profile.GetAuthorByProfileIDResponse, error) {
	if request == nil || request.ProfileId == 0 {
		return nil, ErrInvalidRequest
	}
	authorData, err := p.authorService.GetAuthorByProfileID(ctx, request.ProfileId)
	if err != nil {
		return nil, err
	}
	return authorData, nil
}

func (p ProfileGrpcHandler) ListAuthors(ctx context.Context, request *profile.ListAuthorsRequest) (*profile.ListAuthorsResponse, error) {
	return p.authorService.ListAuthors(ctx)
}

func (p ProfileGrpcHandler) CreateAuthor(ctx context.Context, request *profile.CreateAuthorRequest) (*profile.CreateAuthorResponse, error) {
	if request == nil || request.Name == "" {
		return nil, ErrInvalidRequest
	}

	result, err := p.authorService.CreateAuthor(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p ProfileGrpcHandler) UpdateAuthor(ctx context.Context, request *profile.UpdateAuthorRequest) (*profile.UpdateAuthorResponse, error) {
	if request == nil || request.Id == 0 {
		return nil, ErrInvalidRequest
	}

	result, err := p.authorService.UpdateAuthor(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p ProfileGrpcHandler) DeleteAuthor(ctx context.Context, request *profile.DeleteAuthorRequest) (*profile.DeleteAuthorResponse, error) {
	if request == nil || request.Id == 0 {
		return nil, ErrInvalidRequest
	}
	return p.authorService.DeleteAuthor(ctx, request.Id)
}

func (p ProfileGrpcHandler) mustEmbedUnimplementedAuthorServiceServer() {
	//TODO implement me
	panic("implement me")
}

func NewProfileGrpcHandler(db *sql.DB) *ProfileGrpcHandler {
	return &ProfileGrpcHandler{
		profileService: NewProfileService(db),
		authorService:  NewAuthorService(db),
	}
}
