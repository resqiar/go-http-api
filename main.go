package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	// basic route
	r.GET("/hello", handleHelloRoute)

	r.Run() // run on port 8080
}

func handleHelloRoute(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello There!"})
}
