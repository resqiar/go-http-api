package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type IJwtService interface {
	GenerateToken(id string) string
	Verify(token string) (*jwt.Token, error)
}

type JwtClaims struct {
	Id int64 `json:"id"`
	jwt.StandardClaims
}

func GenerateToken(id int64) string {
	// JWT token will expire after certain time,
	// here we define how long it will expire and invalid.
	// In this case we define to expire after 24 hours (1 day)
	expirationTime := time.Now().Add(24 * time.Hour)

	// If you know how JWT works, you can call claims as "Payload"
	// it basically carry information data inside.
	// Learn more see https://jwt.io/
	claims := &JwtClaims{
		Id: id,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Here we sign the token with HS256 Algorithm
	// including the payload above.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Secret key from .env file
	secretKey := os.Getenv("JWT_SECRET_KEY")

	// After we sign the payload,
	// now time to sign it with our secretKey.
	accessToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		panic(err)
	}
	// return back as accessToken "string"
	return accessToken
}

func Verify(token string) (*jwt.Token, error) {
	// This function will verify whether the token is valid or not
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("invalid token")
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
}
