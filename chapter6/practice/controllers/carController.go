package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Car struct {
	CarID string `json:"car_id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Price int    `json:"price"`
}

var CarDatas = []Car{}

func CreateCar(ctx *gin.Context) {
	var newCar Car

	if err := ctx.ShouldBindJSON(&newCar); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newCar.CarID = fmt.Sprintf("c%d", len(CarDatas)+1) // mengenerate id untuk tiap data mobil baru
	CarDatas = append(CarDatas, newCar)

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"car":     newCar,
	})
}

func UpdateCar(ctx *gin.Context) {
	carID := ctx.Param("CarID")
	fmt.Printf("Received carID: %s\n", carID) // Debug log
	condition := false
	var updateCar Car

	if carID == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_status":  "Bad Request",
			"error_message": "carID is required",
		})
		return
	}

	if err := ctx.ShouldBindJSON(&updateCar); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, car := range CarDatas {
		if carID == car.CarID {
			condition = true
			updateCar.CarID = car.CarID // Ensure CarID remains unchanged
			CarDatas[i] = updateCar
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("car with id %v not found", carID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": fmt.Sprintf("car with id %v has been successfully updated", carID),
		"car":     updateCar,
	})
}

func GetCar(ctx *gin.Context) {
	carID := ctx.Param("carID")
	condition := false
	var carData Car

	for i, car := range CarDatas {
		if carID == car.CarID {
			condition = true
			carData = CarDatas[i]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("car with id %v not found", carID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"car":     carData,
	})
}

func GetAllCar(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "success",
		"cars":    CarDatas,
	})
}

func DeleteCar(ctx *gin.Context) {
	carID := ctx.Param("carID")
	condition := false
	var carIndex int

	for i, car := range CarDatas {
		if carID == car.CarID {
			condition = true
			carIndex = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("car with id %v not found", carID),
		})
		return
	}

	copy(CarDatas[carIndex:], CarDatas[carIndex+1:])
	CarDatas[len(CarDatas)-1] = Car{}
	CarDatas = CarDatas[:len(CarDatas)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": fmt.Sprintf("car with id %v has been successfully deleted", carID),
	})
}
