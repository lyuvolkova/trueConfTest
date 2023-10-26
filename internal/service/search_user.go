package service

import (
	"net/http"

	"github.com/go-chi/render"
)

func SearchUsers(repo repo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		list, err := repo.SearchUsers()
		if err != nil {
			_ = render.Render(w, r, ErrServer(err))
			return
		}
		render.JSON(w, r, list)
	}
}
