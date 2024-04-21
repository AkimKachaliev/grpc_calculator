package auth

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
	"github.com/jackc/pgx/v4"
)

// RegisterUser регистрирует нового пользователя.
func RegisterUser(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}

		if _, err := pool.Exec(context.Background(), INSERT INTO users (login, password) VALUES ($1, $2), user.Login, user.Password); err != nil {
			render.Render(w, r, ErrInternalServer(err))
			return
		}

		render.Status(r, http.StatusCreated)
		render.Render(w, r, render.JSON(user))
	}
}