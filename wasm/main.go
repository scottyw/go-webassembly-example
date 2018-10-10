package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	js.Global().Get("document").Call("getElementById", "helloworld").Set("textContent", "Hello, WebAssembly!")
	fmt.Println("Also available on the console!")
}
