package main

import (
	"fmt"
	"net/url"
)
const gurl string = "https://github.dev:3000/learn?coursename=golang&username=Aniket762"

func main()  {
	fmt.Println("Getting into URLs 🙏")
	fmt.Println(gurl)

	// parsing
	result,_ := url.Parse(gurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("Type: %T\n",qparams)
	
	fmt.Println(qparams["coursename"])

	// getting the params
	for _,val := range qparams{
		fmt.Println(val)
	}

	// imp! the & is important
	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=Aniket",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)

}