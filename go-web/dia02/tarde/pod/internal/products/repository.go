package products

import "github.com/bgw6/pod/internal/domain"

func NewRepository() domain.Repository {
	return mysqlRepository{}
}
