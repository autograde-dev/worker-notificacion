package test

import (
	"context"
	"os"
	"time"

	connection "github.com/autograde-dev/worker-notificacion/rabbitmqconnection"
	amqp "github.com/rabbitmq/amqp091-go"
)

func Test() {
	queue_name := os.Getenv("RABBITMQ_QUEUE_NAME_EVA")
	if queue_name == "" {
		queue_name = "notifications"
	}
	ch := connection.ConnectMQ()
	defer ch.Close()
	q, err := ch.QueueDeclare(
		queue_name, // name
		true,       // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	connection.FailOnError(err, "Failed to declare a queue")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	body := []byte(`{"idEvaluation":1,"isValid":true, "student": {"id_estudiante":1,"primer_nombre":"John","primer_apellido":"Doe", "correo": "a@f.com"}}`)
	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         body,
		})
	connection.FailOnError(err, "Failed to publish a message")

	connection.FailOnError(err, "Failed to connect to RabbitMQ")
}
