package main

import (
	"log"
	"net/http"
	"os"
	"refactoring/internal/router"
	"refactoring/internal/storage"
)

func main() {
	store := os.Getenv("FILE_NAME")
	if store == "" {
		store = `users.json`
	}

	repo := storage.NewRepository(store)
	err := repo.Load()
	if err != nil {
		log.Fatalln(err)
	}

	r := router.Router(repo)

	http.ListenAndServe(":3333", r)
}
