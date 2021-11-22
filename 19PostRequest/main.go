package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main()  {
	fmt.Println("Getting started with Post Requests 🙏")
	PerformPostJsonRequest()
	PerformPostFormRequest()
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

func PerformPostFormRequest()  {
	const uri = "http://localhost:8080/postform"

	data := url.Values{}
	data.Add("fname","Aniket")
	data.Add("lname","Pal")
	
	response, err := http.PostForm(uri,data)
	if err != nil{
		panic(err)
	}

	defer response.Body.Close()

	content , _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

