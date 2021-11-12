package main

import (
	"fmt"
)

func main()  {
	fmt.Println("Getting into maps!")
	languages := make(map[string]string)
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	fmt.Println(languages)
	fmt.Println(languages["JS"])

	// deleting value
	delete(languages,"RB")
	fmt.Println(languages)

	// looping over map
	for key,value := range languages{
		fmt.Printf("For key %v, value is %v \n",key,value)
	}

	// Note:
	// with _ you can ignore anything
}