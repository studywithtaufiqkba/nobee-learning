package handler

import (
	"github.com/gin-gonic/gin"
	"go_rest_api/service"
	"net/http"
)

type ProductHandler struct {
	service service.ProductService
}

func NewProductHandler(service service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (handler *ProductHandler) Route(r *gin.RouterGroup) {
	r.GET("/products", handler.GetAll)
}

func (handler *ProductHandler) GetAll(ctx *gin.Context) {
	data := handler.service.GetAll()
	ctx.JSON(http.StatusOK, data)
}
