package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	a :="aloo"
	fmt.Println(a)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter")

	// comma ok || err ok
	// it's like try catch where we take
	// the input or error or both.
	// incase we wanna ignore one we put _

	// ReadString('char') shows till when
	// the input will read. Here when enter
	// is pressed it will stop taking in the
	// input
	input , _ := reader.ReadString('\n')
	fmt.Println("Entered val:",input)

	// over here reader is a pointer
}