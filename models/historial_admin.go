package models

import "time"

type HistorialAdmin struct {
	ID_Historial      int       `json:"id_historial"`
	ID_Admin          int       `json:"id_admin"`
	Accion            string    `json:"accion"`
	Descripcion       string    `json:"descripcion"`
	TipoObjeto        string    `json:"tipo_objeto"`
	ID_Objeto         *int      `json:"id_objeto"`
	FechaAccion       *time.Time `json:"fecha_accion"`
	Activo            bool      `json:"activo"`
	FechaCreacion     *time.Time `json:"fecha_creacion"`
	FechaModificacion *time.Time `json:"fecha_modificacion"`
}
