package main

import "fmt"

// when first letter is capital it is public
// variable

// defining immutable variables
// capital L means public variable
const LoginToken string = "awhnswhjbnshwbnsh"

func main()  {
	var username string = "Aniket762"
	fmt.Println(username)
	// %T for type 
	fmt.Printf("Variable is of type: %T \n",username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n",isLoggedIn)

	var smallVal uint8 = 12
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n",smallVal)

	// default value of an int is 0

	// implicit type: not defining type also works
	var dob = 672001
	fmt.Println(dob)

	// no var style - can only be defined under a method or func
	// := is volorous operator
	numberOfUser := 3000
	fmt.Println(numberOfUser) 
}