package service

import (
	"net/http"
	"refactoring/internal"

	"github.com/go-chi/render"
)

func SearchUsers(s *internal.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, s.List)
	}
}
