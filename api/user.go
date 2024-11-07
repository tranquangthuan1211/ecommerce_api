package api

import (
	"ecommerce/database"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func loginHandler(router *gin.RouterGroup, loginHandler gin.HandlerFunc) {
	router.POST("/users/login", loginHandler)
}

func registerHandler(router *gin.RouterGroup, db *gorm.DB) {
	router.POST("/users/register", func(c *gin.Context) {
		req := &database.Register{}

		// Bind the JSON payload to the request struct
		if err := c.ShouldBindJSON(req); err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid input: " + err.Error(),
			})
			return
		}
		req.ID = uuid.New().String()
		result := db.Table("users").Create(req)
		if result.Error != nil {
			c.JSON(500, gin.H{
				"error": "Failed to create user",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "User registered successfully",
		})
	})
}
func getUsers(router *gin.RouterGroup, db *gorm.DB) {
	router.GET("/", func(c *gin.Context) {
		var users []database.UserResponse
		db.Raw(`select * from users where deleted_at is null`).Scan(&users)
		c.JSON(200, users)
	})
}

func getUserByID(router *gin.RouterGroup, db *gorm.DB) {
	router.GET("/:id", func(c *gin.Context) {
		var user database.UserResponse
		id := c.Param("id")
		db.Raw(`select * from users where id=?`, id).Scan(&user)
		c.JSON(200, user)
	})
}

// func updateUser(router *gin.RouterGroup, db *gorm.DB) {
// 	router.PUT("/:id", func(c *gin.Context) {
// 		req := &database.UserUpdate{}
// 		id := c.Param("id")

// 		// Bind the JSON payload to the request struct
// 		if err := c.ShouldBindJSON(req); err != nil {
// 			c.JSON(400, gin.H{
// 				"error": "Invalid input: " + err.Error(),
// 			})
// 			return
// 		}

// 		result := db.Table("USERS").Where("id = ?", id).Updates(req)
// 		if result.Error != nil {
// 			c.JSON(500, gin.H{
// 				"error": "Failed to update user",
// 			})
// 			return
// 		}

//			// Return a success response
//			c.JSON(200, gin.H{
//				"message": "User updated successfully",
//			})
//		})
//	}
func deleteUser(router *gin.RouterGroup, db *gorm.DB) {
	router.DELETE("/:id", func(c *gin.Context) {
		id := c.Param("id")
		result := db.Table("USERS").Where("id = ?", id).Delete(&database.UserResponse{})
		if result.Error != nil {
			c.JSON(500, gin.H{
				"error": "Failed to delete user",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "User deleted successfully",
		})
	})
}
