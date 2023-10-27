package service

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"refactoring/internal/storage"
)

var (
	UserNotFound = errors.New("user_not_found")
)

func DeleteUser(repo Repo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		err := repo.DeleteUser(id)
		if err != nil {
			switch err {
			case storage.ErrUserNotFound:
				_ = render.Render(w, r, ErrNotFound(UserNotFound))
			case storage.ErrStorageBusy:
				_ = render.Render(w, r, ErrConflict(err))
			default:
				_ = render.Render(w, r, ErrServer(err))
			}
			return
		}

		render.Status(r, http.StatusNoContent)
	}
}
