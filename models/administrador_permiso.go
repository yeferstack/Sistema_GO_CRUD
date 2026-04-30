package models

import "time"

type AdministradorPermiso struct {
	ID                int       `json:"id"`
	ID_Admin          int       `json:"id_admin"`
	ID_Permiso        int       `json:"id_permiso"`
	Activo            bool      `json:"activo"`
	FechaCreacion     *time.Time `json:"fecha_creacion"`
	FechaModificacion *time.Time `json:"fecha_modificacion"`
}
