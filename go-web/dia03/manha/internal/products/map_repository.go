package products

import (
	"fmt"

	"github.com/bgw6/pod/internal/domain"
)

type ProductAttributes struct {
	Name     string
	Type     string
	Quantity int
	Price    float64
}

type mapRepository struct {
	lastId   int
	products map[int]ProductAttributes
}

func (pr ProductAttributes) ToDomain() domain.Product {
	return domain.Product{
		Id:       0,
		Name:     pr.Name,
		Type:     pr.Type,
		Quantity: pr.Quantity,
		Price:    pr.Price,
	}
}

func (r mapRepository) Get() ([]domain.Product, error) {
	var pr []domain.Product

	for id, value := range r.products {
		productDomain := value.ToDomain()
		productDomain.Id = id
		pr = append(pr, productDomain)
	}

	return pr, nil
}

func (r *mapRepository) Update(p *domain.Product) (err error) {
	attr := ProductAttributes{
		Name:     p.Name,
		Type:     p.Type,
		Quantity: p.Quantity,
		Price:    p.Price,
	}
	// update
	_, ok := r.products[p.Id]

	if !ok {
		return domain.ErrResourceNotFound
	}

	r.products[p.Id] = attr
	return nil
}

func (r *mapRepository) Create(p *domain.Product) (err error) {
	r.lastId++
	p.Id = r.lastId
	attr := ProductAttributes{
		Name:     p.Name,
		Type:     p.Type,
		Quantity: p.Quantity,
		Price:    p.Price,
	}

	r.products[p.Id] = attr
	return nil
}

func (r *mapRepository) Delete(productId int) error {
	if _, ok := r.products[productId]; !ok {
		return fmt.Errorf("%w: product with %d not found", domain.ErrResourceNotFound, productId)
	}
	// delete
	delete(r.products, productId)
	return nil
}
