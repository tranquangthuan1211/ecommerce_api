package api

import (
	"ecommerce/database"
	"ecommerce/utils"
	"log"

	ginJWT "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func returnError(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"message": message,
	})
}
func loginHandler(router *gin.RouterGroup, loginHandler gin.HandlerFunc) {
	router.POST("/users/login", loginHandler)
}

func registerHandler(router *gin.RouterGroup, db *gorm.DB) {
	router.POST("/users/register", func(c *gin.Context) {
		req := &database.Register{}
		if err := c.ShouldBindJSON(req); err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid input: " + err.Error(),
			})
			return
		}
		if err := CheckInputError(req); err != nil {
			returnError(c, 400, err.Error())
			return
		}
		req.ID = uuid.New().String()
		req.Password, _ = utils.HashPassword(req.Password)
		result := db.Table("users").Create(req)
		if result.Error != nil {
			c.JSON(500, gin.H{
				"error": "Failed to create user",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "User registered successfully",
			"data":    req,
		})
	})
}
func getUsers(router *gin.RouterGroup, db *gorm.DB) {
	router.GET("/", func(c *gin.Context) {
		var users []database.UserResponse
		if err := db.Table("users").
			Select("users.*, ranks.name AS rank_name").
			Joins("left join ranks on users.rank_id = ranks.id").
			Order("ranks.name ASC").
			Find(&users).Error; err != nil {
			log.Println("Error fetching users with rank:", err)
			return
		}
		c.JSON(200, gin.H{
			"message": "Users fetched successfully",
			"data":    users,
		})
	})
}

func getUserByID(router *gin.RouterGroup, db *gorm.DB) {
	router.GET("/:id", func(c *gin.Context) {
		var user database.UserResponse
		id := c.Param("id")
		if err := db.Table("users").
			Select("users.*, ranks.name AS rank_name").
			Joins("left join ranks on users.rank_id = ranks.id").
			Where("users.id = ?", id).
			First(&user).Error; err != nil {
			log.Println("Error fetching user with rank:", err)
			return
		}
		c.JSON(200, gin.H{
			"message": "Users fetched successfully",
			"data":    user,
		})
	})
}

func updateUser(router *gin.RouterGroup, db *gorm.DB) {
	router.PATCH("/:id", func(c *gin.Context) {
		pathID := c.Param("id") // ID từ đường dẫn
		tmp, exists := c.Get("JWT_PAYLOAD")
		if !exists {
			returnError(c, 401, "Unauthorized")
			return
		}
		payload, ok := tmp.(ginJWT.MapClaims)
		if !ok {
			returnError(c, 401, "Invalid token payload")
			return
		}
		ID, ok := payload["ID"].(string)
		if !ok {
			returnError(c, 401, "Invalid user ID in token")
			return
		}
		role, ok := payload["Role"].(string)
		if !ok {
			returnError(c, 401, "Invalid role in token")
			return
		}

		// Nếu không phải admin, chỉ cho phép cập nhật chính mình
		if role != "admin" {
			pathID = ID
		}
		req := map[string]interface{}{}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid input: " + err.Error(),
			})
			return
		}

		if password, ok := req["password"].(string); ok && password != "" {
			hashedPassword, _ := utils.HashPassword(password)
			req["password"] = hashedPassword
		}

		result := db.Table("users").Where("id = ?", pathID).Updates(req)
		if result.Error != nil {
			c.JSON(500, gin.H{
				"error": "Failed to update user",
			})
			return
		}

		c.JSON(200, gin.H{
			"message":        "User updated successfully",
			"updated_fields": req,
		})
	})
}

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
