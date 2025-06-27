package main

import (
	"encoding/json"
	"os"
)

func SaveToFileMaterial(materiales []Material) error {
	file, err := os.Create("materiales.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(materiales); err != nil {
		return err
	}
	return nil
}

func LoadFromMateriales(materiales *[]Material) error {
	file, err := os.Open("materiales.json")
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(materiales) // Cambiado
	if err != nil {
		return err
	}
	return nil
}
