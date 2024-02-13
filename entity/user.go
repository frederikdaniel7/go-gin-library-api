package entity

import (
	"database/sql"
	"time"
)

type User struct {
	ID    int64 
	Name  string
	Email string
	Phone string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   sql.NullTime
}
