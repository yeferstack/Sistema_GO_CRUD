package models

import "time"

type Administrador struct {
	ID_Admin          int       `json:"id_admin"`
	ID_Usuario        int       `json:"id_usuario"`
	ID_NivelAcceso    string    `json:"id_nivelacceso"`
	Activo            bool      `json:"activo"`
	FechaAsignacion   *time.Time `json:"fecha_asignacion"`
	FechaCreacion     *time.Time `json:"fecha_creacion"`
	FechaModificacion *time.Time `json:"fecha_modificacion"`
}
