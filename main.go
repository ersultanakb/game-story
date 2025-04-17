package main

import (
	"game-store/config"
	"game-store/controllers"
	"game-store/middleware"
	"game-store/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Подключение к базе данных
	config.ConnectDatabase()

	// Публичные маршруты
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// Защищённые маршруты
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware()) // Использование middleware для защиты маршрутов
	{
		routes.RegisterGameRoutes(api) // Передаем защищенную группу маршрутов api
	}

	// Запуск сервера
	r.Run(":8081")
}
