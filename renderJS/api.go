// main.go
package main

import (
    "syscall/js"
	"fmt"
)

func main() {
    c := make(chan struct{}, 0)
    
    // Funciones de suma y resta
    js.Global().Set("sum", js.FuncOf(sum))
    js.Global().Set("subtract", js.FuncOf(subtract))
    
    // Renderizar el contenido HTML
    renderHTML()

    <-c
}

func sum(this js.Value, p []js.Value) interface{} {
    a := p[0].Float()
    b := p[1].Float()
    result := a + b
    return js.ValueOf(result)
}

func subtract(this js.Value, p []js.Value) interface{} {
    a := p[0].Float()
    b := p[1].Float()
    result := a - b
    return js.ValueOf(result)
}

func renderHTML() {
    document := js.Global().Get("document")
    body := document.Get("body")

    // Crear elementos HTML
    h1 := document.Call("createElement", "h1")
    h1.Set("innerHTML", "Go WebAssembly Example")
    body.Call("appendChild", h1)

    pSum := document.Call("createElement", "p")
    pSum.Set("innerHTML", "Sum: <span id='sumResult'></span>")
    body.Call("appendChild", pSum)

    pSubtract := document.Call("createElement", "p")
    pSubtract.Set("innerHTML", "Subtract: <span id='subtractResult'></span>")
    body.Call("appendChild", pSubtract)

    // Llamar a las funciones de suma y resta
    sumResult := js.Global().Call("sum", 5, 3).Float()
	fmt.Println("Resultado",sumResult)
    subtractResult := js.Global().Call("subtract", 5, 3).Float()
	fmt.Println("Resultado",subtractResult)

    document.Call("getElementById", "sumResult").Set("innerHTML", fmt.Sprintf("%f", sumResult))
    document.Call("getElementById", "subtractResult").Set("innerHTML", fmt.Sprintf("%f", subtractResult))
    jsCode := "alert('Hello from Go!')"
    script := document.Call("createElement", "script")
    script.Set("innerHTML", jsCode)
    body.Call("appendChild", script)
    cssCode := document.Call("createElement", "style")
    cssCode.Set("innerHTML", "#sumResult { color: green; } #subtractResult { color: red; }")
    body.Call("appendChild", cssCode)
}
