APP=nba_simulator

install:
	@echo "Installing..."
	go mod download

build:
	@echo "Building..."
	go build -o ${APP} cmd/app/main.go -dev

run: build
	@echo "Running..."
	./${APP}

dev:
	@echo "Running..."
	go run cmd/app/main.go -dev

clean:
	@echo "Cleaning..."
	go clean
	rm -f ${APP}

up:
	@echo "Running docker-compose..."
	docker-compose up -d --build

down:
	@echo "Stopping docker-compose..."
	docker-compose down

.PHONY: build run clean
