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
	db       *gorm.DB
	host     = os.Getenv("DB_HOST")
	username = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	dbName   = os.Getenv("DB_NAME")
	port     = os.Getenv("DB_PORT")
	err      error
)

func StartDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", host, username, password, dbName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Product{})

}

func main() {
	StartDB()

	if err != nil {
		panic(err)
	}

	var port = os.Getenv("PORT")

	r := gin.Default()

	api := r.Group("/api")

	productController := controller.ProductController{
		DB: db,
	}

	productController.Routes(api)

	r.Run(":" + port)

}
