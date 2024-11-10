package main

import (
	"fmt"

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
}
