package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/bgw6/pod/internal/domain"
	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

type productHandler struct {
	service domain.Service
}

func (ph productHandler) FindAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products, err := ph.service.Get()

		if err != nil {
			switch {
			default:
				{
					response.Error(w, http.StatusInternalServerError, "")
				}
			}
			return
		}

		if len(products) == 0 {
			response.JSON(w, http.StatusNoContent, nil)
			return
		}

		response.JSON(w, http.StatusOK, products)
	}
}

type RequestBodyProduct struct {
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

func (ph productHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqBody RequestBodyProduct
		if err := request.JSON(r, &reqBody); err != nil {
			code := http.StatusUnprocessableEntity
			body := map[string]any{"message": "invalid request body", "data": nil}

			response.JSON(w, code, body)
			return
		}

		pr := domain.Product{
			Name:     reqBody.Name,
			Type:     reqBody.Type,
			Quantity: reqBody.Quantity,
			Price:    reqBody.Price,
		}

		err := ph.service.Create(&pr)

		if err != nil {
			var code int
			var body map[string]any
			switch {
			default:
				code = http.StatusInternalServerError
				body = map[string]any{"message": "internal server error", "data": nil}
			}
			response.JSON(w, code, body)
			return
		}

		code := http.StatusCreated
		body := map[string]any{"message": "product created", "data": pr}

		response.JSON(w, code, body)
	}
}

type RequestBodyUpdate struct {
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

func (ph productHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlId := chi.URLParam(r, "productId")

		if urlId == "" {
			code := http.StatusBadRequest
			body := map[string]any{"message": "product id is required", "data": nil}

			response.JSON(w, code, body)
			return
		}

		// request
		id, err := strconv.Atoi(urlId)
		if err != nil {
			code := http.StatusBadRequest
			body := map[string]any{"message": "invalid id", "data": nil}

			response.JSON(w, code, body)
			return
		}

		var reqBody RequestBodyUpdate
		if err := request.JSON(r, &reqBody); err != nil {
			code := http.StatusUnprocessableEntity
			body := map[string]any{"message": "invalid request body", "data": nil}

			response.JSON(w, code, body)
			return
		}

		pr := domain.Product{
			Id:       id,
			Name:     reqBody.Name,
			Type:     reqBody.Type,
			Quantity: reqBody.Quantity,
			Price:    reqBody.Price,
		}
		// -> update or create
		if err := ph.service.Update(&pr); err != nil {
			var code int
			var body map[string]any
			switch {
			case errors.Is(err, domain.ErrResourceNotFound):
				code = http.StatusNotFound
				body = map[string]any{"message": "product not found", "data": nil}
			default:
				code = http.StatusInternalServerError
				body = map[string]any{"message": "internal server error", "data": nil}
			}
			response.JSON(w, code, body)
			return
		}

		code := http.StatusOK
		body := map[string]any{"message": "product updated", "data": pr}

		response.JSON(w, code, body)
	}

}

func (ph productHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// request
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			code := http.StatusBadRequest
			body := map[string]any{"message": "invalid id", "data": nil}

			response.JSON(w, code, body)
			return
		}

		err = ph.service.Delete(id)
		if err != nil {
			var code int
			var body map[string]any
			switch {
			case errors.Is(err, domain.ErrResourceNotFound):
				code = http.StatusNotFound
				body = map[string]any{"message": "product not found", "data": nil}
			default:
				code = http.StatusInternalServerError
				body = map[string]any{"message": "internal server error", "data": nil}
			}
			response.JSON(w, code, body)
			return
		}

		code := http.StatusNoContent
		body := map[string]any(nil)

		response.JSON(w, code, body)
	}
}

func NewProductHandler(service domain.Service) productHandler {
	return productHandler{service}
}

/*
1xx - nao tenho ideia

2xx - status de sucesso
 200 - OK
 201 - Created
 202 - Accepted
 204 - NoContent

3xx - cache e redirecionamento

4xx - erros do cliente
 400 - Bad Request
 401 - Unauthorized
 403 - Forbideen
 404 - Not Found
 429 - Conflict
 422 - Unprocessable Entity
 459 - Too many request

5xx - erros do servidor
 500 - Internal Server Error
 502 - Bad Gateway

 /customers/{customerId}/usuarios/{userId}
*/
