# Kafka Test Client #

Requirements:

- golang > 1.14
- make

## Kafka Producer ##

Follow steps below for testing kafka producer:

- check `configs/app/config.yaml` file set config parameters
- then execute `make start-producer`

Properties:

- kafka.producer.brokerUrls - [string array] broker urls as list
- kafka.producer.topicName - [string] topic which producer will send a message

PS: it will send 10 test messages to specified topic in `configs/app/config.yaml`.

## Kafka Consumer ##

Follow steps below for testing kafka consumer:

- check configs/app/config.yaml file set config parameters
- then execute `make start-consumer`

Properties:

- kafka.consumer.brokerUrls - [string array] broker urls as list
- kafka.consumer.topicName - [string] topic which consumer will subscribe
- kafka.consumer.groupId - [string] consumer group property

PS: it will consume messages and print key and value as string format untill you interrupt the program.

## Kafka Admin ##

Kafka admin tool, can manage kafka topics (create, info, delete). Parameters are below:

- mode - [string] available values:
  - info - shows info about the topic (default)
  - create - creates topic
  - delete - deletes topic

- brokerUrl - [string] describes kafka lead url
- topic - [string] describes topic, when creating topic you can specify topic details with this pattern `{topicName}:{partitions}:{replicas}`

Available commands:

- for deleting the topic `make -e MODE="delete" BROKER_URL="localhost:9092" TOPIC="test" start-admin`
- for creating the topic with 5 partiotions and 1 replicas `make -e MODE="create" BROKER_URL="localhost:9092" TOPIC="test:5:1" start-admin`
- for fetching info abot the topic `make -e MODE="info" BROKER_URL="localhost:9092" TOPIC="test" start-admin`
