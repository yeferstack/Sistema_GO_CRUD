package controllers

import (
	"Sistema_GO_CRUD/config"
	"Sistema_GO_CRUD/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GET ALL NivelAcceso
func GetAllNivelAcceso(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(`
		SELECT id_admin, id_nivelacceso, moderador, administrador, activo, fecha_asignacion, fecha_creacion, fecha_modificacion 
		FROM "Sistema"."Nivel_Acceso"`)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()

	var list []models.NivelAcceso
	for rows.Next() {
		var n models.NivelAcceso
		rows.Scan(&n.ID_Admin, &n.ID_NivelAcceso, &n.Moderador, &n.Administrador, &n.Activo, &n.FechaAsignacion, &n.FechaCreacion, &n.FechaModificacion)
		list = append(list, n)
	}
	respondJSON(w, 200, list)
}

// GET BY ID NivelAcceso
func GetNivelAccesoByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var n models.NivelAcceso
	err := config.DB.QueryRow(`
		SELECT id_admin, id_nivelacceso, moderador, administrador, activo, fecha_asignacion, fecha_creacion, fecha_modificacion 
		FROM "Sistema"."Nivel_Acceso" WHERE id_admin=$1`, id).
		Scan(&n.ID_Admin, &n.ID_NivelAcceso, &n.Moderador, &n.Administrador, &n.Activo, &n.FechaAsignacion, &n.FechaCreacion, &n.FechaModificacion)
	if err != nil {
		respondJSON(w, 404, map[string]string{"error": "Nivel de acceso no encontrado"})
		return
	}
	respondJSON(w, 200, n)
}