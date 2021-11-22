package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main()  {
	fmt.Println("Getting started with Get Requests 🙏")
	PerformGetRequest()
}

func PerformGetRequest()  {
	const uri = "http://localhost:8080"
	response,err := http.Get(uri) 
	 if err != nil{
		 panic(err)
	 }

	defer  response.Body.Close()

	 fmt.Println(response.StatusCode)
	 fmt.Println("Content length is: ", response.ContentLength)

	content , _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))	

	// another method to show the content 
	// using strings package
	var responseString strings.Builder
	byteCount,_ := responseString.Write(content)
	fmt.Println(byteCount)
	fmt.Println(responseString.String())

	
}