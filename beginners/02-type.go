package main

import (
	"fmt"
	"strconv"
)

func main() {
	booleanType()
	numericType()
	arithmeticOperator()
	bitwiseOperator()
	stringType()
}

func booleanType() {
	var thisIsBoolean bool
	isEqual := 1 == 1
	fmt.Printf("%v %T\n", thisIsBoolean, thisIsBoolean)
	fmt.Printf("%v %T\n", isEqual, isEqual)
}

func numericType() {
	// https://go.dev/ref/spec#Numeric_types
	var nInt int = 42
	var nInt8 int8 = 42
	var nInt16 int16 = 42
	var nInt32 int32 = 42
	var nInt64 int64 = 42
	fmt.Printf("%v %T - on this machine int = int%v\n", nInt, nInt, strconv.IntSize)
	fmt.Printf("%v %T\n", nInt8, nInt8)
	fmt.Printf("%v %T\n", nInt16, nInt16)
	fmt.Printf("%v %T\n", nInt32, nInt32)
	fmt.Printf("%v %T\n", nInt64, nInt64)

	var nFloat32 float32 = 3.14
	nFloat64 := 2.1e14
	fmt.Printf("%v %T\n", nFloat32, nFloat32)
	fmt.Printf("%v %T\n", nFloat64, nFloat64)

	var nComplex64 complex64 = complex(5, 12)
	var nComplex128 complex128 = 1 + 2i
	fmt.Printf("%v %T\n", nComplex64, nComplex64)
	fmt.Printf("%v %T, %v %T\n", real(nComplex64), real(nComplex64), imag(nComplex64), imag(nComplex64))
	fmt.Printf("%v %T\n", nComplex128, nComplex128)
	fmt.Printf("%v %T, %v %T\n", real(nComplex128), real(nComplex128), imag(nComplex128), imag(nComplex128))
}

func arithmeticOperator() {
	a := 10
	b := 3
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
}

func bitwiseOperator() {
	var a uint64 = 10   // 1010
	var b uint64 = 3    // 0011
	fmt.Println(a & b)  // 0010 = 2  (AND)
	fmt.Println(a | b)  // 1011 = 11 (OR)
	fmt.Println(a ^ b)  // 1001 = 9  (XOR)
	fmt.Println(^a)     // 1101      (NOT = 1111 XOR a)
	fmt.Println(a &^ b) // 0100 = 8  (AND NOT = NOT OR)
	fmt.Println(8 << 3) // 0100_0000 = 64
	fmt.Println(8 >> 3) // 0000_0001 = 1
}

func stringType() {
	s := "This is UTF-8" + " 😊"
	fmt.Println(s)
	b := []byte(s) // a slice of bytes
	fmt.Printf("%v %T\n", b, b)
	fmt.Printf("%v %T\n", s[0], s[0])

	r := []rune("This is UTF-32 😊")
	var c rune = 'T'
	fmt.Println(string(r))
	fmt.Printf("%v %T\n", r, r)
	fmt.Printf("%v %T\n", c, c)
}
