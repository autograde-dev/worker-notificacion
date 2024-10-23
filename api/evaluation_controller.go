package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	evaluation "github.com/autograde-dev/worker-notificacion/evaluation"
	connection "github.com/autograde-dev/worker-notificacion/rabbitmqconnection"
	amqp "github.com/rabbitmq/amqp091-go"
)

func postEvaluation(w http.ResponseWriter, r *http.Request) {
	var evaluation evaluation.Evaluation
	err := json.NewDecoder(r.Body).Decode(&evaluation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	conn, ch := connection.ConnectMQ()
	defer connection.CloseMQ(conn, ch)
	queue_name := os.Getenv("QUEUE_NAME")
	if queue_name == "" {
		queue_name = "evaluations"
	}
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
	body, err := json.Marshal(evaluation)
	connection.FailOnError(err, "Failed to marshal evaluation")
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

	json.NewEncoder(w).Encode(evaluation)
}

func handleRequests() {
	http.Handle("/evaluation", http.HandlerFunc(postEvaluation))
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func main() {
	queue_name := os.Getenv("QUEUE_NAME")
	if queue_name == "" {
		queue_name = "evaluations"
	}
	rabbitmq := connection.RabbitMQ{QueueName: queue_name}
	rabbitmq.Consume()
	handleRequests()
}
