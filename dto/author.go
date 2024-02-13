package dto

import "time"

type Author struct {
	ID        *int64     `json:"id"`
	Name      *string    `json:"author_name"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
