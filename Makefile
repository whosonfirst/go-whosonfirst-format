GOROOT=$(shell go env GOROOT)
GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")
LDFLAGS=-s -w

.PHONY: build
build:
	@make cli
	@make wasmjs

cli:
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" \
		-o bin/wof-format \
		./cmd/wof-format/main.go

wasmjs:
	GOOS=js GOARCH=wasm \
		go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -tags wasmjs \
		-o www/wasm/wof_format.wasm \
		cmd/wof-format-wasm/main.go

.PHONY: test
test:
	go test -v ./...
