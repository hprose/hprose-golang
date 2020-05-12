package main

import (
	"github.com/lanfengye2008/hprose-golang/rpc"
	"gopkg.in/gin-gonic/gin.v1"
)

func hello(name string) string {
	return "Hello " + name + "!"
}

func main() {
	service := rpc.NewHTTPService()
	service.AddFunction("hello", hello)
	router := gin.Default()
	router.Any("/", func(c *gin.Context) {
		service.ServeHTTP(c.Writer, c.Request)
	})
	router.Run(":8080")
}
