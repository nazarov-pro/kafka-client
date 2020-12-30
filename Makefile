GOPATH := $(GOPATH):$(shell dirname $(shell pwd))

start-consumer:
	@CONFIG_FILE="./configs/app/config.yaml" go run "github.com/nazarov-pro/kafka-client/cmd/consumer"

start-producer:
	@CONFIG_FILE="./configs/app/config.yaml" go run "github.com/nazarov-pro/kafka-client/cmd/producer"
