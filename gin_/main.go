package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Person struct {
	Name  string `json:"name"`
	Age string `json:"age"`
}
func main() {
	router:=gin.Default()


	router.GET("/user/:name/:age", func(context *gin.Context) {
		name:=context.Param("name")
		age:=context.Param("age")
		p:=Person{name,age}
		context.JSON(http.StatusOK,p)
	})



	router.Run(":8080")
}
