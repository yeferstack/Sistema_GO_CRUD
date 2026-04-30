package routes

import (
	"Sistema_GO_CRUD/controllers"

	"github.com/gorilla/mux"
)

// Registro de rutas para HistorialAdmin
func RegisterHistorialAdminRoutes(r *mux.Router) {
	r.HandleFunc("/historial-admin", controllers.GetAllHistorialAdmin).Methods("GET")
	r.HandleFunc("/historial-admin/{id}", controllers.GetHistorialAdminByID).Methods("GET")
	r.HandleFunc("/historial-admin", controllers.CreateHistorialAdmin).Methods("POST")
	r.HandleFunc("/historial-admin/{id}", controllers.UpdateHistorialAdmin).Methods("PUT")
	r.HandleFunc("/historial-admin/{id}", controllers.DeleteHistorialAdmin).Methods("DELETE")
}
