package service

import (
	"net/http"
	"refactoring/internal"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func GetUser(s *internal.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		render.JSON(w, r, s.List[id])
	}
}
