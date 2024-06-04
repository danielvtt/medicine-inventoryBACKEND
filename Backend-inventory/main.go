package main

import (
	"medicine-inventory/database"
	"medicine-inventory/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Configurar CORS
	r.Use(cors.Default())

	database.Connect()
	database.AutoMigrate()

	// Registrar rutas
	routes.RegisterMedicineRoutes(r)
	routes.RegisterUserRoutes(r)
	routes.RegisterWithdrawRoutes(r)

	r.Run() // por defecto en el puerto 8080
}
