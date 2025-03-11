package router

import (
	"net/http"

	"github.com/bgw6/pod/cmd/http/handlers"
	"github.com/bgw6/pod/internal/products"
	"github.com/go-chi/chi/v5"
)

func buildProductsRoutes() http.Handler {
	r := chi.NewRouter()
	repository := products.NewRepository()
	service := products.NewService(repository)
	handler := handlers.NewProductHandler(service)

	r.Get("/", handler.FindAll())

	return r
}
