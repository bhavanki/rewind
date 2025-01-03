.PHONY: build
build:
	CGO_ENABLED=1 go build -o rewind ./cmd/rewind

.PHONY: fmt
fmt:
	go fmt ./...

internal/store/store_mock.go: internal/store/store.go
	rm internal/store/store_mock.go
	go generate ./internal/store

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test: internal/store/store_mock.go
	CGO_ENABLED=1 go test --coverprofile=coverage.out ./...

.PHONY: clean
clean:
	rm -f rewind
