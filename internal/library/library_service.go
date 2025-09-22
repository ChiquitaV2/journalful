package library

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/chiquitav2/journalful/internal/db"
	"github.com/chiquitav2/journalful/pkg/library/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type LibraryServiceInterface interface {
	SaveArticleToLibrary(ctx context.Context, request *library.SaveArticleToLibraryRequest) (*library.SaveArticleToLibraryResponse, error)
	GetUserLibrary(ctx context.Context, request *library.GetUserLibraryRequest) (*library.GetUserLibraryResponse, error)
	GetLibrary(ctx context.Context, request *library.GetLibraryRequest) (*library.GetLibraryResponse, error)
	CreateLibrary(ctx context.Context, request *library.CreateLibraryRequest) (*library.CreateLibraryResponse, error)
	UpdateLibrary(ctx context.Context, request *library.UpdateLibraryRequest) (*library.UpdateLibraryResponse, error)
	DeleteLibrary(ctx context.Context, request *library.DeleteLibraryRequest) (*library.DeleteLibraryResponse, error)
}
type LibraryService struct {
	repo *db.Queries
}

func NewLibraryService(conn *sql.DB) *LibraryService {
	return &LibraryService{
		repo: db.New(conn),
	}
}

func (l *LibraryService) SaveArticleToLibrary(ctx context.Context, request *library.SaveArticleToLibraryRequest) (*library.SaveArticleToLibraryResponse, error) {
	result, err := l.repo.AddLibraryArticle(ctx, db.AddLibraryArticleParams{
		LibraryID:     request.LibraryId,
		ArticleID:     request.ArticleId,
		ReadingStatus: sql.NullInt16{Int16: int16(request.ReadingStatus), Valid: true},
		Notes:         sql.NullString{String: *request.Notes, Valid: request.Notes != nil},
	})
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &library.SaveArticleToLibraryResponse{Id: id}, nil
}

func (l *LibraryService) GetUserLibrary(ctx context.Context, request *library.GetUserLibraryRequest) (*library.GetUserLibraryResponse, error) {
	libraries, err := l.repo.ListLibrariesByUserID(ctx, request.UserId)
	if err != nil {
		return nil, err
	}
	var defaultLibrary *library.Library

	var response []*library.Library
	for _, lib := range libraries {
		builtLib, err := l.buildLibrary(ctx, lib)
		if err != nil {
			slog.Error("error building library", "error", err.Error())
			return nil, err
		}
		response = append(response, builtLib)
		if lib.Isdefault.Bool {
			defaultLibrary = builtLib
		}
	}
	if defaultLibrary == nil {
		id, err := l.createDefaultLibrary(ctx, request.UserId)
		if err != nil {
			slog.Error("error creating default library", "error", err.Error())
			return nil, err
		}
		lib, err := l.repo.GetLibrary(ctx, id)
		if err != nil {
			slog.Error("error getting default library", "error", err)
			return nil, err
		}
		defaultLibrary, err = l.buildLibrary(ctx, lib)
		if err != nil {
			slog.Error("error building default library", "error", err.Error())
			return nil, err
		}
	}

	return &library.GetUserLibraryResponse{
		DefaultLibrary:   defaultLibrary,
		PrivateLibraries: response,
	}, nil
}

