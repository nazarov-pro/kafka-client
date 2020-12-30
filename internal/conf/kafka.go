package conf

import (
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
func NewWriter(brokerURLs []string, topicName string) *kafka.Writer {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  brokerURLs,
		Topic:    topicName,
		Balancer: &kafka.LeastBytes{},
		Dialer:   dialer,
	})
	return w
}

// NewReader - kafka reader for consuming messages
func NewReader(brokerURLs []string, topicName string, groupID string) *kafka.Reader {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokerURLs,
		GroupID:        groupID,
		Topic:          topicName,
		MinBytes:       1,
		MaxBytes:       10e6,
		CommitInterval: time.Second,
		Dialer:         dialer,
	})
	return r
}

// CreateTopicWithPattern creating topic with pattern '{topicName}:{partitions}:{replicas}'
func CreateTopicWithPattern(brokerURL, topic string) error {
	topicName, partitions, replicas := extractTopicDetailsFromPattern(topic)
	return CreateTopic(brokerURL, topicName, partitions, replicas)
}

// TopicInfo - information about topic
type TopicInfo struct {
	TopicName  string
	Partitions int
	Replicas   int
	LeaderID     int
}

//FetchTopicDetails - fetching topic details
func FetchTopicDetails(brokerURL, topic string) (*TopicInfo, error) {
	conn, err := kafka.Dial("tcp", brokerURL)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	topic, _, _  = extractTopicDetailsFromPattern(topic)
	partitions, err := conn.ReadPartitions(topic)
	if err != nil {
		return nil, err
	}
	replicas := make(map[int]bool)
	leader := 0
	for _, partition := range partitions {
		leader = partition.Leader.ID
		for _, broker := range partition.Replicas {
			replicas[broker.ID] = true
		}
	}
	return &TopicInfo{TopicName: topic, Partitions: len(partitions), Replicas: len(replicas), LeaderID: leader}, nil
}

// CreateTopic - creates topics with topic configs
func CreateTopic(brokerURL, topicName string, partitions, replicas int) error {
	conn, err := kafka.Dial("tcp", brokerURL)
	if err != nil {
		return err
	}
	defer conn.Close()
	return conn.CreateTopics(
		kafka.TopicConfig{
			Topic:             topicName,
			NumPartitions:     partitions,
			ReplicationFactor: replicas,
		},
	)
}

//DeleteTopic - removes topic
func DeleteTopic(brokerURL, topic string) error {
	conn, err := kafka.Dial("tcp", brokerURL)
	if err != nil {
		return err
	}
	defer conn.Close()
	topic, _, _ = extractTopicDetailsFromPattern(topic)
	return conn.DeleteTopics(topic)
}

func extractTopicDetailsFromPattern(pattern string) (string, int, int) {
	var (
		items      = strings.Split(pattern, ":")
		topicName  = items[0]
		partitions = 1
		replicas   = 1
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
