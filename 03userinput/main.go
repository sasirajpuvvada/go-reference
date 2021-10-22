package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating")

	// comma ok || error ok syntax

	input, _ := reader.ReadString('\n')
	fmt.Print(input)
	fmt.Printf("Thanks for rating %T\n", input)
}
