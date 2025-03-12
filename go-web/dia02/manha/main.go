package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Product estrutura que contém as informações de um produto
type Product struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

// ControllerProducts estrutura que contém o armazenamento dos produtos
type ControllerProducts struct {
	storage map[int]*Product
}

// RequestBodyProduct estrutura para o corpo da solicitação
type RequestBodyProduct struct {
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

// ResponseBody estrutura para o corpo da resposta
type ResponseBody struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   bool        `json:"error"`
}

// ResponseBodyProduct estrutura para o corpo da resposta
type ResponseBodyProduct struct {
	Message string `json:"message"`
	Data    *struct {
		Id       int     `json:"id"`
		Name     string  `json:"name"`
		Type     string  `json:"type"`
		Quantity int     `json:"quantity"`
		Price    float64 `json:"price"`
	} `json:"data"`
	Error bool `json:"error"`
}

// Create método para lidar com a criação de um produto
func (c *ControllerProducts) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqBody RequestBodyProduct
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ResponseBodyProduct{Message: "Bad Request", Error: true})
			return
		}

		pr := &Product{
			Id:       len(c.storage) + 1,
			Name:     reqBody.Name,
			Type:     reqBody.Type,
			Quantity: reqBody.Quantity,
			Price:    reqBody.Price,
		}
		c.storage[pr.Id] = pr

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(ResponseBodyProduct{
			Message: "Product created",
			Data: &struct {
				Id       int     `json:"id"`
				Name     string  `json:"name"`
				Type     string  `json:"type"`
				Quantity int     `json:"quantity"`
				Price    float64 `json:"price"`
			}{
				Id:       pr.Id,
				Name:     pr.Name,
				Type:     pr.Type,
				Quantity: pr.Quantity,
				Price:    pr.Price,
			},
			Error: false,
		})
	}
}

// ListProducts método para listar produtos com verificação de token
func (c *ControllerProducts) ListProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Token")
		if token != "123456" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(ResponseBody{Message: "Unauthorized", Error: true})
			return
		}

		// Retorna a lista de produtos
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(ResponseBody{Message: "Products retrieved", Data: c.storage, Error: false})
	}
}

func main() {
	db := make(map[int]*Product)
	ct := &ControllerProducts{storage: db}

	rt := chi.NewRouter()
	rt.Post("/products", ct.Create())

	http.ListenAndServe(":8080", rt)
}
