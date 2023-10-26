package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"refactoring/internal/service"
)

func main() {
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
				r.Get("/", service.SearchUsers)
				r.Post("/", service.CreateUser)

				r.Route("/{id}", func(r chi.Router) {
					r.Get("/", service.GetUser)
					r.Patch("/", service.UpdateUser)
					r.Delete("/", service.DeleteUser)
				})
			})
		})
	})

	http.ListenAndServe(":3333", r)
}
