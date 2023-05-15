build:
	@go build -o ./bin/api

run: build
	@./bin/api --port 3333

seed:
	@go run scripts/seed.go

test:
	@go test -v ./...
