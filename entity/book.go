package entity

import "database/sql"

type Book struct {
	ID int64 `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Quantity int `json:"quantity"`
	Cover string `json:"cover"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}