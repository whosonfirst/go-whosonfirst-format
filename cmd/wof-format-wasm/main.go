//go:build wasmjs
package main

import (
	"log"
	"syscall/js"

	"github.com/whosonfirst/go-whosonfirst-format/wasm"
)

func main() {

	format_func := wasm.FormatFunc()
	defer format_func.Release()

	js.Global().Set("wof_format", format_func)

	c := make(chan struct{}, 0)

	log.Println("wof_format function initialized")
	<-c
}
