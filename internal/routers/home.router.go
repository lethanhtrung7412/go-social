package routers

import (
	"github.com/gin-gonic/gin"
	h "go_social/internal/handlers"
)

func HomeRoutes(route *gin.Engine) {
	auth := route.Group("/")
	{
		auth.GET("", h.Index)
	}
}
