package main

import "fmt"


func main() {
	fmt.Println("welcome to loops in Go")

	days := []string{"sunday", "tuesday", "wednesday", "friday", "saturday"}

	// for d := 0; d<len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for index, day := range days {
		fmt.Printf("index is %v, the day is %v \n", index, day)
	}

	rogueValue := 1
	for rogueValue<10{

		if rogueValue == 2{
			goto affan
		}

		if rogueValue == 5 {
			// break
			rogueValue++
			continue
		}
		
		fmt.Println("Value is:", rogueValue)
		rogueValue++
	}

	affan:
		fmt.Println("jumping to goto statement")
}