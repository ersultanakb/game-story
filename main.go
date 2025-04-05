package main

import (
	"game-store/config"
	"game-store/models"
	"game-store/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.Game{})

	routes.RegisterGameRoutes(r)

	r.Run(":8080")
}
