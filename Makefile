all: fmt

fmt:
	go fmt . ./examples

test: fmt
	go run ./examples/test.go -h
