package Controller

import (
	_ "GolangwithFrame/src/app/service"
	"GolangwithFrame/src/domain/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ProductController interface {
	FindAllProducts(ctx *gin.Context)
	CreateProduct(ctx *gin.Context)
	UpdateProduct(ctx *gin.Context)
	DeleteProduct(ctx *gin.Context)
	GetProduct(ctx *gin.Context)
}

func (c *Controller) FindAllProducts(ctx *gin.Context) {
	ctx.JSON(200, c.service.FindAllProducts())
	//fmt.Printf("ClientIP: %s\n", ctx.ClientIP())
}

func (c *Controller) CreateProduct(ctx *gin.Context) {
	var product model.Product
	err := ctx.ShouldBindJSON(&product)
	fmt.Println(product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.service.CreateProduct(product)
	ctx.JSON(http.StatusOK, gin.H{"message": "Product was created"})

}

func (c *Controller) UpdateProduct(ctx *gin.Context) {
	var product model.Product
	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Used invalid ID"})
		return
	}
	product.Id = id
	err = c.service.UpdateProduct(product)
	//fmt.Println(err)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "There is nothing to update"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Product was updated"})

}

func (c *Controller) DeleteProduct(ctx *gin.Context) {
	var product model.Product
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Used invalid ID"})
		return
	}
	product.Id = id
	err = c.service.DeleteProduct(product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "There is nothing to delete"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Product was deleted!"})

}

func (c *Controller) GetProduct(ctx *gin.Context) {
	var product *model.Product
	ProductId := ctx.Param("id")
	ProductIdInt, err := strconv.Atoi(ProductId)
	if err != nil {
		ctx.JSON(404, gin.H{"message": "Used invalid ID"})
		return
	}
	product = c.cache.Get(ProductId)
	//fmt.Println(product)
	if product == nil {
		fmt.Println("condition 1 (NOT IN CACHE)")
		ctx.ShouldBindJSON(&product)
		prod, err := c.service.GetProduct(ProductIdInt)
		if err != nil {
			ctx.JSON(404, gin.H{"message": "There is no object with this ID"})
			return
		}
		c.cache.Set(ProductId, &prod)

		ctx.JSON(200, gin.H{"message": prod})
		//fmt.Println(err)
	} else {
		fmt.Println("condition 2 (FOUND IN CACHE)")
		ctx.JSON(200, gin.H{"message": product})
	}
}
