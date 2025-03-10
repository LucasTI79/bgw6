package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// kebab-case
// camelCase

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// rota, endpoint, ponto de extremidade
	r.Get("/hello-world", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"Hello World!"}`))
	})

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("ops, um erro aconteceu: %s", err)
	}
}
