package service

import (
	"go_rest_api/controllers"
	"go_rest_api/dto"
	"go_rest_api/models"
)

type ProductService interface {
	GetAll() []models.Product
	Create(dto dto.ProductDto) (*models.Product, error)
	Update(id int, dto dto.ProductDto) (*models.Product, error)
	Delete(id int) error
}

type productService struct {
	controller controllers.ProductController
}

func (p *productService) GetAll() []models.Product {
	return p.controller.GetAll()
}

func (p *productService) Create(dto dto.ProductDto) (*models.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p *productService) Update(id int, dto dto.ProductDto) (*models.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p *productService) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}

func NewProductService(controller controllers.ProductController) ProductService {
	return &productService{controller: controller}
}
