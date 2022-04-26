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
	userService     services.IUserService
}

func QuestionController(questionService services.IQuestionService, userService services.IUserService) *questionController {
	return &questionController{questionService: questionService, userService: userService}
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

func (ctrl *questionController) HandleUpdateQuestion(c *gin.Context) {
	startTime := time.Now()

	// Update Question Input DTO
	var updateInput dtos.UpdateQuestionInput

	// Validate user input (DTO)
	bodyErr := c.ShouldBindJSON(&updateInput)
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

	// Get target question
	targetQuestion, questionErr := ctrl.questionService.FindById(int64(updateInput.ID))
	targetAuthor, authorErr := ctrl.userService.FindById(int64(authorId.(float64)))
	if questionErr != nil || authorErr != nil {
		c.AbortWithStatus(400)
		return
	}

	// Make sure the one who update the data is the author
	if targetQuestion.AuthorID != targetAuthor.ID {
		c.AbortWithStatus(401)
		return
	}

	// Call questions service to update question
	resultErr := ctrl.questionService.UpdateQuestion(updateInput)
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

func (ctrl *questionController) HandleDeleteQuestion(c *gin.Context) {
	startTime := time.Now()

	// Delete Question Input DTO
	var deleteInput dtos.DeleteQuestionInput

	// Validate user input (DTO)
	bodyErr := c.ShouldBindJSON(&deleteInput)
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

	// Get target question
	targetQuestion, questionErr := ctrl.questionService.FindById(int64(deleteInput.ID))
	targetAuthor, authorErr := ctrl.userService.FindById(int64(authorId.(float64)))
	if questionErr != nil || authorErr != nil {
		c.AbortWithStatus(400)
		return
	}

	// Make sure the one who delete the data is the author
	if targetQuestion.AuthorID != targetAuthor.ID {
		c.AbortWithStatus(401)
		return
	}

	// Call questions service to delete question
	resultErr := ctrl.questionService.SoftDeleteQuestion(deleteInput)
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
