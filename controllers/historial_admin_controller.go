package controllers

import (
	"Sistema_GO_CRUD/config"
	"Sistema_GO_CRUD/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GET ALL HistorialAdmin
func GetAllHistorialAdmin(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(`
		SELECT id_historial, id_admin, accion, descripcion, tipo_objeto, id_objeto, fecha_accion, activo, fecha_creacion, fecha_modificacion 
		FROM "Sistema"."HistorialAdmin"`)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()

	var list []models.HistorialAdmin
	for rows.Next() {
		var h models.HistorialAdmin
		rows.Scan(&h.ID_Historial, &h.ID_Admin, &h.Accion, &h.Descripcion, &h.TipoObjeto, &h.ID_Objeto, &h.FechaAccion, &h.Activo, &h.FechaCreacion, &h.FechaModificacion)
		list = append(list, h)
	}
	respondJSON(w, 200, list)
}

// GET BY ID HistorialAdmin
func GetHistorialAdminByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var h models.HistorialAdmin
	err := config.DB.QueryRow(`
		SELECT id_historial, id_admin, accion, descripcion, tipo_objeto, id_objeto, fecha_accion, activo, fecha_creacion, fecha_modificacion 
		FROM "Sistema"."HistorialAdmin" WHERE id_historial=$1`, id).
		Scan(&h.ID_Historial, &h.ID_Admin, &h.Accion, &h.Descripcion, &h.TipoObjeto, &h.ID_Objeto, &h.FechaAccion, &h.Activo, &h.FechaCreacion, &h.FechaModificacion)
	if err != nil {
		respondJSON(w, 404, map[string]string{"error": "Historial no encontrado"})
		return
	}
	respondJSON(w, 200, h)
}

// CREATE HistorialAdmin
func CreateHistorialAdmin(w http.ResponseWriter, r *http.Request) {
	var h models.HistorialAdmin
	if err := json.NewDecoder(r.Body).Decode(&h); err != nil {
		respondJSON(w, 400, map[string]string{"error": "JSON invalido"})
		return
	}
	err := config.DB.QueryRow(`
		INSERT INTO "Sistema"."HistorialAdmin" (id_admin, accion, descripcion, tipo_objeto, id_objeto, activo) 
		VALUES ($1,$2,$3,$4,$5,$6) RETURNING id_historial`,
		h.ID_Admin, h.Accion, h.Descripcion, h.TipoObjeto, h.ID_Objeto, h.Activo,
	).Scan(&h.ID_Historial)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 201, h)
}

// UPDATE HistorialAdmin
func UpdateHistorialAdmin(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var h models.HistorialAdmin
	json.NewDecoder(r.Body).Decode(&h)
	_, err := config.DB.Exec(`
		UPDATE "Sistema"."HistorialAdmin" SET accion=$1, descripcion=$2, tipo_objeto=$3, id_objeto=$4, activo=$5 
		WHERE id_historial=$6`,
		h.Accion, h.Descripcion, h.TipoObjeto, h.ID_Objeto, h.Activo, id,
	)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "Historial actualizado correctamente"})
}

// DELETE HistorialAdmin
func DeleteHistorialAdmin(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	_, err := config.DB.Exec(`DELETE FROM "Sistema"."HistorialAdmin" WHERE id_historial=$1`, id)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "Historial eliminado correctamente"})
}
