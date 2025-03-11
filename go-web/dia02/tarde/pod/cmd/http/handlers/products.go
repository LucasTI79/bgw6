package handlers

import (
	"net/http"

	"github.com/bgw6/pod/internal/domain"
	"github.com/bgw6/pod/pkg/web"
)

type productHandler struct {
	service domain.Service
}

func (ph productHandler) FindAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products, err := ph.service.FindAll()

		if err != nil {
			switch {
			default:
				{
					web.ReponseJson(w, http.StatusInternalServerError, nil, "")
				}
			}
			return
		}

		if len(products) == 0 {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		web.ReponseJson(w, http.StatusOK, products, "lista de produtos")
	}
}

func NewProductHandler(service domain.Service) productHandler {
	return productHandler{service}
}
