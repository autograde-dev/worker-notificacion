package rabbitmqconnection

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	notification "github.com/autograde-dev/worker-notificacion/notification"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	Body      string
	QueueName string
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func ConnectMQ() (*amqp.Connection, *amqp.Channel) {
	rbmq_user := os.Getenv("RABBITMQ_DEFAULT_USER")
	rbmq_pass := os.Getenv("RABBITMQ_DEFAULT_PASS")
	constr := fmt.Sprintf("amqp://%s:%s@%s:5672/", rbmq_user, rbmq_pass, os.Getenv("RABBITMQ_HOST"))
	conn, err := amqp.Dial(constr)
	FailOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")

	return conn, ch
}

func CloseMQ(conn *amqp.Connection, channel *amqp.Channel) {
	defer conn.Close()    //rabbit mq close
	defer channel.Close() //rabbit mq channel close
}

func (r *RabbitMQ) Consume() {

	conn, ch := ConnectMQ()
	defer CloseMQ(conn, ch)

	q, err := ch.QueueDeclare(
		r.QueueName, // name
		true,        // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	FailOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	FailOnError(err, "Failed to declare a queue")

	k := make(chan bool)

	go func() {
		for d := range msgs {
			var msg notification.Notification
			err := json.Unmarshal(d.Body, &msg)
			if err != nil {
				log.Printf("Error parsing JSON: %s", err)
				continue
			}
			
			err = msg.SendNotifications()
			if err != nil {
				log.Printf("Error sending notifications: %s", err)
				ch.Reject(d.DeliveryTag, true)
				continue
			}
			d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-k
}
