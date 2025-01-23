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
func createOrderCompany(router *gin.RouterGroup, db *gorm.DB) {
	router.POST("/", func(c *gin.Context) {
		req := &database.OrderCompanyResponse{}
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

		var product database.OrderCompany
		check := db.Table("order_companies").Where("id = ?", req.ID_PRODUCT).First(&product)

		if check.Error != nil {
			if check.Error == gorm.ErrRecordNotFound {
				result := db.Table("order_companies").Create(req)
				if result.Error != nil {
					c.JSON(500, gin.H{
						"error": "Failed to create order company",
					})
					return
				}
				c.JSON(200, gin.H{
					"message": "Order company created successfully",
					"data":    req,
				})
			} else { // Xử lý lỗi khác
				c.JSON(500, gin.H{
					"error": "Database error: " + check.Error.Error(),
				})
			}
			return
		}

		// Nếu sản phẩm đã tồn tại, cập nhật lại quantity
		product.QUANTITY += req.QUANTITY // Giả sử `Quantity` là số lượng cần cập nhật
		update := db.Table("order_companies").Save(&product)
		if update.Error != nil {
			c.JSON(500, gin.H{
				"error": "Failed to update product quantity",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "Product quantity updated successfully",
			"data":    product,
		})
	})
}

func updateOrderCompany(router *gin.RouterGroup, db *gorm.DB) {
	router.PATCH("/:id", func(c *gin.Context) {
		id := c.Param("id")
		req := &database.OrderCompanyResponse{}
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
		result := db.Table("order_companies").Where("id = ?", id).Updates(req)
		if result.Error != nil {
			c.JSON(500, gin.H{
				"error": "Failed to update order company",
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "Order company updated successfully",
			"data":    req,
		})
	})
}
