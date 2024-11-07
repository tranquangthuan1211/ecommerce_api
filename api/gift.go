package api

import (
	"ecommerce/database"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func getGifts(router *gin.RouterGroup, db *gorm.DB) {
	router.GET("/", func(c *gin.Context) {
		var gifts []database.GiftResponse
		db.Raw(`select * from gifts where deleted_at is null`).Scan(&gifts)
		c.JSON(200, gifts)
	})
}
func createGift(router *gin.RouterGroup, db *gorm.DB) {
	router.POST("/", func(c *gin.Context) {
		req := &database.GiftResponse{}
		if err := c.ShouldBindJSON(req); err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid input: " + err.Error(),
			})
			return
		}
		req.ID = uuid.New().String()
		result := db.Table("gifts").Create(req)
		if result.Error != nil {
			c.JSON(500, gin.H{
				"error": "Failed to create gift",
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "Gift created successfully",
		})
	})
}
