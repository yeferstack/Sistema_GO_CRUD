package routes

import (
	"Sistema_GO_CRUD/controllers"

	"github.com/gorilla/mux"
)

// Registro de rutas para Nivel_Acceso
func RegisterNivelAccesoRoutes(r *mux.Router) {
	r.HandleFunc("/nivel-acceso", controllers.GetAllNivelAcceso).Methods("GET")
	r.HandleFunc("/nivel-acceso/{id}", controllers.GetNivelAccesoByID).Methods("GET")
	r.HandleFunc("/nivel-acceso", controllers.CreateNivelAcceso).Methods("POST")
	r.HandleFunc("/nivel-acceso/{id}", controllers.UpdateNivelAcceso).Methods("PUT")
	r.HandleFunc("/nivel-acceso/{id}", controllers.DeleteNivelAcceso).Methods("DELETE")
}
