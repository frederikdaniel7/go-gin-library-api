package repository

import (
	"database/sql"
	"errors"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/entity"
)

type AuthorRepository interface {
	FindOneById(id int64) (*entity.Author, error)
}

type authorRepository struct {
	db *sql.DB
}

func NewAuthorRepository(db *sql.DB) *authorRepository {
	return &authorRepository{
		db: db,
	}
}

func (r *authorRepository) FindOneById(id int64) (*entity.Author, error) {
	var author entity.Author

	q := `SELECT a.id, a.author_name from author a where a.id = $1`

	row := r.db.QueryRow(q, id)
	if row == nil {
		return nil, errors.New("error query")
	}
	row.Scan(&author.ID, &author.Name)

	return &author, nil
}
