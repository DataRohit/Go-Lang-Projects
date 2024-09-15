package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp.Dial("amqp://root:example@localhost:5672/")
	handleError(err, "Connection failed")
	defer conn.Close()

	channel, err := conn.Channel()
	handleError(err, "Channel creation failed")
	defer channel.Close()

	queue, err := channel.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	handleError(err, "Queue declaration failed")

	messages, err := channel.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	handleError(err, "Consumer registration failed")

	go func() {
		for msg := range messages {
			log.Printf("Received: %s", msg.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. Press CTRL+C to exit")
	select {}
}

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
