package main

import "fmt"

func main() {
	fmt.Println("functions")
	greeter()

	result := add(3, 4)
	fmt.Println("result is: ", result)

	proresult := proAdd(2,5,7,8)
	fmt.Println("pro res is: ", proresult)
}

func add(val1 int, val2 int) int {
	return val1 + val2
}

func proAdd(values ...int) int {
	total := 0
	
	for _, value := range values {
		total += value
	}

	return total
}

func greeter() {
	fmt.Println("Hello from golang")
}