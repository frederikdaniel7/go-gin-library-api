package entity

import (
	"database/sql"
	"time"
)

type Book struct {
	ID          int64
	Title       string
	Description string
	Quantity    int
	Cover       sql.NullString
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   sql.NullTime
}
