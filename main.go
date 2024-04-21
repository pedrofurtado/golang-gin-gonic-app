package main

import (
  "my-app/cmd"
)

func main() {
	go cmd.SetupSQSConsumers()
	cmd.SetupWebApp().Run(":80")
}
