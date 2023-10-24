package product

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterServiceProduct(router fiber.Router, db *gorm.DB) {
	repo := NewPostgresRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)
	var productRouter = router.Group("products")
	{
		productRouter.Post("", handler.CreateProduct)
	}
}
