package models

import "gorm.io/gorm"

type Medicine struct {
	gorm.Model
	Nombre            string
	IngredienteActivo string
	Marca             string
	FechaVencimiento  string
	Cantidad          int
}
