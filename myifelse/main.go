package main

import "fmt"

func main() {
	fmt.Println("If else in go lang")

	loginCount := 10

	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "Something else"
	} else {
		result = "Exactly 10"
	}

	fmt.Println(result)

	// assigning on the condition
	if num := 3; num < 10 {
		fmt.Println("Num less than 10")
	} else {
		fmt.Println("Num is gt than 10")
	}

	// if err != nil {

	// }

}
