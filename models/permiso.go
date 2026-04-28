package models

import "time"

type Permiso struct {
	ID_Permiso        int       `json:"id_permiso"`
	Nombre            string    `json:"nombre"`
	Descripcion       string    `json:"descripcion"`
	Activo            bool      `json:"activo"`
	FechaCreacion     time.Time `json:"fecha_creacion"`
	FechaModificacion time.Time `json:"fecha_modificacion"`
}
