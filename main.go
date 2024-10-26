package main

import (
	"os"

	connection "github.com/autograde-dev/worker-notificacion/rabbitmqconnection"
	test "github.com/autograde-dev/worker-notificacion/test"
)

func main() {
	test.Test()
	queue_name := os.Getenv("QUEUE_NAME")
	if queue_name == "" {
		queue_name = "notifications"
	}
	rabbitmq := connection.RabbitMQ{QueueName: queue_name}
	rabbitmq.Consume()
}
