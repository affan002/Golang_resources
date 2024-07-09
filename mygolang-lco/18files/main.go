package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Files in golang")
	content := "This needs to go in file"

	file, err := os.Create("./mylcogofile.txt")

	// if err!=nil {
	// 	panic(err)
	// }
	checkNilError(err)

	length, err := io.WriteString(file, content)
	checkNilError(err)

	fmt.Println("length is: ", length)

	defer file.Close()

	readFile("./mylcogofile.txt")
}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)
	checkNilError(err)

	fmt.Println("The data in the file is\n", string(databyte))

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
