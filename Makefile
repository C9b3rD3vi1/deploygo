BINARY_NAME=deploygo

build:
	go build -o $(BINARY_NAME) ./cmd/deploygo

install:
	go install ./cmd/deploygo

test:
	go test ./...

clean:
	rm -f $(BINARY_NAME)

.PHONY: build install test clean