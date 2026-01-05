build:
	@go build -o bin/social-api cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/social-api