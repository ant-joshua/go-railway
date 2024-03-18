package main

import (
	"fmt"
	"go-railway/controller"
	"go-railway/entity"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = os.Getenv("DB_HOST")
	username = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	dbName   = os.Getenv("DB_NAME")
	port     = os.Getenv("DB_PORT")
)

func StartDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", host, username, password, dbName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&entity.Product{})

	return db
}

func main() {
	db := StartDB()

	var port = os.Getenv("PORT")

	r := gin.Default()

	api := r.Group("/api")

	productController := controller.ProductController{
		DB: db,
	}

	productController.Routes(api)

	r.Run(":" + port)

}
