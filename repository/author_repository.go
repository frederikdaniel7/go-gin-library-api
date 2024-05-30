package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/frederikdaniel7/go-gin-library-api/entity"
)

type AuthorRepository interface {
	FindOneById(ctx context.Context, id int64) (*entity.Author, error)
}

type authorRepository struct {
	db *sql.DB
}

func NewAuthorRepository(db *sql.DB) *authorRepository {
	return &authorRepository{
		db: db,
	}
}

func (r *authorRepository) FindOneById(ctx context.Context, id int64) (*entity.Author, error) {
	var author entity.Author

	q := `SELECT a.id, a.author_name from author a where a.id = $1`

	row := r.db.QueryRow(q, id)
	if row == nil {
		return nil, errors.New("error query")
	}
	row.Scan(&author.ID, &author.Name)

	return &author, nil
}
