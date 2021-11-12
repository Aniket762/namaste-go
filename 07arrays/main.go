package main

import "fmt"

func main()  {
	fmt.Println("Getting into array!")

	var food[4] string 

	food[0] = "biriyani"
	food[1]="momo"
	food[3] = "dosa"

	fmt.Println(food)
	// len(food) gives 4 bcz even if we don't enter a value it keeps a reserved space

	var vegList = [5]string {"potato","beans","mushroom"}
	fmt.Println(vegList)


}