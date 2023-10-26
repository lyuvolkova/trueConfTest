package service

import (
	"net/http"

	"github.com/go-chi/render"

	"refactoring/internal/storage"
)

type CreateUserRequest struct {
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
}

func (c *CreateUserRequest) Bind(r *http.Request) error { return nil }

func CreateUser(repo repo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := CreateUserRequest{}

		if err := render.Bind(r, &request); err != nil {
			_ = render.Render(w, r, ErrInvalidRequest(err))
			return
		}

		id, err := repo.CreateUser(request.DisplayName, request.Email)
		if err != nil {
			switch err {
			case storage.ErrStorageBusy:
				_ = render.Render(w, r, ErrConflict(err))
			default:
				_ = render.Render(w, r, ErrServer(err))
			}
			return
		}

		render.Status(r, http.StatusCreated)
		render.JSON(w, r, map[string]interface{}{
			"user_id": id,
		})
	}
}
