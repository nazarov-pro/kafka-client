package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/nazarov-pro/kafka-client/internal/conf"
	"github.com/nazarov-pro/kafka-client/internal/kafka/consumer"
	"github.com/oklog/run"
)

func main() {
	fmt.Println("Consumer app is starting...")
	conf.LoadConfiguration()
	
	var g run.Group
	//Consumer
	{
		ctx, cancel := context.WithCancel(context.Background())
		r := conf.NewReader(
			conf.Config.GetStringSlice("kafka.consumer.brokerUrls"),
			conf.Config.GetString("kafka.consumer.topicName"),
			conf.Config.GetString("kafka.consumer.groupId"),
		)
		g.Add(
			func() error {
				return consumer.ConsumeMessages(ctx, r)
			}, func(error) {
				cancel()
				fmt.Printf("Kafka consumer closed, err=%v\n", r.Close())
			},
		)
	}

	//Signal Catcher
	{
		sigChan := make(chan os.Signal)
		g.Add(func() error {
			signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
			return fmt.Errorf("%v", <-sigChan)
		}, func(error) {
			close(sigChan)
		})
	}
	fmt.Println("Application stopped, err: ", g.Run())
}
