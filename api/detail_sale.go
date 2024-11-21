package api

import (
	"ecommerce/database"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func createDetailSale(router *gin.RouterGroup, db *gorm.DB) {
	router.POST("/", func(c *gin.Context) {
		req := &database.DetailSaleResponse{}
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
		req.ID = uuid.New().String()
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
