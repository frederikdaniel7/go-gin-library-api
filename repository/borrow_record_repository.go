package repository

import (
	"database/sql"
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/entity"
)

type BorrowRecordRepository interface {
	CreateBorrowRecord(body dto.CreateBorrowRecordBody) (*entity.BorrowRecord, error)
}

type borrowRecordRepository struct {
	db *sql.DB
}

func NewBorrowRecordRepository(db *sql.DB) *borrowRecordRepository {
	return &borrowRecordRepository{
		db: db,
	}
}

func (r *borrowRecordRepository) CreateBorrowRecord(body dto.CreateBorrowRecordBody) (*entity.BorrowRecord, error) {
	record := entity.BorrowRecord{}

	now := time.Now()
	q := `INSERT INTO borrow_records (user_id, book_id, status, borrowing_date) VALUES ($1, $2,$3,$4,$5,$6) 
	RETURNING id, user_id, book_id, status, borrowing_date,returning_date, created_at, updated_at, deleted_at`
	err := r.db.QueryRow(q, body.UserID, body.BookID, body.Status, now).Scan(&record.ID, &record.UserID, &record.BookID, &record.Status,
		&record.BorrowingDate, &record.ReturningDate, &record.CreatedAt, &record.UpdatedAt, &record.DeletedAt)
	if err != nil {
		return nil, err
	}
	return &record, nil

}
