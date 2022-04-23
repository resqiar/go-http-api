package users

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type userCtrl struct {
	userService IUserService
}

func UserCtrl(userService IUserService) *userCtrl {
	return &userCtrl{userService}
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

func (ctrl *userCtrl) HandleFindUserByUsername(c *gin.Context) {
	startTime := time.Now()

	// Get username from url parameter
	input, exist := c.Params.Get("username")
	if !exist {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	// Call users servic to query user based on given username
	result, err := ctrl.userService.FindByUsername(input)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
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
