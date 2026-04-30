package routes

import (
	"Sistema_GO_CRUD/controllers"

	"github.com/gorilla/mux"
)

// Registro de rutas para Permiso
func RegisterPermisoRoutes(r *mux.Router) {
	r.HandleFunc("/permisos", controllers.GetAllPermiso).Methods("GET")
	r.HandleFunc("/permisos/{id}", controllers.GetPermisoByID).Methods("GET")
	r.HandleFunc("/permisos", controllers.CreatePermiso).Methods("POST")
	r.HandleFunc("/permisos/{id}", controllers.UpdatePermiso).Methods("PUT")
	r.HandleFunc("/permisos/{id}", controllers.DeletePermiso).Methods("DELETE")
}
