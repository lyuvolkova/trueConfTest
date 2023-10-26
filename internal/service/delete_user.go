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

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	s, err := storage.ReadStore()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	id := chi.URLParam(r, "id")

	if _, ok := s.List[id]; !ok {
		_ = render.Render(w, r, ErrInvalidRequest(UserNotFound))
		return
	}

	delete(s.List, id)

	err = storage.WriteStore(s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	render.Status(r, http.StatusNoContent)
}
