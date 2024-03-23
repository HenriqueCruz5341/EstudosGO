package repository

import (
	"api-and-messagerie/internal/entity"
	"database/sql"
)

type ProductRepositoryPostgreSql struct {
	DB *sql.DB
}

func NewProductRepositoryPostgreSql(db *sql.DB) *ProductRepositoryPostgreSql {
	return &ProductRepositoryPostgreSql{DB: db}
}

func (r *ProductRepositoryPostgreSql) Create(product *entity.Product) error {
	_, err := r.DB.Exec("INSERT INTO product (id, name, price) VALUES ($1, $2, $3)",
		product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}
	return nil
}

func (r *ProductRepositoryPostgreSql) FindAll() ([]*entity.Product, error) {
	rows, err := r.DB.Query("SELECT id, name, price FROM product")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entity.Product
	for rows.Next() {
		var product entity.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}
