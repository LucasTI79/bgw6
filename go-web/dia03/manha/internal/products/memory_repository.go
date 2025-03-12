package products

import "github.com/bgw6/pod/internal/domain"

var products []domain.Product

type memoryRepository struct {
}

func (r memoryRepository) Get() ([]domain.Product, error) {
	return products, nil
}
