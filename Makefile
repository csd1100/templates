build:
	go build -o build/ ./cmd/template-generator
format:
	gofmt -w .
run:
	go run .
test:
	go test ./...