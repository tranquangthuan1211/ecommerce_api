package api

import (
	"ecommerce/database"
	"ecommerce/utils"
	"errors"
	"fmt"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Quang Thuan
// Login Callback Flow
// Authenticator
// PayloadFunc
// LoginResponse

// MiddlewareFunc Callback Flow (Loggined)
// IdentityHandler
// Authorizator

// Logout Request flow (using LogoutHandler)
// LogoutResponse

// Refresh Request flow (using RefreshHandler)
// RefreshResponse

// Failures with logging in, bad tokens, or lacking privileges
// Unauthorized

// the jwt middleware
var getAuthMiddleware = func(db *gorm.DB) (*jwt.GinJWTMiddleware, error) {

	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:      "ZoneV1",
		Key:        utils.SECRET_KEY,
		Timeout:    time.Hour * 24 * 30 * 6,
		MaxRefresh: time.Hour * 24 * 30 * 6,
		// IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*database.UserResponse); ok {
				return jwt.MapClaims{
					"ID":      v.ID,
					"Name":    v.Username,
					"Email":   v.Email,
					"Address": v.Address,
					"Phone":   v.Phone,
					"Role":    v.Role,
				}
			}
			return jwt.MapClaims{}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			// util.SetDBSearchPath()
			req := database.Login{}
			c.Bind(&req)
			fmt.Println(req)
			if err := CheckInputError(req); err != nil {
				return nil, err
			}
			user := database.UserResponse{}
			err := (func() error {

				err := db.Raw(`select * from users where email=? and deleted_at is null`,
					req.Email).Scan(&user).Error
				if err != nil {
					return err
				}

				// if user.ID == "" {
				// 	teach := database.Teacher{}
				// 	err = db.Debug().Table("teacher").Where("email=?", req.Email).First(&teach).Error

				// 	if err != nil {
				// 		return err
				// 	}

				// 	if teach.ID == "" {
				// 		return errors.New("tài khoản không tồn tại")
				// 	}
				// 	user := database.Register{
				// 		ID:           teach.ID,
				// 		Last_Name:    teach.Last_Name,
				// 		First_Name:   teach.First_Name,
				// 		Hire_date:    teach.Hire_date,
				// 		Email:        teach.Email,
				// 		Major:        teach.Major,
				// 		Phone_number: teach.Phone_number,
				// 	}
				// 	result := db.Table("USERS").Create(&user)
				// 	if result.Error != nil {
				// 		return errors.New("tạo tài khoản giáo viên không thành công " + result.Error.Error())
				// 	}
				// }
				user = database.UserResponse{}
				db.Debug().Raw(`select * from users where email=? and password=? and deleted_at is null`,
					req.Email, req.Password).Scan(&user)

				if user.ID == "" {
					return errors.New("thông tin đăng nhập sai")
				}

				return nil
			})()

			if err != nil {
				return nil, err
			}
			c.Set("loginedUser", user)
			return &user, nil
		},
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			userp, _ := c.Get("loginedUser")
			user := userp.(database.UserResponse)
			c.JSON(code, database.LoginResponse{
				Code:   code,
				Token:  token,
				Expire: expire.Format(time.RFC3339),
				Data:   user,
			})
		},
		// Authorizator: func(data interface{}, c *gin.Context) bool {
		// 	// if v, ok := data.(*User); ok && v.UserName == "admin" {
		// 	// 	return true
		// 	// }

		// 	return true
		// },
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})
}