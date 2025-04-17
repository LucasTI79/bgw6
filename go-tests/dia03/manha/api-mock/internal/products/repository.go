package products

import (
	"database/sql"

	"github.com/bgw6/pod/internal/domain"
)

func NewRepository(database *sql.DB) domain.Repository {
	return &mysqlRepository{
		db: database,
	}
	// return &mapRepository{
	// 	products: make(map[int]ProductAttributes),
	// }
}
