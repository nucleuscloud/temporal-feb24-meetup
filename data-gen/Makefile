GO=go

default: all

all: build

vet:
	$(GO) vet ./...
.PHONY: vet

build: vet
	CGO_ENABLED=0 $(GO) build -ldflags="-s -w" -o bin/worker worker/*.go
	CGO_ENABLED=0 $(GO) build -ldflags="-s -w" -o bin/starter starter/*.go
.PHONY: build

clean:
	rm -rf bin
.PHONY: clean
