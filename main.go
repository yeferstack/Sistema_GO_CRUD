package main

import (
	"log"
	"net/http"
	"os"

	"Sistema_GO_CRUD/config"
	"Sistema_GO_CRUD/routes"

	"github.com/gorilla/mux"
)

// middleware CORS
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// permiten cualquier origen de la peticion
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	config.ConnectDB() // conectar DB

	// RUTAS DE LA API
	r := mux.NewRouter()
	routes.RegisterNivelAccesoRoutes(r)    // registrar rutas de nivel de acceso
	routes.RegisterPermisoRoutes(r)        // registrar rutas de permiso
	routes.RegisterNotificacionRoutes(r)   // registrar rutas de notificacion
	routes.RegisterAdministradorRoutes(r)  // registrar rutas de administrador
	routes.RegisterAdminPermisoRoutes(r)   // registrar rutas de administrador-permiso
	routes.RegisterHistorialAdminRoutes(r) // registrar rutas de historial admin

	// leer puerto desde variable de entorno
	port := os.Getenv("INSCRIPCION_CRUD_HTTP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Servidor corriendo en el puerto", port)
	http.ListenAndServe(":"+port, enableCORS(r))
}
