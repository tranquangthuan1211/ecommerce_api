package api

import (
	"ecommerce/utils"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	cors "github.com/itsjamie/gin-cors"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

type Routes struct {
	Path    string
	Handler []func(*gin.RouterGroup, *gorm.DB)
}

var router = []Routes{
	{"/users", []func(*gin.RouterGroup, *gorm.DB){getUsers}},
}

func RunServer(db *gorm.DB) {
	r := gin.Default()
	corsConfig := cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE, OPTIONS, PATCH",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          15 * time.Second,
		ValidateHeaders: false,
	}
	r.Use(cors.Middleware(corsConfig))

	authMiddleware, err := getAuthMiddleware(db)
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	errInit := authMiddleware.MiddlewareInit()
	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}
	v1 := r.Group("/api/v1")
	registerHandler(v1, db)
	loginHandler(v1, authMiddleware.LoginHandler)
	v1.Use(authMiddleware.MiddlewareFunc())
	for _, route := range router {
		for _, controller := range route.Handler {
			controller(v1.Group(route.Path), db)
		}
	}

	// use ginSwagger middleware to serve the API docs
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":" + utils.PORT)
}

var validate *validator.Validate

func CheckInputError(input interface{}) error {
	validate = validator.New()
	err := validate.Struct(input)
	if err == nil {
		return nil
	}

	if _, ok := err.(*validator.InvalidValidationError); ok {
		return err
	}

	fields := []string{}
	for _, err := range err.(validator.ValidationErrors) {
		fields = append(fields, err.Namespace()+" "+err.Type().Name())
	}
	return fmt.Errorf("Error:Field validation for: %v", strings.Join(fields, ", "))
}
