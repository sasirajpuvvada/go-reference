package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This is need to go in a file"

	file, err := os.Create("./myfile.txt")

	// if err != nil {
	// 	panic(err)
	// }
	checkNilError(err)

	length, err := io.WriteString(file, content)
	checkNilError(err)

	fmt.Println("length is ", length)
	defer file.Close()
	readFile("./myfile.txt")
}

func readFile(filename string) {
	dataByte, err := ioutil.ReadFile(filename)
	checkNilError(err)

	fmt.Println("Text data inside file is \n", string(dataByte))

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
