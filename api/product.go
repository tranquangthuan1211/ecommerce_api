package api

import (
	"ecommerce/database"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func getProducts(router *gin.RouterGroup, db *gorm.DB) {
	router.GET("/", func(c *gin.Context) {
		var products []database.DetailProduct
		if err := db.Table("products").
			Select("sales.discount as sale ,products.name,products.Price ,products.manufacturer,products.max_quantity,products.currently_quantity,categories.name AS category_name").
			Joins("JOIN categories ON products.category_id = categories.id").
			Joins("JOIN detail_sales ON products.id = detail_sales.product_id").
			Joins("JOIN sales ON detail_sales.sale_id = sales.id").
			Find(&products).Error; err != nil {
			c.JSON(500, gin.H{
				"error": "Failed to fetch products",
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "Products fetched successfully",
			"data":    products,
		})
	})
}
func getproductByID(router *gin.RouterGroup, db *gorm.DB) {
	router.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		var product database.ProductResponse
		if err := db.Table("products").
			Joins("JOIN categories ON products.category_id = categories.id").
			Joins("JOIN detail_sales ON products.id = detail_sales.product_id").
			Joins("JOIN sales ON detail_sales.sale_id = sales.id").
			Where("id = ?", id).
			First(&product).Error; err != nil {
			c.JSON(500, gin.H{
				"error": "Failed to fetch product",
				"data":  nil,
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "Product fetched successfully",
			"data":    product,
		})
	})
}
func createProduct(router *gin.RouterGroup, db *gorm.DB) {
	router.POST("/", func(c *gin.Context) {
		req := &database.ProductResponse{}

		if err := c.ShouldBindJSON(req); err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid input: " + err.Error(),
			})
			return
		}
		if err := CheckInputError(req); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		req.ID = uuid.New().String()
		result := db.Table("products").Create(req)
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
func updateProduct(router *gin.RouterGroup, db *gorm.DB) {
	router.PATCH("/:id", func(c *gin.Context) {
		id := c.Param("id")
		req := &database.ProductUpdate{}

		if err := c.ShouldBindJSON(req); err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid input: " + err.Error(),
			})
			return
		}
		if err := CheckInputError(req); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		result := db.Table("products").Where("id = ?", id).Updates(req)
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
