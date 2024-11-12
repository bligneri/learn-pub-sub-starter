package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"

	"github.com/bootdotdev/learn-pub-sub-starter/internal/routing"

	// If PublishJSON is from pubsub package, import it accordingly:
	"github.com/bootdotdev/learn-pub-sub-starter/internal/pubsub"

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

	channel, err := con.Channel()
	if err != nil {
		fmt.Println("Error :", err)
		return
	}

	state, err := json.Marshal(routing.PlayingState{IsPaused: true})
	if err != nil {
		fmt.Println("Error :", err)
		return
	}

	fmt.Println("PublishJSON to channel", routing.ExchangePerilDirect, routing.PauseKey, state)
	pubsub.PublishJSON(channel, routing.ExchangePerilDirect, routing.PauseKey, state)

	// wait for ctrl+c
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
}
