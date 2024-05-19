build:
	@go build -o bin/authentication-project

run: build
	@./bin/authentication-project

test:
	@go test -v ./...