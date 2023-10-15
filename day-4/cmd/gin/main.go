package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.New()

	router.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]string{
			"status": "success",
		})
	})

	router.Run()
}
