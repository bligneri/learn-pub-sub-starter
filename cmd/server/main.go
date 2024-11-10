package main

import (
	"fmt"
	"os"
	"os/signal"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Println("Starting Peril server...")

	var rabbitQueue = "amqp://guest:guest@localhost:5672"

	con, err := amqp.Dial(rabbitQueue)
	if err != nil {
		fmt.Println("Error :", err)
		return
	}
	defer con.Close()

	// wait for ctrl+c
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
}
