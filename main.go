package main

import (
	"github.com/oleksiivelychko/go-queue-service/mq"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
)

func main() {
	conn, err := mq.New()
	mq.FailOnError(err)
	defer func(conn *amqp.Connection) {
		_ = conn.Close()
	}(conn)

	ch, err := conn.Channel()
	mq.FailOnError(err)
	defer func(ch *amqp.Channel) {
		_ = ch.Close()
	}(ch)

	queue, err := mq.Queue(ch, os.Getenv("MQ_NAME"))

	mq.FailOnError(err)

	messages, err := ch.Consume(queue.Name, "", true, false, false, false, nil)

	mq.FailOnError(err)

	forever := make(chan bool)

	go func() {
		for d := range messages {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

	<-forever
}
