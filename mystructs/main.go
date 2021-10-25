package main

import "fmt"

func main() {
	fmt.Println("Structs in goloang")
	// no inheritance in go loang; No super or parent

	hitesh := User{"Hitesh", "hitesh@go.dev", true, 20}
	fmt.Println(hitesh)
	fmt.Printf("hitesh details are: %+v\n", hitesh)
	fmt.Printf("Name is %v and E-mail is %v \n", hitesh.Name, hitesh.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
