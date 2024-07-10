build:
	@go build -o bin/go-school-api

run: build
	@./bin/go-school-api

test:
	@go test -v ./...