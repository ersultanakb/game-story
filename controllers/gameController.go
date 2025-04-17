package controllers

import (
	"game-store/config"
	"game-store/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetGames(c *gin.Context) {
	var games []models.Game
	config.DB.Find(&games)
	c.JSON(http.StatusOK, games)
}

func GetGame(c *gin.Context) {
	id := c.Param("id")
	var game models.Game
	if err := config.DB.First(&game, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Игра не найдена"})
		return
	}
	c.JSON(http.StatusOK, game)
}

func CreateGame(c *gin.Context) {
	var game models.Game
	if err := c.ShouldBindJSON(&game); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Получаем user_id из контекста (из middleware)
	userIDInterface, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in token"})
		return
	}
	userID := userIDInterface.(uint)

	game.UserID = userID

	// Создаём игру в базе данных
	if result := config.DB.Create(&game); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать игру"})
		return
	}
	c.JSON(http.StatusCreated, game)
}

func UpdateGame(c *gin.Context) {
	id := c.Param("id")
	var game models.Game

	if err := config.DB.First(&game, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Игра не найдена"})
		return
	}

	// Проверка: если пользователь — не владелец и не админ, отказать
	userID := c.GetUint("user_id")
	isAdmin := c.GetBool("is_admin")
	if game.UserID != userID && !isAdmin {
		c.JSON(http.StatusForbidden, gin.H{"error": "Недостаточно прав"})
		return
	}

	if err := c.ShouldBindJSON(&game); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&game)
	c.JSON(http.StatusOK, game)
}

func DeleteGame(c *gin.Context) {
	id := c.Param("id")
	var game models.Game

	if err := config.DB.First(&game, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Игра не найдена"})
		return
	}

	userID := c.GetUint("user_id")
	isAdmin := c.GetBool("is_admin")
	if game.UserID != userID && !isAdmin {
		c.JSON(http.StatusForbidden, gin.H{"error": "Недостаточно прав"})
		return
	}

	config.DB.Delete(&game)
	c.JSON(http.StatusOK, gin.H{"message": "Игра удалена"})
}
