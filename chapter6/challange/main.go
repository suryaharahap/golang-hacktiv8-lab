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
	// Initialize database connection
	// dbConnection, err := db.NewPostgresDB()
	// if err != nil {
	// 	log.Fatalf("failed to connect to database: %v", err)
	// }

	// dbConnection, err := db.NewPostgresDB()
	dbConnection, err := database.NewPostgresDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Initialize repositories, services and handlers
	productRepo := product.NewGormProductRepository(dbConnection)
	productService := services.NewProductService(productRepo)
	productHandler := api.NewProductHandler(productService)

	// Initialize router
	r := router.InitRouter(productHandler)

	// Start server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
