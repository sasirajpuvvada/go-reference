package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to video on slices")

	var fruitList = []string{"Apple", "Tomato", "grapes"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "cherry")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 24
	highScores[1] = 21
	highScores[2] = 22
	highScores[3] = 12
	// highScores[4] = 14

	highScores = append(highScores, 555, 666)

	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))
}
