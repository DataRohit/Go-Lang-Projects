package main

import (
	"context"
	"log"
	"os"
	"strings"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs_topic",
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare an exchange")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := getMessageBody(os.Args)
	routingKey := getRoutingKey(os.Args)

	err = ch.PublishWithContext(ctx,
		"logs_topic",
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	failOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s with routing key %s", body, routingKey)
}

func getMessageBody(args []string) string {
	if len(args) < 3 || args[2] == "" {
		return "hello"
	}
	return strings.Join(args[2:], " ")
}

func getRoutingKey(args []string) string {
	if len(args) < 2 || args[1] == "" {
		return "anonymous.info"
	}
	return args[1]
}
