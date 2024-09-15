package main

import (
	"context"
	"log"
	"os"
	"strings"
	"time"

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
		"task_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	handleError(err, "Queue declaration failed")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := messageBody(os.Args)

	err = channel.PublishWithContext(
		ctx,
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	handleError(err, "Message publishing failed")

	log.Printf(" [x] Sent %s\n", body)
}

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func messageBody(args []string) string {
	if len(args) < 2 || args[1] == "" {
		return "hello"
	}
	return strings.Join(args[1:], " ")
}
