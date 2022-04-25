package controllers

import (
	"fmt"
	"http-api/dtos"
	"http-api/services"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type questionController struct {
	questionService services.IQuestionService
}

func QuestionController(questionService services.IQuestionService) *questionController {
	return &questionController{questionService}
}

func (ctrl *questionController) HandleReadQuestion(c *gin.Context) {
	startTime := time.Now()

	// Call questions service to retrieve all questions
	result, err := ctrl.questionService.FindAll()
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

func (ctrl *questionController) HandleCreateQuestion(c *gin.Context) {
	startTime := time.Now()

	// Question Input DTO
	var questionInput dtos.QuestionInput

	// Validate user input (DTO)
	bodyErr := c.ShouldBindJSON(&questionInput)
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

	// Call questions service to create user obj
	resultErr := ctrl.questionService.Create(questionInput, int64(authorId.(float64)))
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

func (ctrl *questionController) HandleReadDetailQuestion(c *gin.Context) {
	startTime := time.Now()

	// Get string id from url parameter
	rawId, exist := c.Params.Get("id")
	if !exist {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	// convert string to int
	id, conErr := strconv.ParseInt(rawId, 10, 64)
	if conErr != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// Call questions service to retrieve specific questions
	result, err := ctrl.questionService.FindById(id)
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
