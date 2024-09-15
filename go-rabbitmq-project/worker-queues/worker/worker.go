package main

import (
	"bytes"
	"log"
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

	err = channel.Qos(1, 0, false)
	handleError(err, "QoS setup failed")

	messages, err := channel.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	handleError(err, "Consumer registration failed")

	go func() {
		for msg := range messages {
			log.Printf("Received: %s", msg.Body)

			sleepDuration := time.Duration(bytes.Count(msg.Body, []byte("."))) * time.Second
			time.Sleep(sleepDuration)

			log.Println("Processing complete")
			msg.Ack(false)
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
