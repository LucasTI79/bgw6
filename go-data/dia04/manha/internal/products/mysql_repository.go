package products

import (
	"database/sql"
	"errors"

	"github.com/bgw6/pod/internal/domain"
)

type mysqlRepository struct {
	db *sql.DB
}

func (r mysqlRepository) Get() ([]domain.Product, error) {
	var products []domain.Product
	rows, err := r.db.Query(GetQuery)

	if err != nil {
		return products, err
	}

	for rows.Next() {
		var product domain.Product
		if err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity, &product.Type); err != nil {
			return products, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (r *mysqlRepository) GetByID(id int) (*domain.Product, error) {
	// execute query
	row := r.db.QueryRow(GetOneQuery, id)
	if err := row.Err(); err != nil {
		return &domain.Product{}, err
	}
	// scan result
	var product domain.Product
	if err := row.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity, &product.Type); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &domain.Product{}, domain.ErrResourceNotFound
		}
		return &domain.Product{}, err
	}

	return &product, nil
}

func (r *mysqlRepository) Create(p *domain.Product) error {
	// query execution
	result, err := r.db.Exec(
		StoreQuery,
		(*p).Name, (*p).Price, (*p).Quantity, (*p).Type,
	)
	if err != nil {
		return err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return err
	}

	(*p).Id = int(lastInsertId)
	return nil
}

func (r *mysqlRepository) Update(p *domain.Product) error {
	// query execution
	_, err := r.db.Exec(
		UpdateQuery,
		(*p).Name, (*p).Price, (*p).Quantity, (*p).Type, (*p).Id,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *mysqlRepository) Delete(id int) error {
	// query execution
	_, err := r.db.Exec(DeleteQuery, id)
	if err != nil {
		return err
	}
	return nil
}
