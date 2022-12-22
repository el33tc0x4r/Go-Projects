package main

import "fmt"

func main() {
	var primeNumber, count int

	fmt.Println("The number you want to check: ")
	fmt.Scanln(&primeNumber)

	if primeNumber <= 0 {
		fmt.Println("Please enter a positive value!")
		return
	}

	for i := 2; i < primeNumber/2; i++ {
		if primeNumber%i == 0 {
			count++
			break
		}
	}
	if count == 0 && primeNumber != 1 {
		fmt.Println(primeNumber, "prime number")
	} else {
		fmt.Println(primeNumber, "not prime number")
	}

}
