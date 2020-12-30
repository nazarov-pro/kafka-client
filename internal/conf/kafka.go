package conf

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/segmentio/kafka-go"
)

var dialer = newDialer()

func newDialer() *kafka.Dialer {
	return &kafka.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
	}
}

// NewWriter - kafka writer for producing messages
func NewWriter(topicName string) *kafka.Writer {
	brokerUrls := Config.GetStringSlice("kafka.producer.brokerUrls")
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  brokerUrls,
		Topic:    topicName,
		Balancer: &kafka.LeastBytes{},
		Dialer:   dialer,
	})
	return w
}

// NewReader - kafka reader for consuming messages
func NewReader(topicName string, groupID string) *kafka.Reader {
	brokerUrls := Config.GetStringSlice("kafka.consumer.brokerUrls")
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokerUrls,
		GroupID:        groupID,
		Topic:          topicName,
		MinBytes:       1,
		MaxBytes:       10e6,
		CommitInterval: time.Second,
		Dialer:         dialer,
		SessionTimeout: 5 * time.Second,
	})
	return r
}

// CreateTopics - generating topics which defined in config
func CreateTopics() {
	topics := Config.GetStringMapString("kafka.admin.topics")
	for _, value := range topics {
		err := CreateTopic(extractTopicDetailsFromPattern(value))
		if err != nil {
			fmt.Printf("Error occured %v\n", err)
		}
	}
}

// CreateTopic - creates topics with topic configs
func CreateTopic(topicName string, partitions, replicas int) error {
	conn, err := kafka.Dial("tcp", Config.GetString("kafka.admin.brokerUrl"))
	if err != nil {
		return err
	}
	defer conn.Close()
	fmt.Printf("Creating %s topic with %d partition(s) and %d replica(s) if not exist.\n", topicName, partitions, replicas)
	return conn.CreateTopics(
		kafka.TopicConfig{
			Topic:             topicName,
			NumPartitions:     partitions,
			ReplicationFactor: replicas,
		},
	)
}

func extractTopicDetailsFromPattern(pattern string) (string, int, int) {
	var (
		items      = strings.Split(pattern, ":")
		topicName  = items[0]
		partitions = Config.GetInt("kafka.admin.defaults.partitions")
		replicas   = Config.GetInt("kafka.admin.defaults.replicas")
	)
	switch len(items) {
	case 3:
		replicas, _ = strconv.Atoi(items[2])
		fallthrough
	case 2:
		partitions, _ = strconv.Atoi(items[1])
	}
	return topicName, partitions, replicas
}
