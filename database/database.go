package database

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB_URL = "root:12112004@tcp(127.0.0.1:3306)/ecommerce?charset=utf8mb4&parseTime=True&loc=Local"
var DB_ECOMMERCE = "ecommerce"

var (
	db *gorm.DB
	DB *gorm.DB
)

type BaseModel struct {
	CreatedAt time.Time `json:"created_at" example:"2023-08-17T15:40:58.131023+07:00"`
	UpdatedAt time.Time `json:"updated_at" example:"2023-08-17T15:40:58.131023+07:00"`
	// DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index" swaggertype:"string" example:"2023-08-17T15:40:58.131023+07:00"`
}

func GetDB() *gorm.DB {
	if db != nil {
		return db
	}
	db, _ := gorm.Open(mysql.Open(DB_URL), &gorm.Config{})
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("Failed to connect to database!")
	}
	sqlDB.SetConnMaxIdleTime(5 * time.Second)
	sqlDB.SetMaxIdleConns(50)
	sqlDB.SetMaxOpenConns(450)
	sqlDB.SetConnMaxLifetime(10 * time.Minute)
	DB = db

	return db
}

func Migration() (*gorm.DB, error) {
	db := GetDB()
	db.Exec("CREATE SCHEMA IF NOT EXISTS ECOMMERCE;")
	tables := []interface{}{
		&UserResponse{},
		&GiftResponse{},
		&RankResponse{},
		&SaleResponse{},
		&ProductResponse{},
		&CategoryResponse{},
		&DetailSaleResponse{},
		&CartResponse{},
	}
	for _, table := range tables {
		err := db.AutoMigrate(table)
		if err != nil {
			return db, err
		}
	}
	return db, nil
}
