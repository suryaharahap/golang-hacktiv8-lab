package router

import (
	"golang-hactiv8-lab/chapter6/challange/interfaces/api"

	"github.com/gin-gonic/gin"
)

func InitRouter(productHandler *api.ProductHandler) *gin.Engine {
	r := gin.Default()

	r.GET("/products", productHandler.GetAll)
	r.GET("/products/:id", productHandler.GetByID)
	r.POST("/products", productHandler.Create)
	r.PUT("/products/:id", productHandler.Update)
	r.DELETE("/products/:id", productHandler.Delete)

	return r
}
