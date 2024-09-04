build:
	@go build -o bin/todo cmd/main.go

run: build
	@./bin/todo

test:
	@go test -v ./...