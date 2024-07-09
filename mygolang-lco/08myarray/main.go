package main

import "fmt"

func main() {
	fmt.Println("welcome to the array lesson")

	var fruitList [4]string
	fruitList[0] = "apple"
	fruitList[1] = "tomato"
	fruitList[3] = "peach"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("The length of this list is: ", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("veggie list: ", vegList)
}