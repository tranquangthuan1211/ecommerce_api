package api

import (
	"ecommerce/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func getFeedBack(router *gin.RouterGroup, db *gorm.DB) {
	router.GET("/", func(c *gin.Context) {
		var feedbacks []database.FeedBackResponse
		result := db.Table("feedbacks").
			Joins("JOIN users ON users.id = feedbacks.id_user").
			Find(&feedbacks)
		if result.Error != nil {
			c.JSON(500, gin.H{
				"error": "Failed to get feedbacks",
			})
			return
		}
		c.JSON(200, gin.H{
			"data": feedbacks,
		})
	})
}
func getFeedBackByID(router *gin.RouterGroup, db *gorm.DB) {
	router.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		var feedback database.FeedBackResponse
		result := db.Table("feedbacks").
			Joins("JOIN users ON users.id = feedbacks.id_user").
			Where("id = ?", id).First(&feedback)
		if result.Error != nil {
			c.JSON(500, gin.H{
				"error": "Failed to get feedback",
			})
			return
		}
		c.JSON(200, gin.H{
			"data": feedback,
		})
	})
}
func createFeedBack(router *gin.RouterGroup, db *gorm.DB) {
	router.POST("/", func(c *gin.Context) {
		req := &database.FeedBackResponse{}
		if err := c.ShouldBindJSON(req); err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid input: " + err.Error(),
			})
			return
		}

		err := CheckInputError(req)
		if err != nil {
			returnError(c, 400, err.Error())
			return
		}

		result := db.Table("feedbacks").Create(req)
		if result.Error != nil {
			c.JSON(500, gin.H{
				"error": "Failed to create feedback",
			})
			return
		}
		c.JSON(200, gin.H{
			"data": req,
		})
	})
}
func updateFeedBack(router *gin.RouterGroup, db *gorm.DB) {
	router.PATCH("/:id", func(c *gin.Context) {
		id := c.Param("id")
		req := &database.FeedBackResponse{}
		if err := c.ShouldBindJSON(req); err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid input: " + err.Error(),
			})
			return
		}

		err := CheckInputError(req)
		if err != nil {
			returnError(c, 400, err.Error())
			return
		}

		result := db.Table("feedbacks").Where("id = ?", id).Updates(req)
		if result.Error != nil {
			c.JSON(500, gin.H{
				"error": "Failed to update feedback",
			})
			return
		}
		c.JSON(200, gin.H{
			"data": req,
		})
	})
}
func deleteFeedBack(router *gin.RouterGroup, db *gorm.DB) {
	router.Group("/:id", func(c *gin.Context) {
		id := c.Param("id")
		result := db.Table("feedbacks").Where("id = ?", id).Delete(&database.FeedBackResponse{})
		if result.Error != nil {
			c.JSON(500, gin.H{
				"error": "Failed to delete feedback",
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "Feedback deleted successfully",
		})
	})
}
