package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp.Dial("amqp://root:examplet@localhost:5672/")
	handleError(err, "Connection failed")
	defer conn.Close()

	channel, err := conn.Channel()
	handleError(err, "Channel creation failed")
	defer channel.Close()

	err = channel.ExchangeDeclare(
		"logs",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	handleError(err, "Exchange declaration failed")

	queue, err := channel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	handleError(err, "Queue declaration failed")

	err = channel.QueueBind(
		queue.Name,
		"",
		"logs",
		false,
		nil,
	)
	handleError(err, "Queue binding failed")

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
		for message := range messages {
			log.Printf(" [x] %s", message.Body)
		}
	}()

	log.Printf(" [*] Waiting for logs. Press CTRL+C to exit")
	select {}
}

func handleError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
