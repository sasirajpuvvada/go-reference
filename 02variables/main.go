package main

import "fmt"

const LoginToken string = "jn2idn349c" //Public  variable

func main() {
	var username string = "sasi"
	fmt.Println(username)
	fmt.Printf("variable is type of %s \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is type of %T \n", isLoggedIn)

	var anotherDeclaration = 100
	fmt.Println(anotherDeclaration)

	someVar := 200
	fmt.Println(someVar)

	fmt.Printf("Login Token %s \n", LoginToken)
}
