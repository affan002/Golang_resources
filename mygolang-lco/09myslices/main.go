package main

import (
	"fmt"
	"sort"
)


func main() {
	fmt.Println("welcome to slice lesson")
	var fruitList = []string{"apple", "tomato", "peach"}
	fmt.Printf("Type of fruit list, %T\n", fruitList)

	fruitList = append(fruitList, "mango", "banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])   //same as slicing in python
	fmt.Println(fruitList)

	highScore := make([]int, 4)
	highScore[0] = 234
	highScore[1] = 945
	highScore[2] = 465
	highScore[3] = 867
	// highScore[4] = 777

	highScore = append(highScore, 555, 666, 321)

	fmt.Println(highScore)

	sort.Ints(highScore)
	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore))

	//how to remove a value from slice based on index

	var courses = []string{"reactJS", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}