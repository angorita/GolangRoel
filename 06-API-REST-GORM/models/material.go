package models

import (
	"GORM/db"
)
type Material struct {
	Descripcion string  `json:"descripcion"` // or "description" for English
	Precio      float64 `json:"precio"`      // or "price"
	Cantidad    int     `json:"cantidad"`    // or "quantity"
	Fecha       string  `json:"fecha"`       // or "date"
	Dolar       float64 `json:"dolar"`       // or "dollarValue"
	Id          int64   `json:"id"`
	Bhabilitado bool    `json:"habilitado"` // or "isEnabled", "enabled"
}
type Materiales []Material

func MigrarMaterial() {
	db.Database.AutoMigrate(Material{})
}
