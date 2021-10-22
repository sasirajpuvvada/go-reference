package main

import "fmt"

func main() {
	fmt.Println("Welecom to array in go lang")

	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Orange"
	// fruitList[2] = "Apple"
	fruitList[3] = "Cherry"

	fmt.Println("Frulit list is ", fruitList)
	fmt.Println("Frulit list is ", len(fruitList))

	var cars = [3]string{"amaze", "xuv", "city"}
	fmt.Println("cars list is ", len(cars))
}
