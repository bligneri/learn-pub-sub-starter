package pubsub

import (
	"context"
	"encoding/json"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

func PublishJSON[T any](ch *amqp.Channel, exchange, key string, val T) error {
	ctx := context.Background()

	valBytes, err := json.Marshal(val)
	if err != nil {
		return fmt.Errorf("Error marshalling key %s with value: %w", key, val)
	}

	msg := amqp.Publishing{
		ContentType: "application/json",
		Body:        valBytes,
	}

	err = ch.PublishWithContext(ctx, exchange, key, false, false, msg)
	if err != nil {
		return fmt.Errorf("Error publishing message %w", msg)
	}

	return nil
}
