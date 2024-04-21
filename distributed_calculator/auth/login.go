package auth

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
	"github.com/jackc/pgx/v4"
)

// LoginUser авторизует пользователя и возвращает JWT токен.
func LoginUser(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}

		var hashedPassword string
		err = pool.QueryRow(context.Background(), SELECT password FROM users WHERE login = $1, user.Login).Scan(&hashedPassword)
		if err != nil {
			if err == pgx.ErrNoRows {
				render.Render(w, r, ErrNotFound)
				return
			}
			render.Render(w, r, ErrInternalServer(err))
			return
		}

		if user.Password != hashedPassword {
			render.Render(w, r, ErrInvalidPassword)
			return
		}

		token, err := GenerateJWT(user.ID)
		if err != nil {
			render.Render(w, r, ErrInternalServer(err))
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, map[string]string{"token": token})
	}
}