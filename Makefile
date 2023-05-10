build:
	@go build -o ./bin/api

run: build
	@./bin/api --port=3333

test:
	@go test -v ./...
