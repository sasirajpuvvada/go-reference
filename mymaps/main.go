package main

import "fmt"

func main() {

	languages := make(map[string]string)

	languages["JS"] = "javascript"
	languages["py"] = "python"
	languages["rb"] = "ruby"
	fmt.Println(languages)
	fmt.Println(languages["py"])

	delete(languages, "rb")
	fmt.Println(languages)

	// loops
	for key, value := range languages {
		fmt.Printf("Key is %v, value is %v\n", key, value)
	}

}
