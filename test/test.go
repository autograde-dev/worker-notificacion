package test

import (
	"os"
	"context"
	"time"
	amqp "github.com/rabbitmq/amqp091-go"
	connection "github.com/autograde-dev/worker-notificacion/rabbitmqconnection"
)

func Test() {
	queue_name := os.Getenv("QUEUE_NAME")
	if queue_name == "" {
		queue_name = "notifications"
	}
	conn, ch := connection.ConnectMQ()
	defer connection.CloseMQ(conn, ch)
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
	body:= []byte(`{"idEvaluation":1,"isValid":true}`)
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

