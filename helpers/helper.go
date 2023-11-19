package helpers

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GetClaims(c *gin.Context) jwt.Claims {

	cookie, _ := c.Cookie("token")

	token, _ := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	claims := token.Claims
	return claims
}
