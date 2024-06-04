package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	NombreUsuario string
	Contraseña    string
	Rol           string
}
