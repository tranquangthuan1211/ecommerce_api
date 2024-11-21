package api

import (
	"ecommerce/database"
	// ginJWT "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func getRanks(router *gin.RouterGroup, db *gorm.DB) {
	router.GET("/", func(c *gin.Context) {
		var ranks []database.RankResponse
		db.Raw(`select * from ranks where deleted_at is null`).Scan(&ranks)
		c.JSON(200, ranks)
	})
}
func createRank(router *gin.RouterGroup, db *gorm.DB) {
	router.POST("/", func(c *gin.Context) {
		req := &database.RankResponse{}
		if err := c.ShouldBindJSON(req); err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid input: " + err.Error(),
			})
			return
		}
		req.ID = uuid.New().String()
		result := db.Table("ranks").Create(req)
		if result.Error != nil {
			c.JSON(500, gin.H{
				"error": "Failed to create rank",
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "Rank created successfully",
		})
	})
}
func updateRank(router *gin.RouterGroup, db *gorm.DB) {
	router.PATCH("/:id", func(c *gin.Context) {
		id := c.Param("id")
		req := &database.RankUpdate{}
		if err := c.ShouldBindJSON(req); err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid input: " + err.Error(),
			})
			return
		}
		result := db.Table("ranks").Where("id = ?", id).Updates(req)
		if result.Error != nil {
			c.JSON(500, gin.H{
				"error": "Failed to update rank",
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "Rank updated successfully",
		})
	})
}
func deleteRank(router *gin.RouterGroup, db *gorm.DB) {
	router.DELETE("/:id", func(c *gin.Context) {
		id := c.Param("id")
		result := db.Table("ranks").Where("id = ?", id).Delete(&database.RankResponse{})
		if result.Error != nil {
			c.JSON(500, gin.H{
				"error": "Failed to delete rank",
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "Rank deleted successfully",
		})
	})
}
