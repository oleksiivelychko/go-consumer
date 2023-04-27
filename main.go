package main

import (
	"github.com/oleksiivelychko/go-queue-service/mq"
	"log"
	"os"
)

func main() {
	conn, err := mq.New()
	mq.FailOnError(err)
	defer conn.Close()

	ch, err := conn.Channel()
	mq.FailOnError(err)
	defer ch.Close()

	q, err := mq.Queue(ch, os.Getenv("MQ_NAME"))

	mq.FailOnError(err)

	messages, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

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
