MAIN_PACKAGE_PATH := ./cmd/slideshow/main.go
BINARY_NAME := go-slideshow

.PHONY: test
test:
	go test -v -race -buildvcs -coverprofile=/tmp/coverage.out ./...
	go tool cover -html=/tmp/coverage.out

.PHONY: build
build:
	mkdir -p bin
	CGO_LDFLAGS="-Wl,-no_warn_duplicate_libraries" CGO_ENABLED=1 CC=gcc \
		go build -tags static -ldflags "-s -w" -o bin/${BINARY_NAME} ${MAIN_PACKAGE_PATH}
