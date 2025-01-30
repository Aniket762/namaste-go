package main

import(
	"github.com/gin-gonic/gin"
	"log"
)

func HelloWorldHandler() gin.HandlerFunc{
	return func(c*gin.Context){
		c.JSON(200,gin.H{"status":"success","message":"Hello world"})
	}
}

func main(){
	router:=gin.Default()
	router.GET("/hello",HelloWorldHandler())
	err:=router.Run(":8080")
	if err!=nil{
		log.Fatal()
	}
}


