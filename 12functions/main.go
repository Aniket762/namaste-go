package main

import "fmt"

func main()  {
	greeter()
	fmt.Println("Getting into functions in go")

	result := adder(3,5)
	fmt.Println(result)

	fmt.Println(proAdder(1,2,3,4,5))

	result2 , message := adder2(2,3)
	fmt.Println(result2)
	fmt.Println(message)
}

func greeter()  {
	fmt.Println("Namastey 🙏")
}

// here the int shows the return type of func
func adder(valOne int,valTwo int) int  {
	
	return valOne + valTwo
}

// if you want to return 2 things or more
func adder2(valOne int,valTwo int) (int,string ) {
	
	return valOne + valTwo , "Hehe"
}


// it is used when we don't know how many 
// values we should add 
func proAdder(values ...int) int {
	total :=0

	for _,val :=range values {
		total +=val
	}

	return total
}

/*
1. We are not allowed to write functions inside a function
2. Anonymous functions exists and you can define  them without naming the func
*/