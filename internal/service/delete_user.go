package service

import (
	"errors"
	"net/http"
	"refactoring/internal"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"refactoring/internal/storage"
)

var (
	UserNotFound = errors.New("user_not_found")
)

func DeleteUser(s *internal.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		if _, ok := s.List[id]; !ok {
			_ = render.Render(w, r, ErrInvalidRequest(UserNotFound))
			return
		}

		delete(s.List, id)

		err := storage.WriteStore(s)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		render.Status(r, http.StatusNoContent)
	}
}
