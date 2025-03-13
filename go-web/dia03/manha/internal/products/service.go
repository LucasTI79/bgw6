package products

import "github.com/bgw6/pod/internal/domain"

type service struct {
	repository domain.Repository
}

func (s service) Get() ([]domain.Product, error) {
	return s.repository.Get()
}

func (s service) GetByID(productId int) (*domain.Product, error) {
	return s.repository.GetByID(productId)
}

func (s service) Update(p *domain.Product) error {
	return s.repository.Update(p)
}

func (s service) Create(p *domain.Product) error {
	return s.repository.Create(p)
}

func (s service) Delete(productId int) error {
	return s.repository.Delete(productId)
}

func NewService(repo domain.Repository) domain.Service {
	return service{repo}
}
