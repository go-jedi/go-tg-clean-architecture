package user

import "time"

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	IsBan     bool      `json:"is_ban"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
