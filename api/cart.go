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
func updateCart(router *gin.RouterGroup, db *gorm.DB) {
	router.PATCH("/:id", func(c *gin.Context) {
		id := c.Param("id")
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
		result := db.Table("carts").Where("id = ?", id).Updates(req)
		if result.Error != nil {
			c.JSON(500, gin.H{
				"error": "Failed to update cart",
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "Cart updated successfully",
			"data":    req,
		})
	})
}
func deleteCart(router *gin.RouterGroup, db *gorm.DB) {
	router.DELETE("/:id", func(c *gin.Context) {
		id := c.Param("id")
		result := db.Table("carts").Where("id = ?", id).Delete(&database.CartDataBase{})
		if result.Error != nil {
			c.JSON(500, gin.H{
				"error": "Failed to delete cart",
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "Cart deleted successfully",
		})
	})
}
