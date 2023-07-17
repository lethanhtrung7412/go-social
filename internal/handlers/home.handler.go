package handlers

import (
	"github.com/gin-gonic/gin"
	"go_social/internal/common"
)

func Index(c *gin.Context) {
	common.RenderTemplate(c, "index", gin.H{
		"title": "Hello",
	})
}
