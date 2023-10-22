package main

import (
	"github.com/gin-gonic/gin"
	"go_rest_api/controllers"
	"go_rest_api/handler"
	"go_rest_api/models"
	"go_rest_api/service"
)

func main() {
	r := gin.Default()
	db := models.ConnectDatabase()

	productController := controllers.NewProductController(db)
	productService := service.NewProductService(productController)
	handler.NewProductHandler(productService).Route(&r.RouterGroup)

	err := r.Run()
	if err != nil {
		return
	}

}
