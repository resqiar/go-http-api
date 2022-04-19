package handler

import (
	"fmt"
	"net/http"
	"time"

	"http-api/structs"
	"http-api/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func HandleHelloRoute(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello There!"})
}

func HandleDetailRoute(c *gin.Context) {
	startTime := time.Now()
	id := c.Param("id")
	showAll := c.Query("showAll")

	if showAll == "true" {
		var x []interface{}

		for i := 0; i < 5; i++ {
			x = append(x, gin.H{
				"index":                i,
				"address_id":           utils.GenerateRandomString(48),
				"last_contact_address": utils.GenerateRandomString(48),
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

func HandlePostHello(c *gin.Context) {
	var body structs.PostHello
	err := c.ShouldBindJSON(&body)

	if err != nil {
		errorMessage := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage = append(errorMessage, fmt.Sprintf("Error on field %s, reason %s", e.Field(), e.ActualTag()))
		}
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": errorMessage})
		return
	}

	if len(body.Id) == 0 || len(body.Title) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Id and title cannot be empty",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":    body.Id,
		"title": body.Title,
	})
}
