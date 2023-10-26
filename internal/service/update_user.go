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

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	s, err := storage.ReadStore()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	request := UpdateUserRequest{}

	if err := render.Bind(r, &request); err != nil {
		_ = render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	id := chi.URLParam(r, "id")

	if _, ok := s.List[id]; !ok {
		_ = render.Render(w, r, ErrInvalidRequest(UserNotFound))
		return
	}

	u := s.List[id]
	u.DisplayName = request.DisplayName
	s.List[id] = u

	err = storage.WriteStore(s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	render.Status(r, http.StatusNoContent)
}