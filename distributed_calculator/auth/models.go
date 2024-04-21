package auth

import "time"

// User представляет собой модель пользователя в базе данных.
type User struct {
	ID        int64     json:"id"
	Login     string    json:"login"
	Password  string    json:"password"
	CreatedAt time.Time json:"created_at"
	UpdatedAt time.Time json:"updated_at"
}