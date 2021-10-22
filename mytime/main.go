package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Time Study og go lang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	// Format for time
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.August, 20, 10, 5, 6, 0, time.UTC)
	fmt.Println(createdDate)
}
