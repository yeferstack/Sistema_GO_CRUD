package routes

import (
	"Sistema_GO_CRUD/controllers"

	"github.com/gorilla/mux"
)

// Registro de rutas para Administrador
func RegisterAdministradorRoutes(r *mux.Router) {
	r.HandleFunc("/administradores", controllers.GetAllAdministrador).Methods("GET")
	r.HandleFunc("/administradores/{id}", controllers.GetAdministradorByID).Methods("GET")
	r.HandleFunc("/administradores", controllers.CreateAdministrador).Methods("POST")
	r.HandleFunc("/administradores/{id}", controllers.UpdateAdministrador).Methods("PUT")
	r.HandleFunc("/administradores/{id}", controllers.DeleteAdministrador).Methods("DELETE")
}
