package student

// import (
// 	psql_conn "github.com/autograde-dev/worker-notificacion/postgresconnection"
// )

type Student struct {
	id_estudiante int
	primer_nombre      string
	segundo_nombre     string
	primer_apellido   string
	segundo_apellido  string
	correo string	
}


func (s *Student) GetStudentByEvaluationId(evaluationId int) {
	s.id_estudiante = evaluationId
	s.primer_nombre = "Juan"
	s.segundo_nombre = "Pablo"
	s.primer_apellido = "Perez"
	s.segundo_apellido = "Gomez"
	s.correo = "a@g.com"
}