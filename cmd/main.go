package main

import (
  web "my-app/cmd/web"
  consumers "my-app/cmd/consumers"
)

func main() {
	setupConsumers()
	setupWebApp()
}

func setupConsumers() {
	go consumers.SetupRabbitMQConsumers()
	go consumers.SetupKafkaConsumers()
}

func setupWebApp() {
	web.SetupWebApp().Run(":8080")
}
