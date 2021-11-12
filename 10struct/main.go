package main

import "fmt"

// capital letters are important
type User struct{
	Name string `default:"Aniket"`
	Email string `default:"aniketindian8@gmail.com"`
	Status bool `default:"false"`
	Age int	
}

// default values can be passed via constructor

func main()  {
	fmt.Println("Getting into Structs!")
	// no inheritance , no super or parent
	user := User{"Aniket","aniket@gmail.com",true,10}
	fmt.Println(user)
	// for value + parameter %+v
	fmt.Printf("%+v\n",user)

	// accessing single para
	fmt.Printf("%v",user.Name)
}