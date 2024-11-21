package api

import (
	"ecommerce/database"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func getGifts(router *gin.RouterGroup, db *gorm.DB) {
	router.GET("/", func(c *gin.Context) {
		var gifts []database.GiftDetail
		result := db.Table("gifts").
			Select("ranks.name AS name, ranks.coupon AS sale,gifts.created_at,gifts.updated_at").
			Joins("join ranks on gifts.id_rank = ranks.id").
			Scan(&gifts)
		if result.Error != nil {
			c.JSON(500, gin.H{
				"error": "Failed to fetch gifts",
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "Gifts fetched successfully",
			"data":    gifts,
		})
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
func updateGift(router *gin.RouterGroup, db *gorm.DB) {
	router.PATCH("/:id", func(c *gin.Context) {
		req := &database.GiftUpdate{}
		if err := c.ShouldBindJSON(req); err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid input: " + err.Error(),
			})
			return
		}
		id := c.Param("id")
		result := db.Table("gifts").Where("id = ?", id).Updates(req)
		if result.Error != nil {
			c.JSON(500, gin.H{
				"error": "Failed to update gift",
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "Gift updated successfully",
		})
	})
}
func getGiftForUser(router *gin.RouterGroup, db *gorm.DB) {
	router.GET("/:id", func(c *gin.Context) {
		var gifts []database.GiftForUser
		id := c.Param("id")
		fmt.Println(id)
		result := db.Table("gifts").
			Select("ranks.name AS name, ranks.coupon AS sale, gifts.created_at, gifts.updated_at").
			Joins("LEFT JOIN users ON users.id = gifts.id_user").
			Joins("LEFT JOIN ranks ON users.rank_id = ranks.id").
			Where("gifts.id_user = ?", id).
			Scan(&gifts)
		if result.Error != nil {
			c.JSON(500, gin.H{
				"error": "Failed to fetch gifts",
			})

			return
		}
		c.JSON(200, gin.H{
			"message": "Gifts fetched successfully",
			"data":    gifts,
		})
	})
}
func deleteGift(router *gin.RouterGroup, db *gorm.DB) {
	router.DELETE("/:id", func(c *gin.Context) {
		id := c.Param("id")
		result := db.Table("gifts").Where("id = ?", id).Delete(&database.GiftResponse{})
		if result.Error != nil {
			c.JSON(500, gin.H{
				"error": "Failed to delete gift",
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "Gift deleted successfully",
		})
	})
}
