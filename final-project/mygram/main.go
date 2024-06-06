package main

import (
	"golang-hactiv8-lab/final-project/mygram/database"
	"golang-hactiv8-lab/final-project/mygram/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	database.InitDB()

	r := routes.SetupRouter(database.DB)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8089"
	}
	r.Run(":" + port)
}
