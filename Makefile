GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")
LDFLAGS=-s -w

.PHONY: build
build:
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-format ./cmd/wof-format/main.go

.PHONY: test
test:
	go test -v ./test/...
