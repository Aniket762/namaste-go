package main

import (
	"fmt"
	"sort"
)

// slices is like vector of cpp
// dynamic size
// only remove size from array and boom it's slice 😂
func main()  {
	fmt.Println("Getting into Slices!")
	var fruitList = []string {"Apple","Tomato","Peach"}
	fmt.Println(fruitList)
	fmt.Println(len(fruitList))

	// appending data
	fruitList = append(fruitList, "Mango","Banana")
	fmt.Println(fruitList)
	fmt.Println(len(fruitList))

	// type of ds, size of ds
	highScores := make([]int,4)
	highScores[0] = 1
	highScores[1] = 3
	highScores[2] = 2
	highScores[3] = 4

	// now you cannot use an index to place a value
	// but we can append

	highScores = append(highScores, 342,542,112,76)
	fmt.Println(highScores)
	
	// sorting
	sort.Ints(highScores)

	// delete values using highscores[i:j]
	// specify i and j as the indexes to cut off values
	var courses = [] string {"react","node","mongo","ts","flutter"}
	// index to delete from
	var index int = 2
	courses = append(courses[:index],courses[index+1:]... )
	fmt.Println(courses)
}