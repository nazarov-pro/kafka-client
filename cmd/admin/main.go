package main

import (
	"fmt"
	"os"

	"github.com/nazarov-pro/kafka-client/internal/conf"
	"github.com/nazarov-pro/kafka-client/pkg/utils"
)

func main() {
	var (
		args      = utils.ArgumentParser(os.Args)
		mode      = args["mode"].(string)
		brokerURL = args["brokerUrl"].(string)
		topic     = args["topic"].(string)
	)
	switch mode {
	default:
		fallthrough //info is default mode
	case "info":
		topicInfo, err := conf.FetchTopicDetails(brokerURL, topic)
		if err != nil {
			fmt.Printf("Failed to fetch topic info %s, err %v.\n", topic, err)
		} else {
			fmt.Printf(
				"%s topic info (partitions: %d, repicas: %d, leaderID: %d).\n",
				topicInfo.TopicName,
				topicInfo.Partitions,
				topicInfo.Replicas,
				topicInfo.LeaderID,
			)
		}

	case "create":
		err := conf.CreateTopicWithPattern(brokerURL, topic)
		if err != nil {
			fmt.Printf("Failed to create topic %s, err %v.\n", topic, err)
		} else {
			fmt.Printf("%s topic successfully created (skipped if the topic already exists).\n", topic)
		}
	case "delete":
		err := conf.DeleteTopic(brokerURL, topic)
		if err != nil {
			fmt.Printf("Failed to delete topic %s, err %v.\n", topic, err)
		} else {
			fmt.Printf("%s topic successfully deleted.\n", topic)
		}
	}
}
