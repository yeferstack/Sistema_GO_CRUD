package controllers

import (
	"Sistema_GO_CRUD/config"
	"Sistema_GO_CRUD/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GET ALL Administrador
func GetAllAdministrador(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(`
		SELECT id_admin, id_usuario, id_nivelacceso, activo, fecha_asignacion, fecha_creacion, fecha_modificacion 
		FROM "Sistema"."Administrador"`)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()

	var list []models.Administrador
	for rows.Next() {
		var a models.Administrador
		rows.Scan(&a.ID_Admin, &a.ID_Usuario, &a.ID_NivelAcceso, &a.Activo, &a.FechaAsignacion, &a.FechaCreacion, &a.FechaModificacion)
		list = append(list, a)
	}
	respondJSON(w, 200, list)
}

