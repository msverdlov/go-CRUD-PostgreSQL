package routes

import (
	"anyData/internal/app/anyData/controller"
	"github.com/gin-gonic/gin"
)

const proxy = "192.168.0.1"

func InitRoutes() (router *gin.Engine) {
	router = gin.Default()
	err := router.SetTrustedProxies([]string{proxy})
	if err != nil {
		return nil
	}

	router.POST("/add-data", controller.AddData)

	router.GET("/get-content", controller.GetContent)
	router.GET("/get-data", controller.GetData)

	router.PATCH("/update-data", controller.UpdateData)
	router.DELETE("/delete-data", controller.DeleteData)

	return router
}
