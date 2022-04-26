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
	answerService   services.IAnswerService
	questionService services.IQuestionService
	userService     services.IUserService
}

func AnswerController(answerService services.IAnswerService, questionService services.IQuestionService, userService services.IUserService) *answerController {
	return &answerController{answerService: answerService, questionService: questionService, userService: userService}
}

func (ctrl *answerController) HandleReadAnswers(c *gin.Context) {
	startTime := time.Now()

	// Call answer service to retrieve all answers
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

	// Verify target question
	_, questionErr := ctrl.answerService.FindById(int64(answerInput.QuestionID))
	if questionErr != nil {
		c.AbortWithStatus(400)
		return
	}

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

func (ctrl *answerController) HandleUpdateAnswer(c *gin.Context) {
	startTime := time.Now()

	// Update Input DTO
	var updateInput dtos.UpdateAnswerInput

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

	// Get target answer
	targetAnswer, answerErr := ctrl.answerService.FindById(int64(updateInput.ID))
	targetAuthor, authorErr := ctrl.userService.FindById(int64(authorId.(float64)))
	if answerErr != nil || authorErr != nil {
		c.AbortWithStatus(400)
		return
	}

	// Make sure the one who update the data is the author
	if targetAnswer.AuthorID != targetAuthor.ID {
		c.AbortWithStatus(401)
		return
	}

	// Call answers service to update answer
	resultErr := ctrl.answerService.UpdateAnswer(updateInput)
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

func (ctrl *answerController) HandleDeleteAnswer(c *gin.Context) {
	startTime := time.Now()

	// Delete Input DTO
	var deleteInput dtos.DeleteAnswerInput

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

	// Get target answer
	targetAnswer, answerErr := ctrl.answerService.FindById(int64(deleteInput.ID))
	targetAuthor, authorErr := ctrl.userService.FindById(int64(authorId.(float64)))
	if answerErr != nil || authorErr != nil {
		c.AbortWithStatus(400)
		return
	}

	// Make sure the one who delete the data is the author
	if targetAnswer.AuthorID != targetAuthor.ID {
		c.AbortWithStatus(401)
		return
	}

	// Call answer service to delete answer
	resultErr := ctrl.answerService.SoftDeleteAnswer(deleteInput)
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
