package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// basic routes
	r.GET("/hello", handleHelloRoute)
	r.GET("/hello/:id", handleDetailRoute)

	r.Run() // run on port 8080
}

func handleHelloRoute(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello There!"})
}

func handleDetailRoute(c *gin.Context) {
	startTime := time.Now()
	id := c.Param("id")
	showAll := c.Query("showAll")

	if showAll == "true" {
		var x []interface{}

		for i := 0; i < 5; i++ {
			x = append(x, gin.H{
				"index":                i,
				"address_id":           generateRandomString(48),
				"last_contact_address": generateRandomString(48),
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"status":      http.StatusOK,
			"data":        x,
			"path":        id,
			"timestamp":   time.Now(),
			"response_ns": time.Now().UnixNano() - startTime.UnixNano(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":      http.StatusOK,
		"data":        "Hello there, I am item " + id,
		"path":        id,
		"timestamp":   time.Now(),
		"response_ns": time.Now().UnixNano() - startTime.UnixNano(),
	})
}
