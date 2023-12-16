BINARY_NAME=errodus

.PHONY: all build clean
all: build

build:
	go build -o ${BINARY_NAME} cmd/main.go

clean:
	go clean
