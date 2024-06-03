package routes

import (
	"medicine-inventory/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine) {
	router.POST("/users", controllers.RegisterUser)
	router.POST("/login", controllers.LoginUser)
}
