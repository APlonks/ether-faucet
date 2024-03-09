.PHONY: build

hello:
	echo "Hello"

build:
	go build -o build/main cmd/main.go

run:
	./build/main -config config.yaml

compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=arm go build -o build/main-linux-arm cmd/main.go
	GOOS=linux GOARCH=arm64 go build -o build/main-linux-arm64 cmd/main.go
	GOOS=freebsd GOARCH=386 go build -o build/main-freebsd-386 cmd/main.go

all: hello build