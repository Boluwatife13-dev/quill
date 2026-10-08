package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// ---- Integers ----
	var intNum int = 32767
	var intNum16 int16 = 32767
	// var tooBig int16 = 32768 // compile error: constant overflows int16
	var uintNum uint8 = 255 // unsigned: 0..255 (good for an RGB value)
	fmt.Println(intNum, intNum16, uintNum)

	// ---- Floats ----
	var floatNum32 float32 = 12345678.9
	var floatNum64 float64 = 12345678.9
	fmt.Println(floatNum32, floatNum64) // float32 loses precision

	// ---- Arithmetic: no mixed types without a cast ----
	var f32 float32 = 10.1
	var intNum32 int32 = 2
	result := f32 + float32(intNum32)
	fmt.Println(result)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1 / intNum2) // integer division rounds down -> 1
	fmt.Println(intNum1 % intNum2) // remainder -> 1

	// ---- Strings ----
	var myString string = "Hello" + " " + "World"
	fmt.Println(myString)
	multiLine := `Hello
World`
	fmt.Println(multiLine)
	fmt.Println(len("test"))                 // 4 bytes
	fmt.Println(len("γ"))                    // 2 bytes (UTF-8)
	fmt.Println(utf8.RuneCountInString("γ")) // 1 character
	var myRune rune = 'a'
	fmt.Println(myRune) // 97

	// ---- Booleans ----
	var myBool bool = true
	fmt.Println(myBool)

	// ---- Default values ----
	var zeroInt int
	var zeroString string
	var zeroBool bool
	fmt.Printf("%v %q %v\n", zeroInt, zeroString, zeroBool)

	// ---- Type inference and shorthand ----
	var inferred = "type inferred"
	short := "short declaration"
	var a, b = 1, 2
	c, d := "multiple", 3.5
	fmt.Println(inferred, short, a, b, c, d)

	// ---- Constants ----
	const pi float32 = 3.1415
	fmt.Println(pi)
}
