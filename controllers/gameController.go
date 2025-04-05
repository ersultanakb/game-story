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
	config.DB.Create(&game)
	c.JSON(http.StatusCreated, game)
}

func UpdateGame(c *gin.Context) {
	id := c.Param("id")
	var game models.Game
	if err := config.DB.First(&game, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Игра не найдена"})
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
	if err := config.DB.Delete(&game, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Игра не найдена"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Игра удалена"})
}
