package routers

import (
	"github.com/gin-gonic/gin"
	h "go_social/internal/handlers"
)

func UserRouter(router *gin.Engine) {
	user := router.Group("/user")
	{
		user.GET("all", h.GetAllUser)
		user.POST("signup", h.UserSignup)
		user.GET("signup", h.Signup)
		user.GET("login", h.Login)
		user.POST("login", h.UserLogin)
	}
}
