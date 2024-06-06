package routes

import (
	"golang-hactiv8-lab/final-project/mygram/auth"
	"golang-hactiv8-lab/final-project/mygram/controllers"
	"golang-hactiv8-lab/final-project/mygram/handler"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {

	if db == nil {
		log.Fatalf("DB is nil")
	}
	userRepository := controllers.NewRepository(db)
	if userRepository == nil {
		log.Fatalf("userRepository is nil")
	}

	userService := controllers.NewService(userRepository)
	if userService == nil {
		log.Fatalf("userService is nil")
	}

	authServices := auth.NewService()
	if authServices == nil {
		log.Fatalf("authService is nil")
	}

	userHandler := handler.NewUserHandler(userService, authServices)
	if userHandler == nil {
		log.Fatalf("userHandler is nil")
	}

	r := gin.Default()
	api := r.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", userHandler.UploadAvatar)

	return r
}
