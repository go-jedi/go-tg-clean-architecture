package user

import "time"

type User struct {
	ID        int       `json:"id" validate:"required,min=1"`
	Username  string    `json:"username" validate:"required,min=1"`
	IsBan     bool      `json:"is_ban" validate:"boolean"`
	CreatedAt time.Time `json:"created_at" validate:"required"`
	UpdatedAt time.Time `json:"updated_at" validate:"required"`
}
