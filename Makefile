.PHONY: build run help

build:
	go build -o bin/number-guessing-game main.go

run:
	go run main.go

help:
	@echo "Usage: make [command]"
	@echo "Commands:"
	@echo "  build: Build the project"
	@echo "  run: Run the project"
	@echo "  help: Show this help message"
