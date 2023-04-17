GO := go
BINARY_NAME := kaiseki 

.PHONY: all build clean

all: build

build:
	@echo "Building $(BINARY_NAME)..."
	$(GO) build -o $(BINARY_NAME)

clean:
	@echo "Cleaning up..."
	$(GO) clean
	rm -f $(BINARY_NAME)

