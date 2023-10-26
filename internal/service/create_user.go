package service

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/render"

	"refactoring/internal"
	"refactoring/internal/storage"
)

type CreateUserRequest struct {
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
}

func (c *CreateUserRequest) Bind(r *http.Request) error { return nil }

func CreateUser(s *internal.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := CreateUserRequest{}

		if err := render.Bind(r, &request); err != nil {
			_ = render.Render(w, r, ErrInvalidRequest(err))
			return
		}

		s.Increment++
		u := internal.User{
			CreatedAt:   time.Now(),
			DisplayName: request.DisplayName,
			Email:       request.DisplayName,
		}

		id := strconv.Itoa(s.Increment)
		s.List[id] = u

		err := storage.WriteStore(s)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		render.Status(r, http.StatusCreated)
		render.JSON(w, r, map[string]interface{}{
			"user_id": id,
		})
	}
}
