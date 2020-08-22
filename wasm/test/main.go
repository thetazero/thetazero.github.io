package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	c := make(chan struct{}, 0)
	fmt.Println("WASM Go Initialized")
	// register functions
	js.Global().Set("add", js.FuncOf(add))

	<-c
}

func add(this js.Value, i []js.Value) interface{} {
	return js.ValueOf(i[0].Int() + i[1].Int())
}

// exposing to JS
