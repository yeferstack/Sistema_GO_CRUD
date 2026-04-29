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

// GET BY ID AdministradorPermiso
func GetAdminPermisoByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var ap models.AdministradorPermiso
	err := config.DB.QueryRow(`
		SELECT id, id_admin, id_permiso, activo, fecha_creacion, fecha_modificacion 
		FROM "Sistema"."AdministradorPermiso" WHERE id=$1`, id).
		Scan(&ap.ID, &ap.ID_Admin, &ap.ID_Permiso, &ap.Activo, &ap.FechaCreacion, &ap.FechaModificacion)
	if err != nil {
		respondJSON(w, 404, map[string]string{"error": "Registro no encontrado"})
		return
	}
	respondJSON(w, 200, ap)
}

// CREATE AdministradorPermiso
func CreateAdminPermiso(w http.ResponseWriter, r *http.Request) {
	var ap models.AdministradorPermiso
	if err := json.NewDecoder(r.Body).Decode(&ap); err != nil {
		respondJSON(w, 400, map[string]string{"error": "JSON invalido"})
		return
	}
	err := config.DB.QueryRow(`
		INSERT INTO "Sistema"."AdministradorPermiso" (id_admin, id_permiso, activo) 
		VALUES ($1,$2,$3) RETURNING id, fecha_creacion, fecha_modificacion`,
		ap.ID_Admin, ap.ID_Permiso, ap.Activo,
	).Scan(&ap.ID, &ap.FechaCreacion, &ap.FechaModificacion)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 201, ap)
}

// UPDATE AdministradorPermiso
func UpdateAdminPermiso(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var ap models.AdministradorPermiso
	json.NewDecoder(r.Body).Decode(&ap)
	_, err := config.DB.Exec(`
		UPDATE "Sistema"."AdministradorPermiso" SET id_admin=$1, id_permiso=$2, activo=$3 
		WHERE id=$4`,
		ap.ID_Admin, ap.ID_Permiso, ap.Activo, id,
	)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "AdministradorPermiso actualizado correctamente"})
}

// DELETE AdministradorPermiso
func DeleteAdminPermiso(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	_, err := config.DB.Exec(`DELETE FROM "Sistema"."AdministradorPermiso" WHERE id=$1`, id)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "AdministradorPermiso eliminado correctamente"})
}