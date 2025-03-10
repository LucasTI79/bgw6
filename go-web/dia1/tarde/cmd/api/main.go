package main

import (
	"banana/cmd/api/handlers"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	db := map[string]string{
		"1": "Natan",
		"2": "Lucas",
		"3": "Charles",
	}

	handler := handlers.NewEmployeeHanlder(db)

	r.Get("/employee", handler.GetAll())
	r.Get("/employee/{id}", handler.GetById())

	log.Println("Servidor rodando em http://localhost:8080")
	http.ListenAndServe("localhost:8080", r)
}
