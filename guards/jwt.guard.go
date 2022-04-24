package guards

import (
	"fmt"
	"http-api/services"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWTGuard() gin.HandlerFunc {
	// This is basically a middleware or I like
	// to call it Guard (based on NestJS) as it is intercept
	// incoming request before sending to the controller.
	return func(c *gin.Context) {
		BEARER_SCHEMA := "Bearer "
		// Get Bearer token from Authorization header.
		authorization := c.GetHeader("Authorization")

		// If token is not exist.
		if authorization == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Slice the token and get rid of unnecessary part.
		// Before slice => "Bearer <actual token>"
		// After slice => "<actual token>"
		t := authorization[len(BEARER_SCHEMA):]

		// Call Jwt service to veriy if the token valid or not.
		token, err := services.Verify(string(t))
		if token.Valid {
			// If it is valid extract the payload.
			claims := token.Claims.(jwt.MapClaims)

			// Set the context to include authenticated user id.
			c.Set("user_id", claims["id"])
		} else {
			// If not valid, abort with status code 401
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
