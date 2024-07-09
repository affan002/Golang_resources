package main

import "fmt"

const LoginToken string = "shsfhhsla" //public

func main() {
	var username string = "Affan"
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type: %T \n", isLoggedIn)

	var samllVal int = 256
	fmt.Println(samllVal)
	fmt.Printf("variable is of type: %T \n", samllVal)

	var samllFloat float64 = 255.4556662222
	fmt.Println(samllFloat)
	fmt.Printf("variable is of type: %T \n", samllFloat)

	// default values and aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("variable is of type: %T \n", anotherVariable)

	var anotherString string
	fmt.Println(anotherString)
	fmt.Printf("variable is of type: %T \n", anotherString)


	// implicit type
	var website = "affan.com"
	fmt.Println(website)

	// no var style
	numberOfUsers := 300000
	fmt.Println(numberOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("variable is of type: %T \n", LoginToken)

	


}
