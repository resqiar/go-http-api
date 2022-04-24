package auth

import (
	"fmt"
	"http-api/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ILoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService ILoginService
	usersService services.IUserService
}

func LoginController(loginService ILoginService, usersService services.IUserService) *loginController {
	return &loginController{
		loginService: loginService,
		usersService: usersService,
	}
}

func (ctrl *loginController) Login(c *gin.Context) {
	// Not necessary but as an engineer you have to be curious on something.
	startTime := time.Now()

	// Works as a DTO (Data Transfer Object)
	var creds LoginInput

	// Validate the input according to DTO, return an error.
	err := c.ShouldBindJSON(&creds)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Call service to login, passing user's email and password
	// service will return bool if valid or invalid.
	id, isValid := ctrl.loginService.Login(creds.Email, creds.Password)
	if !isValid {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "wrong email or password",
		})
		return
	}

	// At this point, user credential is verified,
	// so generate access token and return back as JSON.
	accessToken := GenerateToken(id)

	c.JSON(http.StatusOK, gin.H{
		"status":       http.StatusOK,
		"access_token": accessToken,
		"timestamp":    time.Now(),
		"response_ms":  time.Now().UnixMilli() - startTime.UnixMilli(),
	})
}

func (ctrl *loginController) Register(c *gin.Context) {
	// Not necessary but as an engineer you have to be curious on something.
	startTime := time.Now()

	// Register DTO
	var userInput RegisterInput

	// Validate the input according to DTO, return an error.
	bodyErr := c.ShouldBindJSON(&userInput)
	if bodyErr != nil {
		errorMessage := []string{}
		for _, e := range bodyErr.(validator.ValidationErrors) {
			errorMessage = append(errorMessage, fmt.Sprintf("Error on field %s, reason %s", e.Field(), e.ActualTag()))
		}
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": errorMessage})
		return
	}

	// Call service to handle user register, if fails, return error
	resultErr := ctrl.loginService.HandleRegister(userInput)
	if resultErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": resultErr.Error()})
		return
	}

	// At this point, user already created in the database,
	// query that data to be used in the response.
	created, findErr := ctrl.usersService.FindByEmail(userInput.Email)
	if findErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": findErr.Error()})
		return
	}

	// Generate access token
	accessToken := GenerateToken(created.ID)

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": gin.H{
			"username":     created.Username,
			"email":        created.Email,
			"access_token": accessToken,
		},
		"timestamp":   time.Now(),
		"response_ms": time.Now().UnixMilli() - startTime.UnixMilli(),
	})
}
