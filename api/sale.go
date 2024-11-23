package api

import (
	"ecommerce/database"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func getSale(Router *gin.RouterGroup, db *gorm.DB) {
	Router.GET("/", func(c *gin.Context) {
		req := &database.SaleResponse{}
		result := db.Table(req.TableName()).Find(&req)
		if result.Error != nil {
			c.JSON(500, gin.H{
				"error": "Failed to fetch product",
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "Product fetched successfully",
			"data":    req,
		})
	})
}
func createSale(Router *gin.RouterGroup, db *gorm.DB) {
	Router.POST("/", func(c *gin.Context) {
		req := &database.SaleResponse{}
		if err := c.ShouldBindJSON(req); err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid input: " + err.Error(),
			})
			return
		}
		req.ID = uuid.New().String()
		err := CheckInputError(req)
		if err != nil {
			returnError(c, 400, err.Error())
			return
		}
		result := db.Table(req.TableName()).Create(req)
		if result.Error != nil {
			c.JSON(500, gin.H{
				"error": "Failed to create product",
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "Product created successfully",
			"data":    req,
		})
	})
}
func updateSale(Router *gin.RouterGroup, db *gorm.DB) {
	Router.PATCH("/:id", func(c *gin.Context) {
		id := c.Param("id")
		req := &database.SaleUpdate{}
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
		result := db.Table("sales").Where("id = ?", id).Updates(req)
		if result.Error != nil {
			c.JSON(500, gin.H{
				"error": "Failed to update product",
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "Product updated successfully",
			"data":    result.RowsAffected,
		})
	})
}
func deleteSale(Router *gin.RouterGroup, db *gorm.DB) {
	Router.DELETE("/:id", func(c *gin.Context) {
		id := c.Param("id")
		result := db.Table("sales").Where("id = ?", id).Delete(&database.SaleResponse{})
		if result.Error != nil {
			c.JSON(500, gin.H{
				"error": "Failed to delete product",
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "Product deleted successfully",
		})
	})
}
