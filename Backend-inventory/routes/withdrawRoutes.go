package routes

import (
	"medicine-inventory/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterWithdrawRoutes(router *gin.Engine) {
	router.POST("/withdraws", controllers.RequestWithdraw)
	router.PUT("/withdraws/:id/approve", controllers.ApproveWithdraw)
	router.PUT("/withdraws/:id/deny", controllers.DenyWithdraw)
	router.GET("/withdraws", controllers.ListWithdraws)
}
