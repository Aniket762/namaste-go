package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://github.com"

func main()  {
	fmt.Println("Getting into handling web request! 🙏")

	response, err := http.Get(url)

	if err != nil{
		panic(err)
	}

	fmt.Printf("Response is of type: %T \n",response)
	// fmt.Println(*response)

	databytes,err := ioutil.ReadAll(response.Body)

	if err != nil{
		panic(err)
	}

	content := string(databytes)
	fmt.Println(content)

	// always close the connection after you do a get request!
	defer response.Body.Close()	
}