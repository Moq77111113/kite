.PHONY: build install test clean run dev kite

# Binary name
BINARY_NAME=kite
INSTALL_PATH=/usr/local/bin

# Build the project
build:
	@echo "Building $(BINARY_NAME)..."
	@go build -o $(BINARY_NAME) ./cmd/kite.go

# Install binary to system
install: build
	@echo "Installing $(BINARY_NAME) to $(INSTALL_PATH)..."
	@sudo mv $(BINARY_NAME) $(INSTALL_PATH)/$(BINARY_NAME)
	@echo "✓ Installed successfully!"


# Development target to run the application with arguments
kite:
	@go run ./cmd/kite.go $(filter-out $@,$(MAKECMDGOALS))

# Catch-all target to prevent make from complaining about unknown targets
%:
	@:

# Run tests
test:
	@echo "Running tests..."
	@go test -v ./...

# Clean build artifacts
clean:
	@echo "Cleaning up..."
	@rm -f $(BINARY_NAME)
	@go clean

# Run the application
run: build
	@./$(BINARY_NAME)
