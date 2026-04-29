package controllers

import (
	"Sistema_GO_CRUD/config"
	"Sistema_GO_CRUD/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GET ALL Notificacion
func GetAllNotificacion(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(`
		SELECT id_notificacion, id_usuario, titulo, mensaje, tipo, id_referencia, tipo_referencia, leido, activo, fecha_creacion, fecha_modificacion 
		FROM "Sistema"."Notificacion"`)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()

	var list []models.Notificacion
	for rows.Next() {
		var n models.Notificacion
		rows.Scan(&n.ID_Notificacion, &n.ID_Usuario, &n.Titulo, &n.Mensaje, &n.Tipo, &n.ID_Referencia, &n.TipoReferencia, &n.Leido, &n.Activo, &n.FechaCreacion, &n.FechaModificacion)
		list = append(list, n)
	}
	respondJSON(w, 200, list)
}

// GET BY ID Notificacion
func GetNotificacionByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var n models.Notificacion
	err := config.DB.QueryRow(`
		SELECT id_notificacion, id_usuario, titulo, mensaje, tipo, id_referencia, tipo_referencia, leido, activo, fecha_creacion, fecha_modificacion 
		FROM "Sistema"."Notificacion" WHERE id_notificacion=$1`, id).
		Scan(&n.ID_Notificacion, &n.ID_Usuario, &n.Titulo, &n.Mensaje, &n.Tipo, &n.ID_Referencia, &n.TipoReferencia, &n.Leido, &n.Activo, &n.FechaCreacion, &n.FechaModificacion)
	if err != nil {
		respondJSON(w, 404, map[string]string{"error": "Notificacion no encontrada"})
		return
	}
	respondJSON(w, 200, n)
}

// CREATE Notificacion
func CreateNotificacion(w http.ResponseWriter, r *http.Request) {
	var n models.Notificacion
	if err := json.NewDecoder(r.Body).Decode(&n); err != nil {
		respondJSON(w, 400, map[string]string{"error": "JSON invalido"})
		return
	}
	err := config.DB.QueryRow(`
		INSERT INTO "Sistema"."Notificacion" (id_usuario, titulo, mensaje, tipo, id_referencia, tipo_referencia, leido, activo) 
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING id_notificacion`,
		n.ID_Usuario, n.Titulo, n.Mensaje, n.Tipo, n.ID_Referencia, n.TipoReferencia, n.Leido, n.Activo,
	).Scan(&n.ID_Notificacion)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 201, n)
}

// UPDATE Notificacion
func UpdateNotificacion(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var n models.Notificacion
	json.NewDecoder(r.Body).Decode(&n)
	_, err := config.DB.Exec(`
		UPDATE "Sistema"."Notificacion" SET titulo=$1, mensaje=$2, tipo=$3, leido=$4, activo=$5 
		WHERE id_notificacion=$6`,
		n.Titulo, n.Mensaje, n.Tipo, n.Leido, n.Activo, id,
	)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "Notificacion actualizada correctamente"})
}