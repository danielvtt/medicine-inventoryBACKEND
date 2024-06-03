package routes

import (
	"medicine-inventory/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterMedicineRoutes(router *gin.Engine) {
	router.POST("/medicines", controllers.AddMedicine)
	router.PUT("/medicines/:id", controllers.EditMedicine)
	router.DELETE("/medicines/:id", controllers.DeleteMedicine)
	router.GET("/medicines", controllers.ListMedicines)
}
