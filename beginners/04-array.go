package main

import "fmt"

func main() {
	grades := [...]int{97, 85, 93}
	fmt.Printf("Grades: %v\n", grades)

	var students [3]string
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
