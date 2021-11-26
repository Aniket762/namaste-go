package main

import (
	"fmt"
	"net/http"

	"github.com/Aniket762/namaste-go/dbapi/router"
)

func main()  {
	fmt.Println("Getting into building APIs with Mongo as DB 🙏")
	r := router.Router()
	http.ListenAndServe(":8081",r)
}