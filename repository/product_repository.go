package repository

import (
	"database/sql"
	"fmt"
	"rest-api-go/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pRepository *ProductRepository) GetAllProducts() ([]model.Product, error) {

	query := "SELECT id, name, price FROM product"
	rows, err := pRepository.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}

	var productList []model.Product
	var productObj model.Product

	for rows.Next() {
		err = rows.Scan(
			&productObj.ID,
			&productObj.Name,
			&productObj.Price,
		)
		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}

		productList = append(productList, productObj)
	}
	rows.Close()

	return productList, nil
}

func (pRepository *ProductRepository) SaveProduct(product model.Product) (int, error) {
	query, err := pRepository.connection.Prepare("INSERT INTO product(name, price) VALUES ($1, $2) RETURNING id")
	if err != nil {
		return -1, err
	}

	var id int
	err = query.QueryRow(product.Name, product.Price).Scan(&id)
	if err != nil {
		return -1, err
	}
	query.Close()

	return id, nil
}
