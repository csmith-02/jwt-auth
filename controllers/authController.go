package controllers

import (
	"jwt-auth/database"
	"jwt-auth/helpers"
	"jwt-auth/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func SignupGET(c *gin.Context) {
	// check to make sure nobody is logged in
	_, err := c.Cookie("token")
	if err == nil {
		c.Redirect(http.StatusSeeOther, "/user")
	}

	c.HTML(http.StatusOK, "signup.html", nil)
}

func Signup(c *gin.Context) {

	var data struct {
		Name     string
		Email    string
		Password string
	}

	data.Name = c.PostForm("name")
	data.Email = c.PostForm("email")
	data.Password = c.PostForm("password")

	err := c.Bind(&data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), 12)
	if err != nil {
		panic(http.StatusInternalServerError)
	}

	var newUser models.User

	if database.DB.Where("email = ?", data.Email).First(&newUser); newUser.Id != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "user already exists",
		})
	}

	newUser = models.User{
		Name:     data.Name,
		Email:    data.Email,
		Password: string(hash),
	}

	database.DB.Create(&newUser)
	c.Redirect(http.StatusSeeOther, "/login")
}

func LoginGET(c *gin.Context) {
	// check to make sure nobody is logged in
	_, err := c.Cookie("token")
	if err == nil {
		c.Redirect(http.StatusSeeOther, "/user")
	}
	c.HTML(http.StatusOK, "login.html", nil)
}

func Login(c *gin.Context) {
	var data struct {
		Email    string
		Password string
	}

	data.Email = c.PostForm("email")
	data.Password = c.PostForm("password")

	err := c.Bind(&data)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not bind data",
		})
		return
	}

	var existingUser models.User

	if database.DB.Where("email = ?", data.Email).First(&existingUser); existingUser.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "user not found",
		})
		return
	}

	// Now need to check passwords to make sure they match
	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(data.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "incorrect password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": existingUser.Id,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "token could not be signed",
		})
		return
	}

	c.SetCookie("token", tokenString, int(time.Hour)*1, "/", "localhost", false, true)

	c.Redirect(http.StatusSeeOther, "/user")

}

func User(c *gin.Context) {
	var user models.User

	claims := helpers.GetClaims(c).(jwt.MapClaims)
	database.DB.Where("id = ?", claims["sub"]).First(&user)

	data := gin.H{
		"name":  user.Name,
		"email": user.Email,
		"ID":    user.Id,
	}

	c.HTML(http.StatusOK, "dashboard.html", data)
}

func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	c.Redirect(http.StatusSeeOther, "/login")
}
