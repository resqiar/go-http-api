package controllers

import (
	"fmt"
	"http-api/dtos"
	"http-api/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type answerController struct {
	answerService services.IAnswerService
}

func AnswerController(answerService services.IAnswerService) *answerController {
	return &answerController{answerService}
}

func (ctrl *answerController) HandleReadAnswers(c *gin.Context) {
	startTime := time.Now()

	// Call answer service to retrieve all questions
	result, err := ctrl.answerService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":      http.StatusOK,
		"data":        result,
		"timestamp":   time.Now(),
		"response_ms": time.Now().UnixMilli() - startTime.UnixMilli(),
	})
}

func (ctrl *answerController) HandleCreateAnswer(c *gin.Context) {
	startTime := time.Now()

	// Answer Input DTO
	var answerInput dtos.AnswerInput

	// Validate input (DTO)
	bodyErr := c.ShouldBindJSON(&answerInput)
	if bodyErr != nil {
		errorMessage := []string{}
		for _, e := range bodyErr.(validator.ValidationErrors) {
			errorMessage = append(errorMessage, fmt.Sprintf("Error on field %s, reason %s", e.Field(), e.ActualTag()))
		}
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": errorMessage})
		return
	}

	// Get author id from JWT Guard context
	authorId, _ := c.Get("user_id")

	// Call answer service to create user obj
	resultErr := ctrl.answerService.Create(answerInput, int64(authorId.(float64)))
	if resultErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": resultErr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":      http.StatusOK,
		"timestamp":   time.Now(),
		"response_ms": time.Now().UnixMilli() - startTime.UnixMilli(),
	})
}
