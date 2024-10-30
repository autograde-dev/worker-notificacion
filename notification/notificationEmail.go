package notification

import (
	"log"
	"os"
	"strconv"

	student "github.com/autograde-dev/worker-notificacion/student"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type NotificationEmail struct {
	IdEvaluation int
	Student      student.Student
	IsValid      bool
}

func (n *NotificationEmail) GetNotificationMessage() string {
	msg := "dear " + n.Student.PrimerNombre + " " + n.Student.PrimerApellido + "this is an email notification that your evaluation with id " + strconv.Itoa(n.IdEvaluation) + " is ready. "
	if n.IsValid {
		msg += "your result is Valid "
	} else {
		msg += "your result is Invalid"
	}
	return msg
}

func (n *NotificationEmail) Notify() {
	from := mail.NewEmail("Example User", "test@example.com")
	subject := ""
	to := mail.NewEmail("Correo estudiante", n.Student.Correo)
	plainTextContent := n.GetNotificationMessage()
	htmlContent := ""
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(response.StatusCode)
		log.Println(response.Body)
		log.Println(response.Headers)
	}
}
