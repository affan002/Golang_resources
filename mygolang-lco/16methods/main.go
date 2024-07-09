package main

import "fmt"

func main() {
	fmt.Println("structs in Go")
	// no inheritance in golang; no super or parent
	affan := User{"Affan", "affan@go.dev", true, 20}
	fmt.Println(affan)
	fmt.Printf("affan's details are %+v\n", affan)
	fmt.Printf("Name is %v, Email is %v\n", affan.Name, affan.Email)
	affan.GetStatus()
	affan.NewMail()
	fmt.Printf("Name is %v, Email is %v\n", affan.Name, affan.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail(){
	u.Email = "test@go.dev"
	fmt.Println("email of this user is: ", u.Email)
}