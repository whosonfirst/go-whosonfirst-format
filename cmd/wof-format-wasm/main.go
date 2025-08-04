//go:build wasmjs
package main

import (
	"fmt"
	"log"
	"syscall/js"

	"github.com/whosonfirst/go-whosonfirst-format"
)

func FormatFunc() js.Func {

	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		feature_str := args[0].String()

		handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

			resolve := args[0]
			reject := args[1]

			feature_fmt, err := format.FormatBytes([]byte(feature_str))

			if err != nil {
				reject.Invoke(fmt.Printf("Failed to format feature, %v\n", err))
				return nil
			}

			resolve.Invoke(string(feature_fmt))
			return nil
		})

		promiseConstructor := js.Global().Get("Promise")
		return promiseConstructor.New(handler)
	})
}

func main() {

	format_func := FormatFunc()
	defer format_func.Release()

	js.Global().Set("wof_format", format_func)

	c := make(chan struct{}, 0)

	log.Println("wof_format function initialized")
	<-c
}
