package dto

import (
	"time"
)

type Book struct {
	ID          int64      `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Quantity    int        `json:"quantity"`
	Cover       *string    `json:"cover,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

type BookDetail struct {
	ID          int64      `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Quantity    int        `json:"quantity"`
	Cover       *string    `json:"cover,omitempty"`
	Author      *Author    `json:"author,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

type CreateBookBody struct {
	Title       string `json:"title" binding:"required,lte=35"`
	Description string `json:"description" binding:"required"`
	Quantity    *int   `json:"quantity" binding:"required,min=0"`
	Cover       string `json:"cover"`
	AuthorID    *int64 `json:"author_id"`
}
