APP_NAME=economicus
APP_DIR=./cmd/economicus

all: build

build:
	GOARCH=amd64 GOOS=darwin go build -o ${APP_NAME} ${APP_DIR}/*.go

run:
	./${APP_NAME}

build_and_run: build run

clean:
	go clean
	rm ${APP_NAME}-darwin
	rm ${APP_NAME}-linux
	rm ${APP_NAME}-windows

test:
	go test ./...

test_api:
	go test ./app/internal/api/.*go
