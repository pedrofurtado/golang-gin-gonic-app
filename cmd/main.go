package main

import (
	models "my-app/app/models"
	consumers "my-app/cmd/consumers"
	web "my-app/cmd/web"
)

func main() {
	setupDatabase()
	setupConsumers()
	setupWebApp()
}

func setupDatabase() {
	models.SetupDatabaseConnection()
}

func setupConsumers() {
	go consumers.SetupRabbitMQConsumers()
	go consumers.SetupKafkaConsumers()
}

func setupWebApp() {
	web.SetupWebApp().Run(":8080")
}
