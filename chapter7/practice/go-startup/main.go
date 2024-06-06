package main

import (
	"golang-hactiv8-lab/chapter7/practice/go-startup/models"
	"golang-hactiv8-lab/chapter7/practice/go-startup/router"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "host=localhost user=gostartup password=mysecretpassword dbname=gostartup port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	userRepository := models.NewRepository(db)
	userService := models.NewService(userRepository)

	userHandler := router.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)

	router.Run()

}
