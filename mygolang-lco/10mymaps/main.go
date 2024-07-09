package main

import "fmt"

func main() {
	fmt.Println("maps in golang")

	languages := make(map[string]string)

	languages["JS"] = "javascript"
	languages["RB"] = "ruby"
	languages["PY"] = "python"

	fmt.Println("list of all languages", languages)
	fmt.Println("JS shorts for: ", languages["JS"])

	delete(languages, "RB")
	fmt.Println("list of all languages", languages)

	//loops are interesting in Go

	for key, val := range languages {
		fmt.Printf("For key %v, value is %v", key, val)
	}



}