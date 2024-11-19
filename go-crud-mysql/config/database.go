package config

import (
	"log"

	"go-crud-mysql/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	dsn := "root:9340@tcp(127.0.0.1:3308)/cruddb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

if db == nil {
    log.Fatal("Database connection is nil")
}

if err := db.AutoMigrate(&model.User{}, &model.Product{}); err != nil {
    log.Fatalf("Failed to migrate database: %v", err)
}

	return db
}
