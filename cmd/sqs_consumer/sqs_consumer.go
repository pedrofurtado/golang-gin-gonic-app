package cmd

import (
	"os"
	"fmt"
	"strconv"

	sqsclient "github.com/inaciogu/go-sqs/consumer"
	"github.com/inaciogu/go-sqs/consumer/message"
	"go.uber.org/zap"
)

type SQSConsumerLogger struct {
	*zap.Logger
}

func (l *SQSConsumerLogger) Log(message string, v ...interface{}) {
	fmt.Printf("[SQS consumer] %v\n", fmt.Sprintf(message, v...))
}

type Message struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func SetupSQSConsumers() {
	sqsWaitTimeSeconds, err := strconv.ParseInt(os.Getenv("APP_SQS_WAIT_TIME_SECONDS"), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("ERRO AO CONVERTER STRING PARA INT %v", err))
	}

	sqsConsumer := sqsclient.New(nil, sqsclient.SQSClientOptions{
		QueueName: os.Getenv("APP_SQS_QUEUE_NAME"),
		Region: os.Getenv("AWS_DEFAULT_REGION"),
		MaxNumberOfMessages: 10,
		WaitTimeSeconds: sqsWaitTimeSeconds,
		VisibilityTimeout: 30,
		Handle: func(message *message.Message) bool {
			myMessage := Message{}

			// Unmarshal the message content
			err := message.Unmarshal(&myMessage)

			if err != nil {
				fmt.Println("ERRO NO UNMARSHAL DA MENSAGEM", err)

				// Do something if the message content cannot be unmarshalled
				return false
			}

			fmt.Println("MESSAGE RECEBIDA OLHA SO OS DADOS", myMessage.Name, myMessage.Email)

			return true // to mark SQS message as success or failure
		},
	})

	sqsConsumer.SetLogger(&SQSConsumerLogger{})

	sqsConsumer.Start()
}
