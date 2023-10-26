package service

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"refactoring/internal/storage"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	s, err := storage.ReadStore()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	id := chi.URLParam(r, "id")

	render.JSON(w, r, s.List[id])
}
