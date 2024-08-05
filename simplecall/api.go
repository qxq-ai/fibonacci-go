package main

import (
    "syscall/js"
)

func main() {
    c := make(chan struct{}, 0)

    // Exponer la función sayHello al JavaScript
    js.Global().Set("sayHello", js.FuncOf(sayHello))

    <-c // Evitar que la función main termine
}

func sayHello(this js.Value, args []js.Value) interface{} {
    return js.ValueOf("Hello from the Go endpoint!")
}