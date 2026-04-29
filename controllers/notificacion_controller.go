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