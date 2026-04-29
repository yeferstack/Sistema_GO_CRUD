package models

import "time"

type Notificacion struct {
	ID_Notificacion   int       `json:"id_notificacion"`
	ID_Usuario        int       `json:"id_usuario"`
	Titulo            string    `json:"titulo"`
	Mensaje           string    `json:"mensaje"`
	Tipo              string    `json:"tipo"`
	ID_Referencia     *int      `json:"id_referencia"`
	TipoReferencia    string    `json:"tipo_referencia"`
	Leido             bool      `json:"leido"`
	Activo            bool      `json:"activo"`
	FechaCreacion     *time.Time `json:"fecha_creacion"`
	FechaModificacion *time.Time `json:"fecha_modificacion"`
}
