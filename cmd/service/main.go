package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"refactoring/internal/service"
	"refactoring/internal/storage"
)

func main() {
	repo := storage.NewRepository()
	err := repo.Load()
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
				r.Get("/", service.SearchUsers(repo))
				r.Post("/", service.CreateUser(repo))

				r.Route("/{id}", func(r chi.Router) {
					r.Get("/", service.GetUser(repo))
					r.Patch("/", service.UpdateUser(repo))
					r.Delete("/", service.DeleteUser(repo))
				})
			})
		})
	})

	http.ListenAndServe(":3333", r)
}
