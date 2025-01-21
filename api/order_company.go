package api

import (
	"ecommerce/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func getOrderCompany(router *gin.RouterGroup, db *gorm.DB) {
	router.GET("/", func(c *gin.Context) {
		var orderCompanies []database.OrderCompanyResponse
		result := db.Table("order_companies").Find(&orderCompanies)
		if result.Error != nil {
			c.JSON(500, gin.H{
				"error": "Failed to get order companies",
			})
			return
		}
		c.JSON(200, gin.H{
			"data": orderCompanies,
		})
	})
}
