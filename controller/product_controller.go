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

	err := p.DB.Find(&productList).Error

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error Get Products",
			"error":   err,
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Get Products",
		"data":    productList,
	})
}
