// package main

// import "fmt"

//	func main() {
//		fmt.Println("hello word ")
//	}
//
// First Go program
// package main

// import "fmt"

// // Main function
// func main() {

//		fmt.Println("!... Hello World ...!")
//	}
package main

import (
	"fmt"
)

// Main function
func main() {
	// Integer types
	var i int = 10                        // Signed integer
	var u uint = 20                       // Unsigned integer
	var i8 int8 = -128                    // 8-bit signed integer
	var u8 uint8 = 255                    // 8-bit unsigned integer
	var i16 int16 = -32768                // 16-bit signed integer
	var u16 uint16 = 65535                // 16-bit unsigned integer
	var i32 int32 = -2147483648           // 32-bit signed integer
	var u32 uint32 = 4294967295           // 32-bit unsigned integer
	var i64 int64 = -9223372036854775808  // 64-bit signed integer
	var u64 uint64 = 18446744073709551615 // 64-bit unsigned integer

	// Floating point types
	var f32 float32 = 3.14              // 32-bit floating point
	var f64 float64 = 3.141592653589793 // 64-bit floating point

	// Boolean
	var b bool = true

	// String
	var s string = "Hello, Go!"

	// Complex types
	var c64 complex64 = complex(5, 7)   // Complex number with 64-bit float parts
	var c128 complex128 = complex(2, 3) // Complex number with 128-bit float parts

	// Byte type (alias for uint8)
	var by byte = 97 // ASCII 'a'

	// Rune type (alias for int32, used for Unicode characters)
	var r rune = 'A'

	// Zero value for variables (when not initialized)
	var defaultInt int       // Defaults to 0
	var defaultFloat float64 // Defaults to 0.0
	var defaultBool bool     // Defaults to false
	var defaultString string // Defaults to ""

	// Output all variables
	fmt.Println("Integer types:")
	fmt.Printf("int: %d\nuint: %d\nint8: %d\nuint8: %d\nint16: %d\nuint16: %d\nint32: %d\nuint32: %d\nint64: %d\nuint64: %d\n",
		i, u, i8, u8, i16, u16, i32, u32, i64, u64)

	fmt.Println("\nFloating point types:")
	fmt.Printf("float32: %f\nfloat64: %f\n", f32, f64)

	fmt.Println("\nBoolean:")
	fmt.Printf("bool: %t\n", b)

	fmt.Println("\nString:")
	fmt.Printf("string: %s\n", s)

	fmt.Println("\nComplex types:")
	fmt.Printf("complex64: %v\ncomplex128: %v\n", c64, c128)

	fmt.Println("\nByte and Rune types:")
	fmt.Printf("byte: %d (char: %c)\nrune: %d (char: %c)\n", by, by, r, r)

	fmt.Println("\nZero value variables:")
	fmt.Printf("default int: %d\ndefault float64: %f\ndefault bool: %t\ndefault string: '%s'\n", defaultInt, defaultFloat, defaultBool, defaultString)
}
