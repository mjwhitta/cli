all: build

build: check fmt
	@go build .

check:
	@which go >/dev/null 2>&1

fmt: check
	@go fmt .

gen: check
	@go generate
