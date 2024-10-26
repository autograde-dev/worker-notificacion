package notification

import (
	"log"
)


func (n *Notification) SendEmailNotification() {
	log.Println("Sending email notification" + n.GetNotificationMessage())
}