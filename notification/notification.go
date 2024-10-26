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
	msg := "dear " + n.student.Primer_nombre + " " + n.student.Primer_apellido + " your evaluation is ready. "
	if n.IsValid {
		msg += "your result is Valid "
	} else {
	msg += "your result is Invalid"
	}
	return msg
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
	log.Println("Log notification:" + n.GetNotificationMessage())
}