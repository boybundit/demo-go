package main

import (
	"fmt"
)

const a int16 = 27

const (
	_               = iota // 0
	catSpecialist          // 1
	dogSpecialist          // 2
	snakeSpecialist        // 3
)

// https://go.dev/doc/effective_go
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
)

const (
	isAdmin = 1 << iota
	isHeadquarter
	canSeeFinacials
	canSeeAfrica
	canSeeAsia
	canSeeEurope
)

func main() {
	fmt.Printf("%v %T\n", a, a)
	// Typed constant
	const a int = 42
	// Shadowing
	fmt.Printf("%v %T\n", a, a)
	// Untyped constant works like c macro
	const b = 12
	fmt.Println(12.1 + b)

	var specialistType int
	fmt.Println(specialistType == catSpecialist)

	fileSzie := 4_000_000_000
	fmt.Printf("%.2fGB\n", float64(fileSzie)/GB)

	var roles byte = isAdmin | canSeeFinacials | canSeeEurope
	fmt.Printf("%b\n", roles)
	fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin)
	fmt.Printf("Is Headquater? %v\n", isHeadquarter&roles == isHeadquarter)
}
