package routes

import (
	"jwt-auth/controllers"
	"jwt-auth/middleware"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {

	router.GET("/signup", controllers.SignupGET)
	router.POST("/signup", controllers.Signup)

	router.GET("/login", controllers.LoginGET)
	router.POST("/login", controllers.Login)

	router.POST("/logout", middleware.AuthenticateUser, controllers.Logout)

	router.GET("/user", middleware.AuthenticateUser, controllers.User)
}