func (s *LibraryService) createDefaultLibrary(ctx context.Context, userID int64) (int64, error) {
	result, err := s.repo.CreateLibrary(ctx, db.CreateLibraryParams{
		OwnerID: userID,
		Name: sql.NullString{
			String: "My Library",
			Valid:  true,
		},
		Description: sql.NullString{
			String: "This is my default library.",
			Valid:  true,
		},
		Ispublic: sql.NullBool{
			Bool:  true,
			Valid: true,
		},
		Isdefault: sql.NullBool{
			Bool:  true,
			Valid: true,
		},
	})
	if err != nil {
		return -1, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (s *LibraryService) buildLibrary(ctx context.Context, lib db.Library) (*library.Library, error) {
	articles, err := s.repo.ListLibraryArticlesByLibraryID(ctx, lib.ID)
	if err != nil {
		return nil, err
	}

	var libraryArticles []*library.LibraryArticle
	for _, article := range articles {
		notes := article.Notes.String
		libraryArticles = append(libraryArticles, &library.LibraryArticle{
			Id:              article.ID,
			ArticleId:       article.ArticleID,
			ReadingStatus:   library.ReadingStatus(article.ReadingStatus.Int16),
			DateAdded:       timestamppb.New(article.Dateadded.Time),
			Notes:           &notes,
			ArticleTitle:    article.ArticleTitle,
			Doi:             article.Doi,
			PublicationYear: article.PublicationYear.Int32,
		})
	}
	return &library.Library{
		Id:          lib.ID,
		OwnerId:     lib.OwnerID,
		Name:        lib.Name.String,
		Description: &lib.Description.String,
		IsPublic:    lib.Ispublic.Bool,
		Articles:    libraryArticles,
		CreatedAt:   timestamppb.New(lib.CreatedAt.Time),
		UpdatedAt:   timestamppb.New(lib.UpdatedAt.Time),
	}, nil
}

func (s *LibraryService) GetLibrary(ctx context.Context, request *library.GetLibraryRequest) (*library.GetLibraryResponse, error) {
	lib, err := s.repo.GetLibrary(ctx, request.LibraryId)
	if err != nil {
		return nil, err
	}

	builtLib, err := s.buildLibrary(ctx, lib)
	if err != nil {
		return nil, err
	}

	return &library.GetLibraryResponse{
		Library: builtLib,
	}, nil
}

func (s *LibraryService) CreateLibrary(ctx context.Context, request *library.CreateLibraryRequest) (*library.CreateLibraryResponse, error) {
	result, err := s.repo.CreateLibrary(ctx, db.CreateLibraryParams{
		OwnerID: request.OwnerId,
		Name: sql.NullString{
			String: request.Name,
			Valid:  true,
		},
		Description: sql.NullString{
			String: *request.Description,
			Valid:  request.Description != nil,
		},
		Ispublic: sql.NullBool{
			Bool:  request.IsPublic,
			Valid: true,
		},
		Isdefault: sql.NullBool{
			Bool:  false, // New libraries are not default
			Valid: true,
		},
	})
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &library.CreateLibraryResponse{
		LibraryId: id,
	}, nil
}

func (s *LibraryService) UpdateLibrary(ctx context.Context, request *library.UpdateLibraryRequest) (*library.UpdateLibraryResponse, error) {
	// Get current library to check ownership
	_, err := s.repo.GetLibrary(ctx, request.LibraryId)
	if err != nil {
		return nil, err
	}

	// Update name if provided
	if request.Name != nil {
		err = s.repo.UpdateLibraryName(ctx, db.UpdateLibraryNameParams{
			ID:   request.LibraryId,
			Name: sql.NullString{String: *request.Name, Valid: true},
		})
		if err != nil {
			return nil, err
		}
	}

	// Update description if provided
	if request.Description != nil {
		err = s.repo.UpdateLibraryDescription(ctx, db.UpdateLibraryDescriptionParams{
			ID:          request.LibraryId,
			Description: sql.NullString{String: *request.Description, Valid: true},
		})
		if err != nil {
			return nil, err
		}
	}

	// Update visibility if provided
	if request.IsPublic != nil {
		err = s.repo.UpdateLibraryVisibility(ctx, db.UpdateLibraryVisibilityParams{
			ID:       request.LibraryId,
			Ispublic: sql.NullBool{Bool: *request.IsPublic, Valid: true},
		})
		if err != nil {
			return nil, err
		}
	}

	return &library.UpdateLibraryResponse{
		Success: true,
	}, nil
}

func (s *LibraryService) DeleteLibrary(ctx context.Context, request *library.DeleteLibraryRequest) (*library.DeleteLibraryResponse, error) {
	err := s.repo.DeleteLibrary(ctx, request.LibraryId)
	if err != nil {
		return nil, err
	}

	return &library.DeleteLibraryResponse{
		Success: true,
	}, nil
}
