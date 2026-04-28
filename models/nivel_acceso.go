package models

import "time"

type NivelAcceso struct {
	ID_Admin          int       `json:"id_admin"`
	ID_NivelAcceso    string    `json:"id_nivelacceso"`
	Moderador         string    `json:"moderador"`
	Administrador     string    `json:"administrador"`
	Activo            bool      `json:"activo"`
	FechaAsignacion   time.Time `json:"fecha_asignacion"`
	FechaCreacion     time.Time `json:"fecha_creacion"`
	FechaModificacion time.Time `json:"fecha_modificacion"`
}
