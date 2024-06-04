package models

import "gorm.io/gorm"

type Withdraw struct {
	gorm.Model
	MedicinaID      uint
	Cantidad        int
	Solicitante     string
	Estado          string
	Aprobador       string
	FechaAprobacion string
}
