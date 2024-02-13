package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/constant"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/entity"
)

type BookRepository interface {
	FindAll() ([]entity.BookDetail, error)
	FindSimilarBookByTitle(title string) ([]entity.BookDetail, error)
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

func (r *bookRepository) FindAll() ([]entity.BookDetail, error) {
	books := []entity.BookDetail{}

	q := `SELECT b.id,b.title,b.book_description, b.quantity,b.cover,a.id, a.author_name, b.created_at,b.updated_at,b.deleted_at from books b
	LEFT JOIN author a ON a.id = b.author_id`

	rows, err := r.db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		author := entity.Author{}
		book := entity.BookDetail{}

		err := rows.Scan(&book.ID, &book.Title, &book.Description, &book.Quantity, &book.Cover,
			&author.ID, &author.Name, &book.CreatedAt, &book.UpdatedAt, &book.DeletedAt)
		if err != nil {
			fmt.Println("error query")
			return nil, err
		}
		if author.ID != nil {
			book.Author = &author
		}
		books = append(books, book)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (r *bookRepository) FindSimilarBookByTitle(title string) ([]entity.BookDetail, error) {
	books := []entity.BookDetail{}

	q := `SELECT b.id,b.title,b.book_description, b.quantity,b.cover,a.id,a.author_name, b.created_at,b.updated_at,b.deleted_at from books b 
	LEFT JOIN author a ON a.id = b.author_id
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
		author := entity.Author{}
		book := entity.BookDetail{}

		err := rows.Scan(&book.ID, &book.Title, &book.Description, &book.Quantity, &book.Cover,
			&author.ID, &author.Name,
			&book.CreatedAt, &book.UpdatedAt, &book.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		if author.ID != nil {
			book.Author = &author
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
	sb.WriteString("INSERT INTO books (title, book_description, quantity, cover, author_id)")
	sb.WriteString("VALUES (")
	for i := 1; i < constant.LenCreateBody; i++ {
		sb.WriteString("$" + fmt.Sprintf("%d", i))
		if i != 5 {
			sb.WriteString(",")
		}
	}
	sb.WriteString(")returning id, title, book_description, quantity, cover, created_at, updated_at, deleted_at;")
	err := r.db.QueryRow(sb.String(), body.Title, body.Description, body.Quantity, body.Cover, body.AuthorID).Scan(
		&book.ID, &book.Title, &book.Description, &book.Quantity, &book.Cover, &book.CreatedAt, &book.UpdatedAt, &book.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &book, nil
}
