package main

import "fmt"

func main() {
	array()
	slice1()
	slice2()
	slice3()
	slice4()
}

func array() {
	// array has fixed size
	grades := [...]int{97, 85, 93}
	fmt.Printf("Grades: %v\n", grades)

	var students [3]string
	// zero-based index
	students[0] = "Lisa"
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Number of students: %v\n", len(students))

	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1, 0, 0}
	identityMatrix[1] = [3]int{0, 1, 0}
	identityMatrix[2] = [3]int{0, 0, 1}
	fmt.Println(identityMatrix)

	// array is copied by value not reference!
	copyOfGrades := grades
	copyOfGrades[0] = 88
	fmt.Println(grades)
	fmt.Println(copyOfGrades)

	// use pointer to copy by reference
	pointerToGrades := &grades
	pointerToGrades[0] = 99
	fmt.Println(grades)
	fmt.Println(*pointerToGrades)
}

func slice1() {
	a := []int{1, 2, 3}
	// slice is copied by reference
	b := a
	b[0] = 5
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("Length: %v\n", len(a))
	// length of an underlying array
	fmt.Printf("Capacity: %v\n", cap(a))
}

func slice2() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[:]
	c := a[3:]
	d := a[:6]
	e := a[3:6]
	a[5] = 66
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Printf("Length: %v\n", len(e))
	fmt.Printf("Capacity: %v\n", cap(e))
}

func slice3() {
	// more efficient if slice may grow pass initial underlying array
	a := make([]int, 3, 100)
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))
	a = append(a, 1)
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))
	a = append(a, 2, 3, 4, 5)
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))
	a = append(a, []int{6, 7, 8}...)
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))
}

func slice4() {
	a := []int{1, 2, 3, 4, 5}
	b := a[1:]                   // remove the first item
	c := a[:len(a)-1]            // remove the last item
	d := append(a[:2], a[3:]...) // remove the 3rd item
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	// BUT this changes a!
	fmt.Println(a)
}
