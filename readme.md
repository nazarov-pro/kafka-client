# Kafka Test Client #

Requirements:

- golang > 1.14
- make

## Kafka Producer ##

Follow steps below for test kafka producer:

- check `configs/app/config.yaml` file set config parameters
- then execute `make start-producer`

Properties:

- kafka.producer.brokerUrls - [string array] broker urls as list
- kafka.producer.topicName - [string] topic which producer will send a message

PS: it will send 10 test messages to specified topic in `configs/app/config.yaml`.

## Kafka Consumer ##

Follow steps below for test kafka consumer:

- check configs/app/config.yaml file set config parameters
- then execute `make start-consumer`

Properties:

- kafka.consumer.brokerUrls - [string array] broker urls as list
- kafka.consumer.topicName - [string] topic which consumer will subscribe
- kafka.consumer.groupId - [string] consumer group property

PS: it will consume messages and print key and value as string format untill you interrupt the program.

## Kafka Admin ##

Kafka admin will create topics if topic is not exist in kafka. It will trigger with `kafka.admin.createTopics`-property, defined in config.

Properties:

- kafka.admin.brokerUrl - [string] broker url, specify lead broker url
- kafka.admin.defaults.partitions - [int] default partition size for the topic (default is 1)
- kafka.admin.defaults.replicas - [int] default replication factor size for the topic (default is 1)
- kafka.admin.createTopics - [bool] will create topics automatically in `kafka.admin.topics` config group.
- kafka.admin.topics.topicName - [string] as template `{topicName}:{partitionSize}:{replicaSize}`, if partition size or replicaSize not defined it will use defaults.
