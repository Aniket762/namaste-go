package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main()  {
	fmt.Println("Getting into file system 🙏")
	content :="Some random gibrish content!"

	file,err := os.Create("./main.txt")
	checkNilErr(err)

	length,err := io.WriteString(file,content)
	checkNilErr(err)

	fmt.Println("Length is: ",length)
	defer file.Close()

	// will get binary
	readFile("./main.txt")

}

func readFile(filename string)  {
	databyte,err := ioutil.ReadFile(filename)
	if err != nil{
		panic(err)
	}

	fmt.Println("Text data inside file is \n",databyte)
	fmt.Println("Text data inside file is in string format \n",string(databyte))

}

func checkNilErr(err error)  {
	if err != nil{
		panic(err)
	}
}



/* 
file creation is the OS operation
but reading isn't 
*/