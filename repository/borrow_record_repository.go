package repository

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/frederikdaniel7/go-gin-library-api/constant"
	"github.com/frederikdaniel7/go-gin-library-api/database"
	"github.com/frederikdaniel7/go-gin-library-api/dto"
	"github.com/frederikdaniel7/go-gin-library-api/entity"
	"github.com/frederikdaniel7/go-gin-library-api/exception"
)

type BorrowRecordRepository interface {
	CreateBorrowRecord(ctx context.Context, body dto.CreateBorrowRecordBody, userId int64) (*entity.BorrowRecord, error)
	FindOneById(ctx context.Context, id int64) (*entity.BorrowRecord, error)
	UpdateRecordReturnBook(ctx context.Context, id int64) (*entity.BorrowRecord, error)
}

type borrowRecordRepository struct {
	db *sql.DB
}

func NewBorrowRecordRepository(db *sql.DB) *borrowRecordRepository {
	return &borrowRecordRepository{
		db: db,
	}
}

func (r *borrowRecordRepository) CreateBorrowRecord(ctx context.Context, body dto.CreateBorrowRecordBody, userId int64) (*entity.BorrowRecord, error) {
	record := entity.BorrowRecord{}

	q := `INSERT INTO borrow_records (user_id, book_id, status, borrowing_date) VALUES ($1, $2,$3,$4) 
	RETURNING id, user_id, book_id, status, borrowing_date,returning_date, created_at, updated_at, deleted_at`

	runner := database.PickQuerier(ctx, r.db)
	err := runner.QueryRowContext(ctx, q, userId, body.BookID, body.Status, body.BorrowingDate).Scan(&record.ID, &record.UserID, &record.BookID, &record.Status,
		&record.BorrowingDate, &record.ReturningDate, &record.CreatedAt, &record.UpdatedAt, &record.DeletedAt)
	if err != nil {
		return nil, err
	}
	return &record, nil

}

func (r *borrowRecordRepository) FindOneById(ctx context.Context, id int64) (*entity.BorrowRecord, error) {
	var borrowRecord entity.BorrowRecord

	q := `SELECT id, user_id, book_id, status, borrowing_date,returning_date, created_at, updated_at, deleted_at from borrow_records where id = $1`

	row := r.db.QueryRowContext(ctx, q, id)
	if row == nil {
		return nil, exception.NewErrorType(http.StatusBadRequest, constant.ResponseMsgBadRequest)
	}
	err := row.Scan(&borrowRecord.ID, &borrowRecord.UserID, &borrowRecord.BookID, &borrowRecord.Status,
		&borrowRecord.BorrowingDate, &borrowRecord.ReturningDate, &borrowRecord.CreatedAt, &borrowRecord.UpdatedAt, &borrowRecord.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &borrowRecord, nil
}

func (r *borrowRecordRepository) UpdateRecordReturnBook(ctx context.Context, id int64) (*entity.BorrowRecord, error) {
	var borrowRecord entity.BorrowRecord

	q := `UPDATE borrow_records SET status = 'returned', returning_date = now() where id = $1 
	RETURNING id, user_id, book_id, status, borrowing_date,returning_date, created_at, updated_at, deleted_at`
	row := r.db.QueryRowContext(ctx, q, id)
	if row == nil {
		return nil, exception.NewErrorType(http.StatusBadRequest, "here")
	}
	err := row.Scan(&borrowRecord.ID, &borrowRecord.UserID, &borrowRecord.BookID, &borrowRecord.Status,
		&borrowRecord.BorrowingDate, &borrowRecord.ReturningDate, &borrowRecord.CreatedAt, &borrowRecord.UpdatedAt, &borrowRecord.DeletedAt)
	if err != nil {
		return nil, exception.NewErrorType(http.StatusBadRequest, q)
	}
	return &borrowRecord, nil

}
