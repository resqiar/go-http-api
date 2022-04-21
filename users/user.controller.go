package users

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type userCtrl struct {
	userService IUserService
}

func UserCtrl(userService IUserService) *userCtrl {
	return &userCtrl{userService: userService}
}

func (ctrl *userCtrl) HandleReadUsers(c *gin.Context) {
	startTime := time.Now()
	result, err := ctrl.userService.FindAll()
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

func (ctrl *userCtrl) HandleCreateUser(c *gin.Context) {
	startTime := time.Now()

	var userInput UserInput

	bodyErr := c.ShouldBindJSON(&userInput)
	if bodyErr != nil {
		errorMessage := []string{}
		for _, e := range bodyErr.(validator.ValidationErrors) {
			errorMessage = append(errorMessage, fmt.Sprintf("Error on field %s, reason %s", e.Field(), e.ActualTag()))
		}
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": errorMessage})
		return
	}

	result, resultErr := ctrl.userService.Create(userInput)
	if resultErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": resultErr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":      http.StatusOK,
		"data":        result,
		"timestamp":   time.Now(),
		"response_ms": time.Now().UnixMilli() - startTime.UnixMilli(),
	})
}
