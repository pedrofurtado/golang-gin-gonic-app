package consumers

import (
	"fmt"
	"os"
	"time"
	"math/rand"

	amqp "github.com/rabbitmq/amqp091-go"
)

func generateRandomNumber123() int {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 3
	return rand.Intn(max - min + 1) + min
}

func SetupRabbitMQConsumers() {
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_CONSUMER_URL"))

	if err != nil {
		panic(fmt.Sprintf("[RabbitMQ consumer] Failed to connect to RabbitMQ: %v", err))
	}

	defer conn.Close()

	ch, err := conn.Channel()

	if err != nil {
		panic(fmt.Sprintf("[RabbitMQ consumer] Failed to open a RabbitMQ channel: %v", err))
	}

	defer ch.Close()

	queueName := os.Getenv("RABBITMQ_CONSUMER_QUEUE")
	q, err := ch.QueueDeclare(
		queueName,
		true,      // durable
		false,     // autoDelete
		false,     // exclusive
		false,     // noWait
		nil,       // args
	)

	if err != nil {
		panic(fmt.Sprintf("[RabbitMQ consumer] Failed to declare a RabbitMQ queue: %v", err))
	}

	messages, err := ch.Consume(
		q.Name,
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	if err != nil {
		panic(fmt.Sprintf("[RabbitMQ consumer] Failed to register a RabbitMQ consumer: %v", err))
	}

	fmt.Printf("[RabbitMQ consumer] Waiting for messages\n")

	for d := range messages {
		fmt.Printf("[RabbitMQ consumer] Message received: %v\n", string(d.Body))

		fmt.Printf("[RabbitMQ consumer] Simulating a slow processing for Message received: %v\n", string(d.Body))
		time.Sleep(5 * time.Second)

		switch randomNumber := generateRandomNumber123(); {
		case randomNumber == 1:
			err := d.Ack(false)

			if err != nil {
				fmt.Printf("[RabbitMQ consumer] Error when ACK true message: %s\n", err)
			} else {
				fmt.Printf("[RabbitMQ consumer] ACK true for Message received: %v\n", string(d.Body))
			}
		case randomNumber == 2:
			fmt.Printf("[RabbitMQ consumer] REJECT requeue TRUE for Message received: %v\n", string(d.Body))
			d.Reject(true)
		case randomNumber == 3:
			fmt.Printf("[RabbitMQ consumer] REJECT requeue FALSE for Message received: %v\n", string(d.Body))
			d.Reject(false)
		default:
			panic(fmt.Sprintf("[RabbitMQ consumer] Invalid random number %v", randomNumber))
		}
	}
}
