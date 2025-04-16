package main

import (
	"game-store/config"
	"game-store/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.ConnectDatabase()

	routes.RegisterGameRoutes(r)

	r.Run(":8080")
}
