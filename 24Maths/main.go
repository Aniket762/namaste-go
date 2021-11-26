package main

import (
	"fmt"
	"math/big"

	// "math/rand"
	"crypto/rand"
)

func main()  {
	fmt.Println("Getting into Maths functions and variables 🙏")
	/*
	var One int =3
	var Two float64 =5.3
	
	fmt.Println(One + int(Two)) // loosing precision over here!
	*/

	// random number math/rand 
	// driven by seed so changing the value of seed
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(100))

	// random numbers from cryptp/rand
	// doesn't depend on seed
	RandomNum , _ := rand.Int(rand.Reader, big.NewInt(100))
	fmt.Println(RandomNum)
}