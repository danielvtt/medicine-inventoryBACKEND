package controllers

import (
	"medicine-inventory/database"
	"medicine-inventory/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Contraseña), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo crear el usuario"})
		return
	}
	user.Contraseña = string(hashedPassword)

	database.DB.Create(&user)
	c.JSON(http.StatusOK, user)
}

func LoginUser(c *gin.Context) {
	var request models.User
	var user models.User

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Where("nombre_usuario = ?", request.NombreUsuario).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no encontrado"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Contraseña), []byte(request.Contraseña)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Contraseña incorrecta"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Inicio de sesión exitoso"})
}
