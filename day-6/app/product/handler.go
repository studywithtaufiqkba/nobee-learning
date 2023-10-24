package product

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type Handler struct {
	svc Service
}

func NewHandler(svc Service) Handler {
	return Handler{svc: svc}
}

func (h Handler) CreateProduct(c *fiber.Ctx) error {
	var req = CreateProductRequest{}

	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "BAD REQUEST",
			"error":   err.Error(),
		})
	}
	model := Product{
		Name:     req.Name,
		Category: req.Category,
		Price:    req.Price,
		Stock:    req.Stock,
	}

	err = h.svc.CreateProduct(c.UserContext(), model)
	if err != nil {
		payload := fiber.Map{}
		httpCode := 400
		switch {
		case errors.Is(err, ErrEmptyName), errors.Is(err, ErrEmptyCategory), errors.Is(err, ErrEmptyPrice), errors.Is(err, ErrEmptyStock):
			payload = fiber.Map{
				"success": false,
				"message": "BAD REQUEST",
				"error":   err.Error(),
			}
		default:
			payload = fiber.Map{
				"success": false,
				"message": "ERR INTERNAL",
				"error":   err.Error(),
			}
		}
		return c.Status(httpCode).JSON(payload)
	}

	return c.Status(http.StatusBadRequest).JSON(fiber.Map{
		"success": true,
		"message": "CREATE SUCCESS",
	})
}
