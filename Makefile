.PHONY: build
build:
	CGO_ENABLED=1 go build -o rewind ./cmd/rewind

.PHONY: fmt
fmt:
	go fmt ./...

internal/store/store_mock.go: internal/store/store.go
	go generate ./internal/store

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test: internal/store/store_mock.go
	CGO_ENABLED=1 go test ./...

.PHONY: clean
clean:
	rm -f rewind
