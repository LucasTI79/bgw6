package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type router struct {
}

func (router *router) MapRoutes() http.Handler {
	r := chi.NewRouter()

	r.Route("/products", func(rp chi.Router) {
		rp.Mount("/", buildProductsRoutes())
	})

	return r
}

func NewRouter() *router {
	return &router{}
}
