package main

import "fmt"

func main()  {
	fmt.Println("Getting into Poiners!")
	
	// if we don't initialize pointer with anything it will be <nill>
	var ptr *int 
	fmt.Println("Value of pointer is ", ptr)

	// creating a pointer referencing to memory
	myNumber := 23
	var ptr2 = &myNumber
	fmt.Println("The address of myNumber is: ptr2 ", ptr2)
	fmt.Println("The value at myNumber is: *ptr2 ", *ptr2)
	fmt.Println("The address of ptr2 is: &ptr2 ", &ptr2)

	*ptr2  = *ptr2 +2 
	fmt.Println("New after *ptr2  = *ptr2 +2 of myNumber is: ", myNumber)

	// Note:
	// Whenever we use pointer we directly point to that value
	// we don't create copies
}