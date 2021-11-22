package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main()  {
	fmt.Println("Getting started with Post Requests 🙏")
	PerformPostJsonRequest()
}

func PerformPostJsonRequest()  {
	const uri = "http://localhost:8080/post"

	// dummy json payload
	requestBody := strings.NewReader(`
		{
			"fname":"Aniket",
			"lname":"Pal",
			"pin":91
		}
	`)

	// url,type of content, content
	response,err:=http.Post(uri, "application/json",requestBody)
	
	if	err != nil{
		panic(err)
	}

	defer response.Body.Close()
	content , _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}