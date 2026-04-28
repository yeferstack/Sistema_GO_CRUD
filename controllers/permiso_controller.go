package controllers

import (
	"Sistema_GO_CRUD/config"
	"Sistema_GO_CRUD/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GET ALL Permiso
func GetAllPermiso(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(`
		SELECT id_permiso, nombre, descripcion, activo, fecha_creacion, fecha_modificacion 
		FROM "Sistema"."Permiso"`)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()

	var list []models.Permiso
	for rows.Next() {
		var p models.Permiso
		rows.Scan(&p.ID_Permiso, &p.Nombre, &p.Descripcion, &p.Activo, &p.FechaCreacion, &p.FechaModificacion)
		list = append(list, p)
	}
	respondJSON(w, 200, list)
}

// GET BY ID Permiso
func GetPermisoByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var p models.Permiso
	err := config.DB.QueryRow(`
		SELECT id_permiso, nombre, descripcion, activo, fecha_creacion, fecha_modificacion 
		FROM "Sistema"."Permiso" WHERE id_permiso=$1`, id).
		Scan(&p.ID_Permiso, &p.Nombre, &p.Descripcion, &p.Activo, &p.FechaCreacion, &p.FechaModificacion)
	if err != nil {
		respondJSON(w, 404, map[string]string{"error": "Permiso no encontrado"})
		return
	}
	respondJSON(w, 200, p)
}

// CREATE Permiso
func CreatePermiso(w http.ResponseWriter, r *http.Request) {
	var p models.Permiso
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		respondJSON(w, 400, map[string]string{"error": "JSON invalido"})
		return
	}
	err := config.DB.QueryRow(`
		INSERT INTO "Sistema"."Permiso" (nombre, descripcion, activo) 
		VALUES ($1,$2,$3) RETURNING id_permiso`,
		p.Nombre, p.Descripcion, p.Activo,
	).Scan(&p.ID_Permiso)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 201, p)
}

// UPDATE Permiso
func UpdatePermiso(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var p models.Permiso
	json.NewDecoder(r.Body).Decode(&p)
	_, err := config.DB.Exec(`
		UPDATE "Sistema"."Permiso" SET nombre=$1, descripcion=$2, activo=$3 
		WHERE id_permiso=$4`,
		p.Nombre, p.Descripcion, p.Activo, id,
	)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "Permiso actualizado correctamente"})
}

// DELETE Permiso
func DeletePermiso(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	_, err := config.DB.Exec(`DELETE FROM "Sistema"."Permiso" WHERE id_permiso=$1`, id)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "Permiso eliminado correctamente"})
}

// GET BY ID Permiso
func GetPermisoByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var p models.Permiso
	err := config.DB.QueryRow(`
		SELECT id_permiso, nombre, descripcion, activo, fecha_creacion, fecha_modificacion 
		FROM "Sistema"."Permiso" WHERE id_permiso=$1`, id).
		Scan(&p.ID_Permiso, &p.Nombre, &p.Descripcion, &p.Activo, &p.FechaCreacion, &p.FechaModificacion)
	if err != nil {
		respondJSON(w, 404, map[string]string{"error": "Permiso no encontrado"})
		return
	}
	respondJSON(w, 200, p)
}

// CREATE Permiso
func CreatePermiso(w http.ResponseWriter, r *http.Request) {
	var p models.Permiso
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		respondJSON(w, 400, map[string]string{"error": "JSON invalido"})
		return
	}
	err := config.DB.QueryRow(`
		INSERT INTO "Sistema"."Permiso" (nombre, descripcion, activo) 
		VALUES ($1,$2,$3) RETURNING id_permiso`,
		p.Nombre, p.Descripcion, p.Activo,
	).Scan(&p.ID_Permiso)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 201, p)
}

// UPDATE Permiso
func UpdatePermiso(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var p models.Permiso
	json.NewDecoder(r.Body).Decode(&p)
	_, err := config.DB.Exec(`
		UPDATE "Sistema"."Permiso" SET nombre=$1, descripcion=$2, activo=$3 
		WHERE id_permiso=$4`,
		p.Nombre, p.Descripcion, p.Activo, id,
	)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "Permiso actualizado correctamente"})
}

// DELETE Permiso
func DeletePermiso(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	_, err := config.DB.Exec(`DELETE FROM "Sistema"."Permiso" WHERE id_permiso=$1`, id)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "Permiso eliminado correctamente"})
}
