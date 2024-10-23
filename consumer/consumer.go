package main

import (
	"os"

	connection "github.com/autograde-dev/worker-notificacion/rabbitmqconnection"
)

func main() {
	queue_name := os.Getenv("QUEUE_NAME")
	if queue_name == "" {
		queue_name = "evaluations"
	}
	rabbitmq := connection.RabbitMQ{QueueName: queue_name}
	rabbitmq.Consume()
}
