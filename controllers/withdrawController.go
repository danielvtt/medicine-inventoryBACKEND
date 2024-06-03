package controllers

import (
	"medicine-inventory/database"
	"medicine-inventory/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequestWithdraw(c *gin.Context) {
	var withdraw models.Withdraw
	if err := c.ShouldBindJSON(&withdraw); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	withdraw.Estado = "pendiente"
	database.DB.Create(&withdraw)
	c.JSON(http.StatusOK, withdraw)
}

func ApproveWithdraw(c *gin.Context) {
	var withdraw models.Withdraw
	id := c.Param("id")
	if err := database.DB.First(&withdraw, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Solicitud de retiro no encontrada"})
		return
	}
	withdraw.Estado = "aprobado"
	database.DB.Save(&withdraw)
	c.JSON(http.StatusOK, withdraw)
}

func DenyWithdraw(c *gin.Context) {
	var withdraw models.Withdraw
	id := c.Param("id")
	if err := database.DB.First(&withdraw, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Solicitud de retiro no encontrada"})
		return
	}
	withdraw.Estado = "denegado"
	database.DB.Save(&withdraw)
	c.JSON(http.StatusOK, withdraw)
}

func ListWithdraws(c *gin.Context) {
	var withdraws []models.Withdraw
	database.DB.Find(&withdraws)
	c.JSON(http.StatusOK, withdraws)
}
