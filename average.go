package main

import "fmt"

func main() {

	fmt.Println("Name: ")
	var name string
	fmt.Scanln(&name)

	fmt.Println("Welcome, ", name)

	fmt.Println("**********")

	fmt.Println("Your math grade: ")
	var mathGrade int
	fmt.Scanln(&mathGrade)

	fmt.Println("Your physics grade: ")
	var physicsGrade int
	fmt.Scanln(&physicsGrade)

	fmt.Println("your Physical education grade: ")
	var physicalGrade int
	fmt.Scanln(&physicalGrade)

	fmt.Println("Your chemistry grade: ")
	var chemistryGrade int
	fmt.Scanln(&chemistryGrade)

	array := []int{mathGrade, physicsGrade, physicalGrade, chemistryGrade}

	n := 4

	sum := 0

	for i := 0; i < n; i++ {

		sum += (array[i])
	}

	avg := (float64(sum)) / (float64(n))

	fmt.Println("Average: ", avg)

}
