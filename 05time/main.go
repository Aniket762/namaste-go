package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println("Welcome to time study of go")
	presentTime := time.Now()
	fmt.Println(presentTime)

	// the time mentioned below is the std format as in docs
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2021,time.November,6,21,51,10,6,time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))

} 

/*
Extras:
To build exes for different os, say I am using MAC
i> Windows build: GOOS="windows" go build
ii> Mac build: go build
iii> Linux build: GOOS="linux" go build
*/