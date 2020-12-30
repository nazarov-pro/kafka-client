package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/oklog/run"

	"github.com/nazarov-pro/kafka-client/internal/conf"
	"github.com/nazarov-pro/kafka-client/internal/kafka/producer"
)

func main() {
	fmt.Println("Producer app starting...")
	conf.LoadConfiguration()

	var g run.Group
	// kafka producer
	{
		ctx, cancel := context.WithCancel(context.Background())
		w := conf.NewWriter(
			conf.Config.GetStringSlice("kafka.producer.brokerUrls"),
			conf.Config.GetString("kafka.producer.topicName"),
		)
		g.Add(func() error {
			for i := 0; i < 10; i++ {
				key := fmt.Sprintf("KEY-%d", i)
				value := fmt.Sprintf("VALUE-%d", i)
				err := producer.SendMessage(ctx, w, key, value)
				if err != nil {
					return err
				}
				fmt.Printf("Message has been sent %s - %s\n", key, value)
			}
			return nil
		}, func(error) {
			cancel()
			fmt.Printf("Kafka producer closed err=%v\n", w.Close())
		})
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
