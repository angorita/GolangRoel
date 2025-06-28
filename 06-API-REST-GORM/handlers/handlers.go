package handlers

import (
	"GORM/db"
	"GORM/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// api para listar materiales
func GetMateriales(rw http.ResponseWriter, r *http.Request) {
	materiales := models.Materiales{}
	db.Database.Find(&materiales)
	sendData(rw, materiales, http.StatusOK)
}

// api para encontrar materiales
func GetMaterial(rw http.ResponseWriter, r *http.Request) {
	if material, err := getMaterialById(r); err != nil {
		sendError(rw, http.StatusNotFound)
	} else {
		sendData(rw, material, http.StatusOK)
	}

}

func getMaterialById(r *http.Request) (models.Material, *gorm.DB) {
	//obtener id
	vars := mux.Vars(r)
	materialId, _ := strconv.Atoi(vars["id"])
	material := models.Material{}
	if err := db.Database.First(&material, materialId); err != nil {
		return material, err
	} else {
		return material, nil
	}
}
func CreateMaterial(rw http.ResponseWriter, r *http.Request) {
	//obtener registro
	material := models.Material{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&material); err != nil {
		sendError(rw, http.StatusUnprocessableEntity)
	} else {
		db.Database.Save(&material)
		sendData(rw, material, http.StatusCreated)
	}

}
func UpdateMaterial(rw http.ResponseWriter, r *http.Request) {
	//obtener registro
	var materialId int64
	if material_ant, err := getMaterialById(r); err != nil {
		sendError(rw, http.StatusNotFound)

	} else {
		materialId = material_ant.Id
		material := models.Material{}
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&material); err != nil {
			sendError(rw, http.StatusUnprocessableEntity)
		} else {
			material.Id = materialId
			db.Database.Save(&material)
			sendData(rw, material, http.StatusOK)
		}
	}
}
func Delete(rw http.ResponseWriter, r *http.Request) {
	if material, err := getMaterialById(r); err != nil {
		sendError(rw, http.StatusNotFound)
	} else {
		db.Database.Delete(&material)
		sendData(rw, material, http.StatusOK)
	}
}
