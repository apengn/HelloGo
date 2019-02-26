package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"os"
	"io"
	"fmt"
	"strings"
	"errors"
	_ "net/http/pprof"
)

//gin  redis   nginx  gorm etcd  kafka   es  kibane



type Person struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}


//参数校验
func Middleware(out io.Writer) gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println(context.Request.URL, "=====================")
		if strings.Contains(context.Request.URL.Path, "888") {
			context.JSON(http.StatusOK, "404 middleware")
			fmt.Println("xxxxxx")
			//权限验证不通过后 不需要执行以后的HandlerFunc
			context.Abort()
			return
		}
		context.Next()

	}
}

func main() {
	router := gin.New()

	router.Use(gin.Logger())

	router.Use(gin.Recovery())
	router.Use(Middleware(gin.DefaultErrorWriter))

	f, _ := os.Create("gin.log")

	gin.DefaultWriter = io.MultiWriter(f)

	v1 := router.Group("v1")

	v1.GET("/user/:name/:age", func(context *gin.Context) {
		name := context.Param("name")
		age := context.Param("age")
		p := Person{name, age}
		context.Error(errors.New(name))
		context.JSON(http.StatusOK, p)
	})

	v1.Handle("GET", "/test/:id", func(context *gin.Context) {
		id := context.Param("id")

		context.JSON(http.StatusOK, id)
	})

	router.Run(":8080")
}
