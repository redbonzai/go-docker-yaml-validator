# Variables
BINARY_NAME=docker-compose-validate
SRC_DIR=./cmd
VALIDATE_CMD=validate

# Build the Go binary
build:
	go build -o $(BINARY_NAME) $(SRC_DIR)

# Run the validator with a specified Docker Compose file
run: build
	./$(BINARY_NAME) $(VALIDATE_CMD) $(FILE)

# Clean the build
clean:
	rm -f $(BINARY_NAME)

# Run tests
test:
	go test ./... -v

# Help
help:
	@echo "Makefile commands:"
	@echo "  build      - Build the Go binary"
	@echo "  run FILE=<path/to/docker-compose.yml> - Run the validator with a specified Docker Compose file"
	@echo "  clean      - Clean the build"
	@echo "  test       - Run tests"
	@echo "  help       - Show this help message"
