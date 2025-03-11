package domain

type Product struct {
}

type Repository interface {
	FindAll() []Product
}

type Service interface {
	FindAll() []Product
}
