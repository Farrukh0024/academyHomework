package postgres

import (
	"database/sql"
	"fmt"
	
	"test/api/models"
	"test/storage"
)

type repo struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) storage.IRepository {
	return &repo{db: db}
}

func (p *repo) Search(productName string, quantity uint) (models.ProductSell, error) {
	product := models.ProductSell{}

	query := `select name, quantity, price, original_price from products where name = $1`
	if err := p.db.QueryRow(query, productName).Scan(
		&product.Name,
		&product.Quantity,
		&product.Price,
		&product.OriginalPrice); err != nil {
		if product.Name != productName {
			fmt.Println("we do not have this product")
			return models.ProductSell{}, err
		} else {
			fmt.Println("error is while searching product", err.Error())
		}
		return models.ProductSell{}, err
	}

	if product.Quantity < quantity {
		fmt.Println("we do not have enough product")
		return models.ProductSell{}, nil
	}

	soldQuantity := product.Quantity - quantity
	if product.Quantity >= quantity {
		p.TakeProduct(productName, soldQuantity)
	}

	return models.ProductSell{
		Name:          product.Name,
		Price:         product.Price,
		OriginalPrice: product.OriginalPrice,
		Quantity:      soldQuantity,
	}, nil
}

func (p *repo) TakeProduct(productName string, quantity uint) {
	query := `update products set quantity = $1 where name = $2`
	if _, err := p.db.Exec(query, quantity, productName); err != nil {
		fmt.Println("error is while taking product from database", err.Error())
		return
	}
}
