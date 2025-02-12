.PHONY: build test run docker-up docker-down proto

build:
	go build -o bin/server cmd/server/main.go

test:
	go test -v ./...

run:
	go run cmd/server/main.go

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

proto:
	protoc --go_out=. --go-grpc_out=. api/grpc/proto/*.proto