package domain

type Product struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

type Repository interface {
	// Get returns all the products
	Get() ([]Product, error)
	// // GetByID returns a product by id
	// GetByID(id int) (*Product, error)
	// // Save saves a product
	Create(p *Product) error
	Update(p *Product) error
	Delete(productId int) error
}

type Service interface {
	Get() ([]Product, error)
	// // GetByID returns a product by id
	// GetByID(id int) (*Product, error)
	// // Save saves a product
	Create(p *Product) error
	Update(p *Product) error
	Delete(productId int) error
}
