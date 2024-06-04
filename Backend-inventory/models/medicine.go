package models

type Medicine struct {
	ID                uint   `json:"id" gorm:"primary_key"`
	Nombre            string `json:"nombre"`
	IngredienteActivo string `json:"ingrediente_activo"`
	Marca             string `json:"marca"`
	FechaVencimiento  string `json:"fecha_vencimiento"`
	Cantidad          int    `json:"cantidad"`
}
