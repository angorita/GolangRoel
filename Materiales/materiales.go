package main

type Material struct {
	Descripcion string  `json:"descripcion"`
	Precio      float64 `json:"precio"`
	Cantidad    int     `json:"cantidad"`
	Fecha       string  `json:"fecha"`
	Dolar       float64 `json:"dolar"`
	Id          int     `json:"id"`
	Bhabilitado bool    `json:"habilitado"`
}
