package handlers

import (
	"05-api-rest/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// api para listar materiales
func GetMateriales(rw http.ResponseWriter, r *http.Request) {
	if materiales, err := models.ListarMateriales(); err != nil {
		models.SendNotFound(rw)
	} else {
		models.SendData(rw, materiales)
	}

}

// api para encontrar materiales
func GetMaterial(rw http.ResponseWriter, r *http.Request) {
	if material, err := getUserByRequest(r); err != nil {
		models.SendNotFound(rw)
	} else {
		models.SendData(rw, material)
	}
}

// api para crear
func CrearMaterial(rw http.ResponseWriter, r *http.Request) {
	material := models.Material{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&material); err != nil {
		models.SendUnprocessableEntity(rw)
	} else {
		material.Save()
		models.SendData(rw, material)
	}
}
func UpdateMaterial(rw http.ResponseWriter, r *http.Request) {
	//obtener id
	var materialId int64
	if material, err := getUserByRequest(r); err != nil {
		models.SendNotFound(rw)
	} else {
		materialId = material.Id
	}
	material := models.Material{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&material); err != nil {
		models.SendUnprocessableEntity(rw)
	} else {
		material.Id = materialId
		//le enviamos a save()una estructura ya construida
		//si tiene id la actualiza sino tiene id crea
		material.Save()
		models.SendData(rw, material)
	}
}
func DeleteMaterial(rw http.ResponseWriter, r *http.Request) {
	if material, err := getUserByRequest(r); err != nil {
		models.SendNotFound(rw)
	} else {
		material.Delete()
		models.SendData(rw, material)
	}
}
func getUserByRequest(r *http.Request) (models.Material, error) {
	//obtener id
	vars := mux.Vars(r)
	materialId, _ := strconv.Atoi(vars["id"])
	if material, err := models.GetMaterial(materialId); err != nil {
		return *material, err
	} else {
		return *material, nil
	}
}
