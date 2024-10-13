build:
	go build -o build/ .
format:
	gofmt -w .
run:
	go run .
test:
	go test ./...