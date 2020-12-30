GOPATH := $(GOPATH):$(shell dirname $(shell pwd))
BROKER_URL := "localhost:9092"
TOPIC := "test"
MODE := "info"

start-consumer:
	@CONFIG_FILE="./configs/app/config.yaml" go run "github.com/nazarov-pro/kafka-client/cmd/consumer"

start-producer:
	@CONFIG_FILE="./configs/app/config.yaml" go run "github.com/nazarov-pro/kafka-client/cmd/producer"

start-admin:
	go run "github.com/nazarov-pro/kafka-client/cmd/admin" -mode ${MODE} -brokerUrl ${BROKER_URL} -topic ${TOPIC}
