package main

import (
	"github.com/oleksiivelychko/go-queue-service/initmq"
	"log"
)

func main() {
	initmq.LoadEnv("rabbitmq")

	conn, err := initmq.MQ()
	initmq.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	initmq.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	initmq.FailOnError(err, "Failed to declare a queue")

	messages, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	initmq.FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range messages {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
