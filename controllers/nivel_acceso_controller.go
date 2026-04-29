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

// CREATE NivelAcceso
func CreateNivelAcceso(w http.ResponseWriter, r *http.Request) {
	var n models.NivelAcceso
	if err := json.NewDecoder(r.Body).Decode(&n); err != nil {
		respondJSON(w, 400, map[string]string{"error": "JSON invalido"})
		return
	}
	err := config.DB.QueryRow(`
		INSERT INTO "Sistema"."Nivel_Acceso" (id_nivelacceso, moderador, administrador, activo) 
		VALUES ($1,$2,$3,$4) RETURNING id_admin`,
		n.ID_NivelAcceso, n.Moderador, n.Administrador, n.Activo,
	).Scan(&n.ID_Admin)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 201, n)
}

// UPDATE NivelAcceso
func UpdateNivelAcceso(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var n models.NivelAcceso
	json.NewDecoder(r.Body).Decode(&n)
	_, err := config.DB.Exec(`
		UPDATE "Sistema"."Nivel_Acceso" SET moderador=$1, administrador=$2, activo=$3 
		WHERE id_admin=$4`,
		n.Moderador, n.Administrador, n.Activo, id,
	)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "Nivel de acceso actualizado correctamente"})
}