package contactos

import (
	"encoding/json"
	"os"
)

type Contacto struct {
	Nom  string
	Ape  string
	Edad int
}

/* carga los contactos desde un archivo json */
func LoadContactosFromFile(contactos *[]Contacto) error {
	file, err := os.Open("contactos.json")
	if err != nil {
		return err

	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&contactos)
	if err != nil {
		return err
	}
	return nil
}
