package routes

import (
	"Sistema_GO_CRUD/controllers"

	"github.com/gorilla/mux"
)

// Registro de rutas para Notificacion
func RegisterNotificacionRoutes(r *mux.Router) {
	r.HandleFunc("/notificaciones", controllers.GetAllNotificacion).Methods("GET")
	r.HandleFunc("/notificaciones/{id}", controllers.GetNotificacionByID).Methods("GET")
	r.HandleFunc("/notificaciones", controllers.CreateNotificacion).Methods("POST")
	r.HandleFunc("/notificaciones/{id}", controllers.UpdateNotificacion).Methods("PUT")
	r.HandleFunc("/notificaciones/{id}", controllers.DeleteNotificacion).Methods("DELETE")
}
