package routes

import (
	"game-store/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterGameRoutes(r *gin.Engine) {
	games := r.Group("/games")
	{
		games.GET("/", controllers.GetGames)
		games.GET("/:id", controllers.GetGame)
		games.POST("/", controllers.CreateGame)
		games.PUT("/:id", controllers.UpdateGame)
		games.DELETE("/:id", controllers.DeleteGame)
	}
}
