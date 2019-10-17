all: fmt

fmt:
	go fmt .

test:
	go run ./examples/test.go -h
