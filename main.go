package main

import (
	"syscall/js"
)

func add(this js.Value, inputs []js.Value) interface{} {
	return inputs[0].Float() + inputs[1].Float()
}

func main() {
	// Go WebAssemblies are different than Rust or other languages.
	// Binaries are programs that runs for once they exit, instead of being like a static library.
	// So here we prevent the program to exit by creating an open channel.
	c := make(chan int)

	// Export Go function
	js.Global().Set("add", js.FuncOf(add))

	// Call JS function
	n := js.Global().Call("add", 3, 4).Int()
	s := js.Global().Call("hello").String()

	println(n, s)

	// Access JS variables
	one := js.Global().Get("env").String()
	two := js.Global().Get("env2").String()

	println(one, two)

	js.Global().Set("env2", "CHANGED!")
	three := js.Global().Get("env2").String()

	println(one, three)

	// Manipulate DOM
	document := js.Global().Get("document")
	h1 := document.Call("createElement", "h1")
	h1.Set("innerText", "This is H1")

	document.Get("body").Call("appendChild", h1)

	// Program is now exited
	<-c
}
