package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_social/internal/common"
	"go_social/internal/routers"
	"os"

	"go_social/database"
	"go_social/database/entities"
)

func main() {
	common.EnvLoad()
	db := database.DB()
	db.AutoMigrate(&entities.UserEntity{})

	fmt.Println("Hello world")
	server := gin.Default()
	server.LoadHTMLGlob("internal/views/*.html")
	routers.HomeRoutes(server)
	routers.UserRouter(server)
	server.Run(":" + os.Getenv("SERVER_PORT"))
}
