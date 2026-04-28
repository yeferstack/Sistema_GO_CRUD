package routes

import (
	"Sistema_GO_CRUD/controllers"

	"github.com/gorilla/mux"
)

// Registro de rutas para AdministradorPermiso
func RegisterAdminPermisoRoutes(r *mux.Router) {
	r.HandleFunc("/admin-permisos", controllers.GetAllAdminPermiso).Methods("GET")
	r.HandleFunc("/admin-permisos/{id}", controllers.GetAdminPermisoByID).Methods("GET")
	r.HandleFunc("/admin-permisos", controllers.CreateAdminPermiso).Methods("POST")
	r.HandleFunc("/admin-permisos/{id}", controllers.UpdateAdminPermiso).Methods("PUT")
	r.HandleFunc("/admin-permisos/{id}", controllers.DeleteAdminPermiso).Methods("DELETE")
}