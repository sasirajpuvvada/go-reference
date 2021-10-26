package main

import "fmt"

func main() {
	fmt.Println("Structs in goloang")
	// no inheritance in go loang; No super or parent

	hitesh := User{"Hitesh", "hitesh@go.dev", true, 20}
	// fmt.Println(hitesh)
	// fmt.Printf("hitesh details are: %+v\n", hitesh)
	// fmt.Printf("Name is %v and E-mail is %v \n", hitesh.Name, hitesh.Email)
	hitesh.GetStatus()
	hitesh.NewEmail()
	fmt.Println(hitesh) // doesn't change original value bcz it creates copy of object
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("status is ", u.Status)
}

func (u User) NewEmail() {
	u.Email = "test@go.dev"
	fmt.Println(u)
}
