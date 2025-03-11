package products

import "github.com/bgw6/pod/internal/domain"

type service struct {
	repository domain.Repository
}

func (s service) FindAll() ([]domain.Product, error) {
	return s.repository.FindAll()
}

func NewService(repo domain.Repository) domain.Service {
	return service{repo}
}
