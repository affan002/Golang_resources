package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers lesson")

	// var ptr *int
	// fmt.Println("The value of pointer is", ptr)

	myNumber := 23
	var ptr = &myNumber
	fmt.Println("The actual value of pointer is", ptr)
	fmt.Println("The actual value of pointer is", *ptr)

	*ptr = *ptr + 2
	fmt.Println("new value is", myNumber)

}