BINARY_NAME=celes
MAIN_PATH=./cmd/celes

.PHONY: all build run listen fmt lint test clean help \
        planets moons describe-planet describe-moon

# Default target
all: build

# Build the Go binary
build:
	go build -o $(BINARY_NAME) $(MAIN_PATH)

# Run the server
listen: build
	./$(BINARY_NAME) listen

# Run without building a binary

# Usage: make run ARGS="planets"
# Usage: make run ARGS="describe planet Earth"
run:
	go run $(MAIN_PATH) $(ARGS)
# Format all Go files
fmt:
	go fmt ./...

# Run the linter
lint:
	golangci-lint run ./...

# Run tests
test:
	go test ./...

# Clean up the binary
clean:
	rm -f $(BINARY_NAME)

# ── CLI commands ────────────────────────────────────────

# List all planet names
# Usage: make planets
planets: build
	./$(BINARY_NAME) planets

# List moons of a planet
# Usage: make moons PLANET=Jupiter
moons: build
	./$(BINARY_NAME) moons $(PLANET)

# Describe a planet
# Usage: make describe-planet NAME=Earth
describe-planet: build
	./$(BINARY_NAME) describe planet $(NAME)

# Describe a moon
# Usage: make describe-moon NAME=Luna
describe-moon: build
	./$(BINARY_NAME) describe moon $(NAME)

# ── Help ────────────────────────────────────────────────
help:
	@echo "Usage: make <target>"
	@echo ""
	@echo "Build:"
	@echo "  build            - Compile the binary"
	@echo "  run              - Run without compiling"
	@echo "  clean            - Remove the binary"
	@echo ""
	@echo "Server:"
	@echo "  listen           - Start the HTTP server"
	@echo ""
	@echo "CLI:"
	@echo "  planets          - List all planets"
	@echo "  moons            - List moons  (PLANET=Jupiter)"
	@echo "  describe-planet  - Describe a planet (NAME=Earth)"
	@echo "  describe-moon    - Describe a moon   (NAME=Luna)"
	@echo ""
	@echo "Quality:"
	@echo "  fmt              - Format code"
	@echo "  lint             - Run linter"
	@echo "  test             - Run tests"