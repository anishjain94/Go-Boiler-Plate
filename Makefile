BINARY_NAME=kulu-server
GO_CMD=go build -o

build:
	@echo "Building the application..."
	cd cmd/kulu && $(GO_CMD) ../../$(BINARY_NAME)

clean:
	@echo "Cleaning up..."
	rm -f $(BINARY_NAME)

run: build
	./$(BINARY_NAME)

.PHONY: build clean run
