package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	// days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	// fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("index is %v, day is %v\n", index, day)
	// }

	rogueValue := 1

	for rogueValue < 10 {

		if rogueValue == 2 {
			goto lco
		}

		fmt.Println("Value is: ", rogueValue)
		rogueValue++

		if rogueValue == 5 {
			// break
			continue
		}
	}

lco:
	fmt.Println("Jumping!!")

}
