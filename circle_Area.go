package main

import "fmt"

func main() {
	Pi := 3.14
	
	var radius, Area float64
	
	fmt.Println("Enter radius: ")
	fmt.Scanln(&radius)

	Area = (Pi) * (radius * radius)
	fmt.Println("Circle Area: ", Area)
}
