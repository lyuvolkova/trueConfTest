package service

import (
	"errors"
	"net/http"
	"refactoring/internal"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"refactoring/internal/storage"
)

var (
	UserNotFound = errors.New("user_not_found")
)

func DeleteUser(s *internal.UserStore, sMutex *sync.Mutex) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		if !sMutex.TryLock() {
			http.Error(w, "file storage busy", http.StatusConflict)
			return
		}
		defer sMutex.Unlock()

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
