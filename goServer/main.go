package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request){
	// if not able to parse form
	if err := r.ParseForm(); err !=nil{
		fmt.Fprintf(w,"ParseForm() err:%v",err)
		return
	}
	fmt.Fprintf(w,"POST request successful")

	// assigning variable according to label value
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w,"Name= %s\n",name)
	fmt.Fprintf(w,"Address=%s\n",address)
}

func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello"{
		http.Error(w,"Status 404 Error",http.StatusNotFound)
		return 
	}

	if r.Method != "GET"{
		http.Error(w,"Method Not Supported",http.StatusNotFound)
		return 
	}
	
	fmt.Fprintf(w,"Hello World")
}

func main(){
	// fileServer will look for index.html from static directory
	fileServer :=http.FileServer(http.Dir("./static"))
	http.Handle("/",fileServer)

	http.HandleFunc("/form",formHandler)
	http.HandleFunc("/hello",helloHandler)

	fmt.Println("Server running at port 8080")

	if err := http.ListenAndServe(":8080",nil); err !=nil{
		log.Fatal(err)
	}
}