# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get
GOFMT = $(GOCMD) fmt
GOVET = $(GOCMD) vet

# Main package path
MAIN_PACKAGE_PATH = ./cmd/server

# Main binary name
BINARY_NAME = charmander

# Build target
build:
	$(GOBUILD) -o $(BINARY_NAME) $(MAIN_PACKAGE_PATH)

# Test target
test:
	$(GOTEST) -v ./...

# Format source code
fmt:
	$(GOFMT) ./...

# Vet source code
vet:
	$(GOVET) ./...

# Clean target
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

# Install dependencies
deps:
	$(GOGET) ./...

# Default target
all: deps fmt vet test build

# Run target (build and run)
run: build
	./$(BINARY_NAME)

# Phony targets
.PHONY: build test fmt vet clean deps all run