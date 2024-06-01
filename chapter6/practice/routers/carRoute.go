package routers

import (
	"golang-hactiv8-lab/chapter6/practice/controllers"

	"github.com/gin-gonic/gin"
)

// MARK: function StartServer untuk menjalankan server dan mengembalikan suatu darta tipe data struct `*gin.Engine` yang berasal dari Gin .
func StartServer() *gin.Engine {
	// variabel `router` digunakan sebagai penampung untuk engine dari router kita dapatkan dari pemangilan function gin.Default .
	router := gin.Default()

	// Menghubungkan client dengan endpoint CreateCar, UpdateCar, GetCar and DeleteCar
	router.POST("/cars", controllers.CreateCar)
	router.PUT("/cars/:carID", controllers.UpdateCar)
	router.GET("/cars/:carID", controllers.GetCar)
	router.DELETE("/cars/:carID", controllers.DeleteCar)
	router.GET("/cars", controllers.GetAllCar)
	return router
}
