package products

import "github.com/bgw6/pod/internal/domain"

type service struct {
	repository domain.Repository
}

func (s service) FindAll() []domain.Product {
	return s.repository.FindAll()
}

func NewService() domain.Service {
	return service{}
}
