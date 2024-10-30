package notification

import (
	"log"
	"strconv"

	student "github.com/autograde-dev/worker-notificacion/student"
)

type Notification interface {
	GetNotificationMessage() string
	Notify()
}

type NotificationFactory struct {
	IdEvaluation      int             `json:"idEvaluation"`
	IsValid           bool            `json:"isValid"`
	Student           student.Student `json:"student"` // Datos del estudiante
	NotificationTypes []string        `json:"notificationTypes"`
}

func (n *NotificationFactory) CreateNotification() []Notification {
	var notifications []Notification
	if len(n.NotificationTypes) == 0 {
		n.NotificationTypes = []string{"log", "email"}
	}
	for _, notificationType := range n.NotificationTypes {
		switch notificationType {
		case "email":
			notifications = append(notifications, &NotificationEmail{IdEvaluation: n.IdEvaluation, Student: n.Student, IsValid: n.IsValid})
		case "log":
			notifications = append(notifications, &LogNotification{IdEvaluation: n.IdEvaluation, Student: n.Student, IsValid: n.IsValid})
		}
	}
	return notifications
}

func (n *NotificationFactory) Notify() {
	notifications := n.CreateNotification()
	for _, notification := range notifications {
		notification.Notify()
	}
}

type LogNotification struct {
	IdEvaluation int
	Student      student.Student
	IsValid      bool
}

func (n *LogNotification) GetNotificationMessage() string {
	msg := "dear " + n.Student.PrimerNombre + " " + n.Student.PrimerApellido + "this is a log notification that your evaluation with id " + strconv.Itoa(n.IdEvaluation) + " is ready. "
	if n.IsValid {
		msg += "your result is Valid "
	} else {
		msg += "your result is Invalid"
	}
	return msg
}

func (n *LogNotification) Notify() {
	log.Println("Log notification:" + n.GetNotificationMessage())
}
