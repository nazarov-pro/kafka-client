package consumer

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

//ConsumeMessages - consume messages
func ConsumeMessages(ctx context.Context, r *kafka.Reader) error {
	for {
		m, err := r.ReadMessage(ctx)
		if err != nil {
			return err
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}
}
