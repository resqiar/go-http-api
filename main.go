package main

import (
	"http-api/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v1 := r.Group("v1")
	// basic routes
	v1.GET("/hello", handler.HandleHelloRoute)
	v1.GET("/hello/:id", handler.HandleDetailRoute)
	v1.POST("/hello", handler.HandlePostHello)

	r.Run() // run on port 8080
}
