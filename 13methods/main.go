package main

import "fmt"

type User struct{
	Name string 
	Email string 
	Status bool 
	Age int	
}

// function wrapped in struct is method

// getter
func (u User) GetStatus()  {
	fmt.Println(u.Status)
}

// setter
func (u User) NewMail()  {
	u.Email = "test@google.com"
	fmt.Println(u.Email)
}

// these objects when passes in params
// are sent by reference ie copies hence 
// actual values aren't changed

func main()  {
	fmt.Println("Getting into Methods 🙏 ")
	user := User{"Aniket","aniket@gmail.com",true,10}
	fmt.Println(user)
	user.GetStatus()
	// for value + parameter %+v
	fmt.Printf("%+v\n",user)

	// accessing single para
	fmt.Printf("%v",user.Name)
}