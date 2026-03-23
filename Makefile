# Define variables for easy updates
BINARY_NAME=celes
MAIN_PATH=./cmd/celes
.PHONY:
test:
	go test ./...
# Run go fmt on all Go files
fmt:
	go fmt ./...

# The 'all' target is usually the default
all: build

# Build the Go binary
build:
	go build -o $(BINARY_NAME) $(MAIN_PATH)

# Run the project
run:
	go run $(MAIN_PATH)

# Run the linter you just installed
lint:
	golangci-lint run ./...

# Clean up the binary
clean:
	rm -f $(BINARY_NAME)

# Help command to list available tasks
help:
	@echo "Available targets:"
	@echo "  build  - Compile the project"
	@echo "  run    - Run the API"
	@echo "  lint   - Run golangci-lint"
	@echo "  clean  - Remove the binary"