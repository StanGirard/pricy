.DEFAULT_TARGET=help
VERSION:=$(shell cat VERSION)

## help: Display list of commands
.PHONY: help
help: Makefile
	@sed -n 's|^##||p' $< | column -t -s ':' | sed -e 's|^| |'

## build: Build tfautomv binary
.PHONY: build
build: fmt vet lint
	go build -o bin/pricy

## fmt: Format source code
.PHONY: fmt
fmt:
	go fmt ./...

## vet: Vet source code
.PHONY: vet
vet:
	go vet ./...

## lint: Lint source code
.PHONY: lint
lint:
	golangci-lint run

## test: Run unit tests
.PHONY: test
test:
	go test ./...

## release: Release a new version
.PHONY: release
release: test
	git tag -a "$(VERSION)" -m "$(VERSION)"
	git push origin "$(VERSION)"
	# goreleaser release --rm-dist