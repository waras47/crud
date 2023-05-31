package routers

import (
	"crud_test/controllers"

	"github.com/gin-gonic/gin"
)

func StarttServer() *gin.Engine {
	router := gin.Default()

	router.GET("/cars", controllers.GetAllCar)
	router.POST("/cars", controllers.CreateCar)
	router.GET("/cars/:carID", controllers.GetByIdCar)
	router.PUT("/cars/:carID", controllers.UpdateCar)
	router.DELETE("/cars/:carID", controllers.DeleteCar)

	return router
}