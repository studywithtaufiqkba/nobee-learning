package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var products []Product

func main() {
	router := gin.Default()

	// Get all products
	router.GET("/products", func(c *gin.Context) {
		c.JSON(http.StatusOK, products)
	})

	// Get a product by ID
	router.GET("/products/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		for _, p := range products {
			if p.ID == id {
				c.JSON(http.StatusOK, p)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
	})

	// Create a product
	router.POST("/products", func(c *gin.Context) {
		var newProduct Product
		if err := c.ShouldBindJSON(&newProduct); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		products = append(products, newProduct)
		c.JSON(http.StatusCreated, newProduct)
	})

	// Update a product by ID
	router.PUT("/products/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		var updatedProduct Product
		if err := c.ShouldBindJSON(&updatedProduct); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		for i, p := range products {
			if p.ID == id {
				products[i] = updatedProduct
				c.JSON(http.StatusOK, updatedProduct)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
	})

	// Delete a product by ID
	router.DELETE("/products/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		for i, p := range products {
			if id == p.ID {
				products = append(products[:i], products[i+1:]...)
				c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
	})

	// Run the server
	router.Run(":8080")
}
