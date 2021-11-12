package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
	fmt.Println("Please rate our pizza between 1 and 5")

	reader := bufio.NewReader(os.Stdin)
	input,errors := reader.ReadString('\n')

	fmt.Println("Thanks for rating, ", input)
	// fmt.Printf("The type of rating %T,",input) // string
	if errors != nil {
		fmt.Println(errors)
	}

	// coverting string to number

	// strings.TrimSpace is needed to remove the trailing character \n
	numRating,errors2 := strconv.ParseFloat(strings.TrimSpace(input),64) 
	
	if errors2 != nil {
		fmt.Println(errors2)
	}
	
	numRating2 := numRating+1
	fmt.Println(numRating2) 
}