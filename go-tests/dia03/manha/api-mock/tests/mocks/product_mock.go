package mocks

import (
	"github.com/bgw6/pod/internal/domain"
	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	mock.Mock
}

func (r *ProductRepositoryMock) Get() ([]domain.Product, error) {
	args := r.Called()

	arg0, ok := args.Get(0).([]domain.Product)

	if ok {
		return arg0, args.Error(1)
	}

	return []domain.Product{}, args.Error(1)
}

func (r *ProductRepositoryMock) GetByID(id int) (*domain.Product, error) {
	args := r.Called(id)
	return args.Get(0).(*domain.Product), args.Error(1)
}

func (r *ProductRepositoryMock) Create(p *domain.Product) error {
	args := r.Called(p)
	return args.Error(0)
}

func (r *ProductRepositoryMock) Update(p *domain.Product) error {
	args := r.Called(p)
	return args.Error(0)
}

func (r *ProductRepositoryMock) Delete(id int) error {
	args := r.Called(id)
	return args.Error(0)
}
