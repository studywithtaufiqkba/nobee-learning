package main

import (
	"day_6/app/product"
	"day_6/config"
	"day_6/pkg/database"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	router := fiber.New(fiber.Config{
		AppName: "Products Services",
		Prefork: true,
	})

	err := config.LoafConfig("./config/config.yaml")
	if err != nil {
		log.Println("error when try to LoadConfig with error: ", err.Error())
	}

	db, err := database.ConnectGORMMySQL(config.Cfg.DB)
	if db != nil {
		log.Println("db connected")
	}

	product.RegisterServiceProduct(router, db)

	err = router.Listen(config.Cfg.App.Port)
	if err != nil {
		return
	}
}
