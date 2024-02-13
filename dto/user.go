package dto

import "time"

type User struct {
	ID        int64      `json:"id"`
	Name      string     `json:"user_name"`
	Email     string     `json:"email"`
	Phone     string     `json:"phone"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
