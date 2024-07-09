package main

import "fmt"

func main() {
	fmt.Println("structs in Go")
	// no inheritance in golang; no super or parent
	affan := User{"Affan", "affan@go.dev", true, 20}
	fmt.Println(affan)
	fmt.Printf("affan's details are %+v\n", affan)
	fmt.Printf("Name is %v, Email is %v", affan.Name, affan.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
