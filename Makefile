APP=nba-simulator

install:
	@echo "Installing..."
	go mod download

build:
	@echo "Building..."
	go build -o ${APP} cmd/app/main.go

run: build
	@echo "Running..."
	./${APP}

dev:
	@echo "Running..."
	go run cmd/app/main.go