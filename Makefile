build:
	@go build -o bin/goecomm cmd/main.go

run: build
	@./bin/goecomm

test:
	@go test -v ./...