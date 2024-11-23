package api

import (
	"ecommerce/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func getCart(router *gin.RouterGroup, db *gorm.DB) {
	router.GET("/", func(c *gin.Context) {
		var carts []database.CartDataBase
		if err := db.Table("carts").Find(&carts).Error; err != nil {
			c.JSON(500, gin.H{
				"error": "Failed to fetch carts",
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "Carts fetched successfully",
			"data":    carts,
		})
	})
}
func createCart(router *gin.RouterGroup, db *gorm.DB) {
	router.POST("/", func(c *gin.Context) {
		req := &database.CartDataBase{}
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
		result := db.Table("carts").Create(req)
		if result.Error != nil {
			c.JSON(500, gin.H{
				"error": "Failed to create cart",
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "Cart created successfully",
			"data":    req,
		})
	})
}
