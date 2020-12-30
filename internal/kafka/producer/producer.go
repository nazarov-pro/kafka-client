package producer

import (
	"context"

	"github.com/segmentio/kafka-go"
)

//SendMessage send a message
func SendMessage(ctx context.Context, w *kafka.Writer, key string, val string) error {
	err := w.WriteMessages(
		ctx,
		kafka.Message{
			Key:   []byte(key),
			Value: []byte(val),
		},
	)
	if err != nil {
		return err
	}
	return nil
}
