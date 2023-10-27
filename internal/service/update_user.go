package service

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"refactoring/internal/storage"
)

type UpdateUserRequest struct {
	DisplayName string `json:"display_name"`
}

func (c *UpdateUserRequest) Bind(r *http.Request) error { return nil }

func UpdateUser(repo Repo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := UpdateUserRequest{}

		if err := render.Bind(r, &request); err != nil {
			_ = render.Render(w, r, ErrInvalidRequest(err))
			return
		}

		id := chi.URLParam(r, "id")
		err := repo.UpdateUser(id, request.DisplayName)
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
