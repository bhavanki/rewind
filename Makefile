.PHONY: build
build:
	CGO_ENABLED=1 go build ./...

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: test
test:
	CGO_ENABLED=1 go test ./...
