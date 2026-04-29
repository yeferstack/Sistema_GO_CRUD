package controllers

import (
	"Sistema_GO_CRUD/config"
	"Sistema_GO_CRUD/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GET ALL AdministradorPermiso
func GetAllAdminPermiso(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(`
		SELECT id, id_admin, id_permiso, activo, fecha_creacion, fecha_modificacion 
		FROM "Sistema"."AdministradorPermiso"`)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()

	var list []models.AdministradorPermiso
	for rows.Next() {
		var ap models.AdministradorPermiso
		rows.Scan(&ap.ID, &ap.ID_Admin, &ap.ID_Permiso, &ap.Activo, &ap.FechaCreacion, &ap.FechaModificacion)
		list = append(list, ap)
	}
	respondJSON(w, 200, list)
}