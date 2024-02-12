package repository

import (
	"database/sql"
	"log"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/entity"
)

type BookRepository interface {
	FindAll() ([]entity.Book, error)
	FindOneBookByTitle(title string) ([]entity.Book, error)
}

type bookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) *bookRepository {
	return &bookRepository{
		db: db,
	}
}

func (r *bookRepository) FindAll() ([]entity.Book, error) {
	books := []entity.Book{}

	q := `SELECT id,title,book_description, quantity,cover,created_at,updated_at,deleted_at from books`

	rows, err := r.db.Query(q)
	if err != nil {
		log.Println("query error", err)
	}
	defer rows.Close()

	for rows.Next() {
		var book entity.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Description, &book.Quantity, &book.Cover,
			&book.CreatedAt, &book.UpdatedAt, &book.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (r *bookRepository) FindOneBookByTitle(title string) ([]entity.Book, error) {
	books := []entity.Book{}

	q := `SELECT id,title,book_description, quantity,cover,created_at,updated_at,deleted_at from books
	where title ILIKE '%' ||$1|| '%'`

	rows, err := r.db.Query(q, title)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var book entity.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Description, &book.Quantity, &book.Cover,
			&book.CreatedAt, &book.UpdatedAt, &book.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return books, nil

}
