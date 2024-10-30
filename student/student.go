package student

// import (
// 	psql_conn "github.com/autograde-dev/worker-notificacion/postgresconnection"
// )

type Student struct {
	IdEstudiante    int    `json:"id_estudiante"`
	PrimerNombre    string `json:"primer_nombre"`
	SegundoNombre   string `json:"segundo_nombre"`
	PrimerApellido  string `json:"primer_apellido"`
	SegundoApellido string `json:"segundo_apellido"`
	Correo          string `json:"correo"`
}
