# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get

# Main package name
MAIN_PACKAGE = ./cmd/main.go

# Output binary name
BINARY_NAME = ./build/bin

# Build targets
all: clean build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v $(MAIN_PACKAGE)

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

test:
	$(GOTEST) -v ./...

run:
	$(GOBUILD) -o $(BINARY_NAME) -v $(MAIN_PACKAGE)
	./$(BINARY_NAME)

deps:
	$(GOGET) -v $(MAIN_PACKAGE)

.PHONY: all build clean test run deps
