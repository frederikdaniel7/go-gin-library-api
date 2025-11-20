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

type BookPackage struct {
	Books     []BookDetail
	Count     int
	TotalData int
}

type BookDetail struct {
	ID          int64
	Title       string
	Description string
	Quantity    int
	Author      *Author
	Cover       sql.NullString
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   sql.NullTime
}
