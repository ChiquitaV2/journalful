package article

import (
	"context"
	"database/sql"
	"testing"

	"github.com/chiquitav2/journalful/internal/db"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock db.Querier
type MockQueries struct {
	mock.Mock
}

func (m *MockQueries) GetArticle(ctx context.Context, id int64) (db.Article, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(db.Article), args.Error(1)
}

func (m *MockQueries) ListArticleAuthorsByArticleID(ctx context.Context, articleID int64) ([]db.ListArticleAuthorsByArticleIDRow, error) {
	args := m.Called(ctx, articleID)
	return args.Get(0).([]db.ListArticleAuthorsByArticleIDRow), args.Error(1)
}

func (m *MockQueries) ListArticlesWithAuthors(ctx context.Context) ([]db.ListArticlesWithAuthorsRow, error) {
	args := m.Called(ctx)
	return args.Get(0).([]db.ListArticlesWithAuthorsRow), args.Error(1)
}

func (m *MockQueries) CreateArticle(ctx context.Context, params db.CreateArticleParams) (sql.Result, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(sql.Result), args.Error(1)
}

func (m *MockQueries) UpdateArticle(ctx context.Context, params db.UpdateArticleParams) error {
	args := m.Called(ctx, params)
	return args.Error(0)
}

func (m *MockQueries) DeleteArticle(ctx context.Context, id int64) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockQueries) AddArticleAuthor(ctx context.Context, params db.AddArticleAuthorParams) (sql.Result, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(sql.Result), args.Error(1)
}

func (m *MockQueries) GetAuthorByName(ctx context.Context, name string) (db.Author, error) {
	args := m.Called(ctx, name)
	return args.Get(0).(db.Author), args.Error(1)
}

func (m *MockQueries) CreateAuthor(ctx context.Context, params db.CreateAuthorParams) (sql.Result, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(sql.Result), args.Error(1)
}

func TestArticleService_GetArticle(t *testing.T) {
	// Create a new mock querier
	mockQueries := new(MockQueries)

	// Create a new article service with the mock querier
	articleService := &ArticleSerivceImp{
		queries: mockQueries,
	}

	// Set up the mock expectations
	articleID := int64(1)
	expectedArticle := db.Article{
		ID:    articleID,
		Title: "Test Article",
	}
	expectedAuthors := []db.ListArticleAuthorsByArticleIDRow{
		{
			AuthorID:   1,
			AuthorName: "Test Author",
		},
	}

	mockQueries.On("GetArticle", mock.Anything, articleID).Return(expectedArticle, nil)
	mockQueries.On("ListArticleAuthorsByArticleID", mock.Anything, articleID).Return(expectedAuthors, nil)

	// Call the GetArticle method
	response, err := articleService.GetArticle(context.Background(), articleID)

	// Assert the results
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, expectedArticle.ID, response.Article.Id)
	assert.Equal(t, expectedArticle.Title, response.Article.Title)
	assert.Equal(t, len(expectedAuthors), len(response.Article.Authors))
	assert.Equal(t, expectedAuthors[0].AuthorName, response.Article.Authors[0].Name)

	// Assert that the mock expectations were met
	mockQueries.AssertExpectations(t)
}
