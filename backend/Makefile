.PHONY: build

hello:
	echo "Hello"

dev:
	go run cmd/main.go

build:
	go build -o build/main.out cmd/main.go

run:
	./build/main.out

compile:
	echo "Compiling for every OS and Platform"
	CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o build/main-linux-arm cmd/main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o build/main-linux-arm64 cmd/main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/main-linux-amd64 cmd/main.go
	CGO_ENABLED=0 GOOS=freebsd GOARCH=386 go build -o build/main-freebsd-386 cmd/main.go

all: hello build