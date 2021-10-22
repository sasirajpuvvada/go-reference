package main

import "fmt"

func main() {
	fmt.Println("Welcome to poniters")

	// var ptr *int
	// fmt.Println("Value of Poniter is", ptr)

	myNumer := 23

	//  reference
	var ptr = &myNumer
	fmt.Println("Value of actual pointer is", ptr)
	fmt.Println("Value of actual pointer is", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New value is:", myNumer)
}
