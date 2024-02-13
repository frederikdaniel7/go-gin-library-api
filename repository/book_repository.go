package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/entity"
)

type BookRepository interface {
	FindAll() ([]entity.Book, error)
	FindOneBookByTitle(title string) ([]entity.Book, error)
	CreateBook(body dto.CreateBookBody) (*entity.Book, error)
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

func (r *bookRepository) CreateBook(body dto.CreateBookBody) (*entity.Book, error) {
	book := entity.Book{}

	var sb strings.Builder
	sb.WriteString("INSERT INTO books (title, book_description, quantity")
	if body.Cover == "" {
		sb.WriteString(") VALUES (")
		for i := 1; i < 4; i++ {
			sb.WriteString("$" + fmt.Sprintf("%d", i))
			if i != 3 {
				sb.WriteString(",")
			}
		}
	} else {
		sb.WriteString(",cover) VALUES (")
		for i := 1; i < 5; i++ {
			sb.WriteString("$" + fmt.Sprintf("%d", i))
			if i != 4 {
				sb.WriteString(",")
			}
		}
	}
	sb.WriteString(")returning id, title, book_description,cover, created_at, updated_at, deleted_at;")
	if body.Cover != "" {
		err := r.db.QueryRow(sb.String(), body.Title, body.Description, body.Quantity, body.Cover).Scan(
			&book.ID, &book.Title, &book.Description, &book.Cover, &book.CreatedAt, &book.UpdatedAt, &book.DeletedAt)
		if err != nil {
			return nil, errors.New(sb.String())
		}

		return &book, nil
	}
	err := r.db.QueryRow(sb.String(), body.Title, body.Description, body.Quantity).Scan(
		&book.ID, &book.Title, &book.Description, &book.Cover, &book.CreatedAt, &book.UpdatedAt, &book.DeletedAt)
	if err != nil {
		return nil, errors.New(sb.String())
	}

	return &book, nil
}
