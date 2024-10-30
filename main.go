package main

import (
	"os"

	connection "github.com/autograde-dev/worker-notificacion/rabbitmqconnection"
	test "github.com/autograde-dev/worker-notificacion/test"
)

func main() {
	if os.Getenv("RUN_TEST") == "true" {
		test.Test()
	}
	queue_name := os.Getenv("RABBITMQ_QUEUE_NAME_EVA")
	if queue_name == "" {
		queue_name = "notifications"
	}
	rabbitmq := connection.RabbitMQ{QueueName: queue_name}
	rabbitmq.Consume()
}
