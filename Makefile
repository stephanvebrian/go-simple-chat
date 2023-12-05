test:
	@echo ">  Running tests..."
	go test -cover -race -v ./...

build:
	go build -o ./bin/go-simple-chat ./cmd

run: build
	./bin/go-simple-chat