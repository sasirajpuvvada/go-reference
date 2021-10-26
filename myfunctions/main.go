package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()
	greeterTwo()

	result := adder(3, 5)
	fmt.Println("result is ", result)

	result2, _ := proAdder(3, 5, 6, 6, 8)
	fmt.Println("result is ", result2)
}

func adder(valOne, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "total val"
}

func greeterTwo() {
	fmt.Println("Another method")
}

func greeter() {
	fmt.Println("Namaste from golang")
}
