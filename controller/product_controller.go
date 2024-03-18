package controller

import (
	"go-railway/entity"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductController struct {
	DB *gorm.DB
}

func (p *ProductController) Routes(r *gin.RouterGroup) {
	r.GET("/products", p.GetProducts)
}

func (p *ProductController) GetProducts(c *gin.Context) {

	var productList []entity.Product

	p.DB.Find(&productList)

	c.JSON(200, gin.H{
		"message": "Get Products",
		"data":    productList,
	})
}
