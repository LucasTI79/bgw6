package domain

type Product struct {
}

type Repository interface {
	FindAll() ([]Product, error)
}

type Service interface {
	FindAll() ([]Product, error)
}
