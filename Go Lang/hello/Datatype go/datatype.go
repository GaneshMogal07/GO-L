package main

import "fmt"

func main() {
	// Integer
	var i int = 42
	// Float
	var f float64 = 3.14
	// Boolean
	var b bool = true
	// String
	var s string = "Hello, Go!"
	// Byte (alias for uint8)
	var by byte = 65 // ASCII for 'A'
	// Rune (alias for int32, used for Unicode characters)
	var r rune = 'G'

	// Print all values
	fmt.Println("Integer:", i)
	fmt.Println("Float:", f)
	fmt.Println("Boolean:", b)
	fmt.Println("String:", s)
	fmt.Println("Byte (as character):", string(by))
	fmt.Println("Rune (as character):", string(r))
}
