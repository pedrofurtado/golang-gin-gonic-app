package main

import (
  web "my-app/cmd/web"
  sqsConsumer "my-app/cmd/sqs_consumer"
)

func main() {
	go sqsConsumer.SetupSQSConsumers()
	web.SetupWebApp().Run(":8080")
}
