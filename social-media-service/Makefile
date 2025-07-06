# Project variables
PROTO_DIR=proto
PROTO_FILE=$(PROTO_DIR)/user.proto
GEN_GO_DIR=.
MAIN_FILE=cmd/main.go
SEED_FILE=scripts/seed.go

BIN_NAME=social-service
IMAGE_NAME=social-service:latest
CONTAINER_NAME=social-service-container

# Default target
all: build

build:
	go build -o $(BIN_NAME) $(MAIN_FILE)

proto:
	protoc --go_out=$(GEN_GO_DIR) --go-grpc_out=$(GEN_GO_DIR) $(PROTO_FILE)

seed:
	go run $(SEED_FILE)

run:
	go run $(MAIN_FILE)

clean:
	rm -f $(BIN_NAME)

install-tools:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Docker build and run
docker-build: proto
	docker build -t $(IMAGE_NAME) .

docker-run:
	docker run -d --name $(CONTAINER_NAME) -p 8080:8080 $(IMAGE_NAME)

docker-stop:
	docker stop $(CONTAINER_NAME) && docker rm $(CONTAINER_NAME)

.PHONY: all build proto seed run clean install-tools docker-build docker-run docker-stop
