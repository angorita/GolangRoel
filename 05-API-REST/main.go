package main

import (
	"05-api-rest/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//rutas o end point mux se asocian a rutas
	mux := mux.NewRouter()
	//creacion de end point obtener
	mux.HandleFunc("/api/material/", handlers.GetMateriales).Methods("GET")
	mux.HandleFunc("/api/material/{id:[0-9]+}", handlers.GetMaterial).Methods("GET")
	//post
	mux.HandleFunc("/api/material/", handlers.CrearMaterial).Methods("POST")
	mux.HandleFunc("/api/material/{id:[0-9]+}", handlers.UpdateMaterial).Methods("PUT")
	mux.HandleFunc("/api/material/{id:[0-9]+}", handlers.DeleteMaterial).Methods("DELETE")
	log.Println("Run server: http://localhost:8000/api/material/")
	log.Fatal(http.ListenAndServe(":8000", mux))

}
