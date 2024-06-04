package controllers

import (
	"medicine-inventory/database"
	"medicine-inventory/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddMedicine(c *gin.Context) {
	var medicine models.Medicine
	if err := c.ShouldBindJSON(&medicine); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&medicine)
	c.JSON(http.StatusOK, medicine)
}

func EditMedicine(c *gin.Context) {
	var medicine models.Medicine
	id := c.Param("id")
	if err := database.DB.First(&medicine, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Medicina no encontrada"})
		return
	}
	if err := c.ShouldBindJSON(&medicine); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&medicine)
	c.JSON(http.StatusOK, medicine)
}

func DeleteMedicine(c *gin.Context) {
	var medicine models.Medicine
	id := c.Param("id")
	if err := database.DB.First(&medicine, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Medicina no encontrada"})
		return
	}
	database.DB.Delete(&medicine)
	c.JSON(http.StatusOK, gin.H{"message": "Medicina eliminada"})
}

func ListMedicines(c *gin.Context) {
	var medicines []models.Medicine
	database.DB.Find(&medicines)
	c.JSON(http.StatusOK, medicines)
}
