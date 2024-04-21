package main

import (
  "my-app/cmd/web"
  "my-app/cmd/sqs_consumer"
)

func main() {
	go cmd.SetupSQSConsumers()
	cmd.SetupWebApp().Run(":8080")
}
