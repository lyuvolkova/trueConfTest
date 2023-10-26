package main

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"refactoring/internal/service"
	"refactoring/internal/storage"
)

func main() {
	sMutex := new(sync.Mutex)
	s, err := storage.ReadStore()
	if err != nil {
		log.Fatalln(err)
	}
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	//r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(time.Now().String()))
	})

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/users", func(r chi.Router) {
				r.Get("/", service.SearchUsers(s))
				r.Post("/", service.CreateUser(s, sMutex))

				r.Route("/{id}", func(r chi.Router) {
					r.Get("/", service.GetUser(s))
					r.Patch("/", service.UpdateUser(s, sMutex))
					r.Delete("/", service.DeleteUser(s, sMutex))
				})
			})
		})
	})

	http.ListenAndServe(":3333", r)
}
