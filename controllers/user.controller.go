package controllers

import (
	"http-api/services"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type userCtrl struct {
	userService services.IUserService
}

func UserCtrl(userService services.IUserService) *userCtrl {
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

func (ctrl *userCtrl) HandleFindUserById(c *gin.Context) {
	startTime := time.Now()

	// Get id from url parameter
	rawId, exist := c.Params.Get("id")
	if !exist {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	id, conErr := strconv.ParseInt(rawId, 10, 64)
	if conErr != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// Call user service to query user based on given id
	result, err := ctrl.userService.FindById(id)
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
