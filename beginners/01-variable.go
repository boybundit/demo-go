package main

import (
	"fmt"     // https://pkg.go.dev/fmt
	"strconv" // https://pkg.go.dev/strconv
)

// We can group var block for readability
var (
	title = "Learn Go Programming - Golang Tutorial for Beginners"
	// Acronym should be UPPERCASE
	linkURL = "https://www.youtube.com/watch?v=YS4e4q9oBaU&t=1843s"
)

var ThisIsExported int = 42
var thisHasPackageScope int = 42

func main() {
	variableDeclaration()
	variableScope()
	typeConversion()
}

func variableDeclaration() {
	// This will get a default value
	var thisIsZero int
	// We can let compiler infer a default type from an initial value
	thisIsInt := 42
	thisIsFloat64 := 42.0
	// We want a specific type instead of a default one
	var thisIsFloat32 float32 = 42

	fmt.Printf("%v %T\n", thisIsZero, thisIsZero)
	fmt.Printf("%v %T\n", thisIsInt, thisIsInt)
	fmt.Printf("%v %T\n", thisIsFloat64, thisIsFloat64)
	fmt.Printf("%v %T\n", thisIsFloat32, thisIsFloat32)
}

func variableScope() {
	var thisHasBlockScope int = 42
	fmt.Printf("%v from block scope\n", thisHasBlockScope)

	// Shadowing: a variable with more specific scope takes a precedence
	fmt.Printf("%v from package scope\n", thisHasPackageScope)
	var thisHasPackageScope = 43
	fmt.Printf("%v from block scope by shadowing\n", thisHasPackageScope)
}

func typeConversion() {
	var i float64 = 42.5
	var j int = int(i)
	var k string = strconv.Itoa(j)
	fmt.Printf("%v %T\n", i, i)
	fmt.Printf("%v %T\n", j, j)
	fmt.Printf("%v %T\n", k, k)
}
