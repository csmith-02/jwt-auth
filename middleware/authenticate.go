package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthenticateUser(c *gin.Context) {

	cookie, err := c.Cookie("token")

	if err != nil {
		fmt.Println("\nYou cannot access this page until you are LOGGED IN.")
		c.Abort()
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	// Make sure the cookie is correct
	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		fmt.Println("Invalid Token.")
		c.Abort()
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	_ = token

	// if the cookie exists and is correct then the user is authenticated
	c.Next()
}
