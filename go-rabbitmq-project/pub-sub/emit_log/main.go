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
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := messageBody(os.Args)

	err = channel.PublishWithContext(
		ctx,
		"logs",
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	handleError(err, "Message publishing failed")

	log.Printf(" [x] Sent %s", body)
}

func messageBody(args []string) string {
	if len(args) < 2 || args[1] == "" {
		return "hello"
	}
	return strings.Join(args[1:], " ")
}

func handleError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
