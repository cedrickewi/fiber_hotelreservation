include .env

docker-up:
	@echo "Starting container"
	@docker-compose -p ${BINARY} up --build -d --remove-orphans

docker-down:
	@echo "stopping container"
	@docker-compose down 

build:
	@go build  -o ${BINARY} ./cmd/api/

run: build
	@./${BINARY}

restart: build run 

test:
	@go test -v ./...

# connection string