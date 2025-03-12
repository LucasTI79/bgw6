package products

import "github.com/bgw6/pod/internal/domain"

func NewRepository() domain.Repository {
	return &mapRepository{
		products: make(map[int]ProductAttributes),
	}
}
