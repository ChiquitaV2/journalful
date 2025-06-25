package article

import (
	"database/sql"
	"github.com/chiquitav2/journalful/internal/db"
)

type articleRepository struct {
	query *db.Queries
}

// NewArticleRepository creates a new instance of articleRepository
func NewArticleRepository(conn *sql.DB) *articleRepository {
	return &articleRepository{
		query: db.New(conn),
	}
}
