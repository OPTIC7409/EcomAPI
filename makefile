build:
	@go build -o bin/ecomapi cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/ecomapi