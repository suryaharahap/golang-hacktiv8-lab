package main

import (
	"golang-hactiv8-lab/chapter6/challange/domain/product"
	"golang-hactiv8-lab/chapter6/challange/infrastructure/database"
	"golang-hactiv8-lab/chapter6/challange/infrastructure/router"
	"golang-hactiv8-lab/chapter6/challange/interfaces/api"
	"golang-hactiv8-lab/chapter6/challange/services"
	"log"
)

func main() {
	dbConnection, err := database.NewPostgresDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	productRepo := product.NewGormProductRepository(dbConnection)
	productService := services.NewProductService(productRepo)
	productHandler := api.NewProductHandler(productService)

	r := router.InitRouter(productHandler)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
