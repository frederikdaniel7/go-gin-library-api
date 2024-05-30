package dto

import (
	"time"
)

type BorrowRecord struct {
	ID            int64      `json:"id"`
	UserID        int64      `json:"user_id"`
	BookID        int64      `json:"book_id"`
	Status        string     `json:"status"`
	BorrowingDate time.Time  `json:"borrowing_date"`
	ReturningDate *time.Time `json:"returning_date,omitempty"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at,omitempty"`
}

type CreateBorrowRecordBody struct {
	BookID        int64     `json:"book_id" binding:"required"`
	Status        string    `json:"status" binding:"required"`
	BorrowingDate time.Time `json:"borrowing_date" binding:"required"`
}

type ReturnBookParam struct {
	ID int `uri:"id" binding:"required"`
}
