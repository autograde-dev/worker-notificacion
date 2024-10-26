package student

// import (
// 	psql_conn "github.com/autograde-dev/worker-notificacion/postgresconnection"
// )

type Student struct {
	Id_estudiante int
	Primer_nombre      string
	Segundo_nombre     string
	Primer_apellido   string
	Segundo_apellido  string
	Correo string	
}


func (s *Student) GetStudentByEvaluationId(evaluationId int) {
	s.Id_estudiante = evaluationId
	s.Primer_nombre = "Juan"
	s.Segundo_nombre = "Pablo"
	s.Primer_apellido = "Perez"
	s.Segundo_apellido = "Gomez"
	s.Correo = "a@g.com"
}