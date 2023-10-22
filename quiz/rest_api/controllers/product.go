package controllers

import (
	"database/sql"
	"go_rest_api/models"
	"log"
)

type ProductController interface {
	GetAll() []models.Product
	Create(model models.Product) (*models.Product, error)
	Update(model models.Product) (*models.Product, error)
	Delete(model models.Product) error
}

type productController struct {
	db *sql.DB
}

func (p *productController) GetAll() []models.Product {
	rows, err := p.db.Query(`SELECT id, name, price FROM products`)
	if err != nil {
		return nil
	}

	var products []models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
			log.Fatal(err)
		}
		products = append(products, p)
	}
	return products
}

func (p *productController) Create(model models.Product) (*models.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p *productController) Update(model models.Product) (*models.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p *productController) Delete(model models.Product) error {
	//TODO implement me
	panic("implement me")
}

func NewProductController(db *sql.DB) ProductController {
	return &productController{db: db}
}
