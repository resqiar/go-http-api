package tasks

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type taskController struct {
	taskService IService
}

func TaskController(taskService IService) *taskController {
	return &taskController{taskService}
}

func (ctrl *taskController) HandleReadTask(c *gin.Context) {
	startTime := time.Now()

	// Call tasks service to retrieve all tasks
	result, err := ctrl.taskService.FindAll()
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

func (ctrl *taskController) HandleCreateTask(c *gin.Context) {
	startTime := time.Now()

	// Task Input DTO
	var taskInput TaskInput

	// Validate user input (DTO)
	bodyErr := c.ShouldBindJSON(&taskInput)
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

	// Call tasks service to create user obj
	resultErr := ctrl.taskService.Create(taskInput, int64(authorId.(float64)))
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
