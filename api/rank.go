package api

import (
	"ecommerce/database"

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
