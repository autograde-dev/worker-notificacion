package notification

import (
	"log"
	student "github.com/autograde-dev/worker-notificacion/student"
)

type Notification struct {
	IdEvaluation int
	IsValid bool
	student student.Student

}

func (n *Notification) GetNotificationMessage() string {
	return "dear " + n.student.Primer_nombre + " " + n.student.Primer_apellido + " your evaluation is ready"
}

func (n *Notification) SendNotifications() (error) {
	student := student.Student{}
	student.GetStudentByEvaluationId(n.IdEvaluation)
	n.student = student
	n.SendLogNotification()
	n.SendEmailNotification()
	return nil
}

func (n *Notification) SendLogNotification() {
	log.Println("Sending log notification")
}