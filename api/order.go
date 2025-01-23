package api

import (
	"ecommerce/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func getOrder(router *gin.RouterGroup, db *gorm.DB) {
	router.GET("/", func(c *gin.Context) {
		var orders []database.OrderResponse
		result := db.Table("orders").Find(&orders)
		if result.Error != nil {
			c.JSON(500, gin.H{
				"error": "Failed to get orders",
			})
			return
		}
		c.JSON(200, gin.H{
			"data": orders,
		})
	})
}
func getOrderByID(router *gin.RouterGroup, db *gorm.DB) {
	router.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		var order database.OrderResponse
		result := db.Table("orders").Where("id = ?", id).First(&order)
		if result.Error != nil {
			c.JSON(500, gin.H{
				"error": "Failed to get order",
			})
			return
		}
		c.JSON(200, gin.H{
			"data": order,
		})
	})
}

func createOrder(router *gin.RouterGroup, db *gorm.DB) {
	router.POST("/", func(c *gin.Context) {
		req := &database.OrderResponse{}
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

		result := db.Table("orders").Create(req)
		if result.Error != nil {
			c.JSON(500, gin.H{
				"error": "Failed to create order",
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "Order created successfully",
			"data":    req,
		})
	})
}
