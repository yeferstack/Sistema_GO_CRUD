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

// GET BY ID Administrador
func GetAdministradorByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var a models.Administrador
	err := config.DB.QueryRow(`
		SELECT id_admin, id_usuario, id_nivelacceso, activo, fecha_asignacion, fecha_creacion, fecha_modificacion 
		FROM "Sistema"."Administrador" WHERE id_admin=$1`, id).
		Scan(&a.ID_Admin, &a.ID_Usuario, &a.ID_NivelAcceso, &a.Activo, &a.FechaAsignacion, &a.FechaCreacion, &a.FechaModificacion)
	if err != nil {
		respondJSON(w, 404, map[string]string{"error": "Administrador no encontrado"})
		return
	}
	respondJSON(w, 200, a)
}

// CREATE Administrador
func CreateAdministrador(w http.ResponseWriter, r *http.Request) {
	var a models.Administrador
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		respondJSON(w, 400, map[string]string{"error": "JSON invalido"})
		return
	}
	err := config.DB.QueryRow(`
		INSERT INTO "Sistema"."Administrador" (id_usuario, id_nivelacceso, activo) 
		VALUES ($1,$2,$3) RETURNING id_admin`,
		a.ID_Usuario, a.ID_NivelAcceso, a.Activo,
	).Scan(&a.ID_Admin)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 201, a)
}

// UPDATE Administrador
func UpdateAdministrador(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var a models.Administrador
	json.NewDecoder(r.Body).Decode(&a)
	_, err := config.DB.Exec(`
		UPDATE "Sistema"."Administrador" SET id_nivelacceso=$1, activo=$2 
		WHERE id_admin=$3`,
		a.ID_NivelAcceso, a.Activo, id,
	)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "Administrador actualizado correctamente"})
}

// DELETE Administrador
func DeleteAdministrador(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	_, err := config.DB.Exec(`DELETE FROM "Sistema"."Administrador" WHERE id_admin=$1`, id)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "Administrador eliminado correctamente"})
}
