package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// basic route
	r.GET("/hello", handleHelloRoute)

	r.Run() // run on port 8080
}

func handleHelloRoute(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello There!"})
}
