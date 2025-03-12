package products

import (
	"database/sql"

	"github.com/bgw6/pod/internal/domain"
)

type mysqlRepository struct {
	db *sql.DB
}

func (r mysqlRepository) Get() []domain.Product {
	// r.db.Query("SELECT * FROM products")

	return nil
}
