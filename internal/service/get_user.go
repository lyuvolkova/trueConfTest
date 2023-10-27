package service

import (
	"net/http"
	"refactoring/internal/storage"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func GetUser(repo Repo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		user, err := repo.GetUser(id)
		if err != nil {
			switch err {
			case storage.ErrUserNotFound:
				_ = render.Render(w, r, ErrNotFound(UserNotFound))
			default:
				_ = render.Render(w, r, ErrServer(err))
			}
			return
		}
		render.JSON(w, r, user)
	}
}
