package entity

import (
	"database/sql"
	"time"
)

type BorrowRecord struct {
	ID            int64
	UserID        int64
	BookID        int64
	Status        string
	BorrowingDate time.Time
	ReturningDate sql.NullTime
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     sql.NullTime
}
