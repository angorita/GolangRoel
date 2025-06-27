package u

import (
	"encoding/json"
	"fmt"
	"os"
)

func (b *Bal) CalcularTotal() float64 {
	return (b.Combo96 + b.HorasExtras + b.DiferenciaCaja + b.DiferenciaAguinaldo - b.DiferenciaBlanco)
}

type Bal struct {
	MesAnio             int     `json:"mesanio"`
	Combo96             float64 `json:"combo96"`
	HorasExtras         float64 `json:"horasextras"`
	DiferenciaCaja      float64 `json:"diferenciacaja"`
	DiferenciaAguinaldo float64 `json:"diferenciaaguinaldo"`
	DiferenciaBlanco    float64 `json:"diferenciablanco"`
	Total               float64 `json:"total"`
}

// Guarda contatos en un archivo json
func SaveBalanceToFile(balances []Bal) error {
	file, e := os.Create("balances.json")
	if e != nil {
		return e
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	e = encoder.Encode(balances)
	fmt.Println(encoder)
	if e != nil {
		return e
	}
	return nil
}

// Carga contactos desde un archivo json
func LoadBalanceFromFile(balances *[]Bal) error {
	file, err := os.Open("balances.json")
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&balances)
	if err != nil {
		return err
	}

	return nil //si anduvo bien
}
