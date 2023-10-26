package service

import (
	"net/http"

	"github.com/go-chi/render"

	"refactoring/internal/storage"
)

func SearchUsers(w http.ResponseWriter, r *http.Request) {
	s, err := storage.ReadStore()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	render.JSON(w, r, s.List)
}
