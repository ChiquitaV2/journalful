package library

import (
	"context"
	"database/sql"
	"github.com/chiquitav2/journalful/internal/db"
	"github.com/chiquitav2/journalful/pkg/library/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type LibraryServiceInterface interface {
	SaveArticleToLibrary(ctx context.Context, request *library.SaveArticleToLibraryRequest) (*library.SaveArticleToLibraryResponse, error)
	GetUserLibrary(ctx context.Context, request *library.GetUserLibraryRequest) (*library.GetUserLibraryResponse, error)
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
		articles, err := l.repo.ListLibraryArticlesByLibraryID(ctx, lib.ID)
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
		name := lib.Name.String
		response = append(response, &library.Library{
			Id:        lib.ID,
			OwnerId:   lib.OwnerID,
			Name:      name,
			CreatedAt: timestamppb.New(lib.CreatedAt.Time),
			UpdatedAt: timestamppb.New(lib.UpdatedAt.Time),
			Articles:  libraryArticles,
		})
		if lib.Isdefault.Bool {
			defaultLibrary = &library.Library{
				Id:        lib.ID,
				OwnerId:   lib.OwnerID,
				Name:      name,
				CreatedAt: timestamppb.New(lib.CreatedAt.Time),
				UpdatedAt: timestamppb.New(lib.UpdatedAt.Time),
				Articles:  libraryArticles,
			}
		}
	}

	return &library.GetUserLibraryResponse{
		DefaultLibrary:   defaultLibrary,
		PrivateLibraries: response,
	}, nil
}
