package notification

import (
	"log"
	"os"
	"strconv"

	student "github.com/autograde-dev/worker-notificacion/student"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	gomail "gopkg.in/mail.v2"
)

type NotificationEmail struct {
	IdEvaluation int
	Student      student.Student
	IsValid      bool
}

func (n *NotificationEmail) GetNotificationMessage() string {
	msg := "Dear " + n.Student.PrimerNombre + " " + n.Student.PrimerApellido + " your evaluation in autograde.dev with id " + strconv.Itoa(n.IdEvaluation) + " is ready. "
	if n.IsValid {
		msg += "your result is Valid "
	} else {
		msg += "your result is Invalid"
	}
	return msg
}

func (n *NotificationEmail) Notify() {
	log.Println("Sending email notification to", n.Student.Correo)
	if os.Getenv("SENDGRID_API_KEY") != "" {
		n.notifyWithSengrid()
	} else {
		n.notifyWithGomail()
	}
}

func (n *NotificationEmail) notifyWithGomail() {
	m := gomail.NewMessage()
	m.SetHeader("From", "autograde.dev@capibaraeventos.com")
	m.SetHeader("To", n.Student.Correo)
	m.SetHeader("Subject", "Evaluation Notification")
	m.SetBody("text/plain", n.GetNotificationMessage())
	port, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		log.Println("Invalid SMTP_PORT:", err)
		return
	}
	d := gomail.NewDialer(os.Getenv("SMTP_SERVER"), port, os.Getenv("SMTP_USER"), os.Getenv("SMTP_PASSWORD"))
	if err := d.DialAndSend(m); err != nil {
		log.Println(err)
	} else {
		log.Println("Email sent successfully")
	}
}

func (n *NotificationEmail) notifyWithSengrid() {
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
