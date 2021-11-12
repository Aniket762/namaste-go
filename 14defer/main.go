package main

import "fmt"

func main()  {
	fmt.Println("Getting into defer 🙏 ")

	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	myDefer()
}


func myDefer()  {
	for i:=0;i<5;i++{
		defer fmt.Println(i)
	}
}
