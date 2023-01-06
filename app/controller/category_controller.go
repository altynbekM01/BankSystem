package Controller

import (
	_ "GolangwithFrame/src/app/service"
	"GolangwithFrame/src/domain/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CategoryController interface {
	FindAllCategory(ctx *gin.Context)
	CreateCategory(ctx *gin.Context)
	UpdateCategory(ctx *gin.Context)
	DeleteCategory(ctx *gin.Context)
	GetCategory(ctx *gin.Context)
	FindProductByCategory(ctx *gin.Context)
}

func (c *Controller) FindAllCategory(ctx *gin.Context) {
	ctx.JSON(200, c.service.FindAllCategory())
	//fmt.Printf("ClientIP: %s\n", ctx.ClientIP())
}

func (c *Controller) CreateCategory(ctx *gin.Context) {
	var category model.Category
	err := ctx.ShouldBindJSON(&category)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.service.CreateCategory(category)
	ctx.JSON(http.StatusOK, gin.H{"message": "Category was created"})

}

func (c *Controller) UpdateCategory(ctx *gin.Context) {
	var category model.Category
	err := ctx.ShouldBindJSON(&category)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Used invalid ID"})
		return
	}
	category.Id = id
	err = c.service.UpdateCategory(category)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "There is nothing to update"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Category was updated"})

}

func (c *Controller) DeleteCategory(ctx *gin.Context) {
	var category model.Category
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Used invalid ID"})
		return
	}
	category.Id = id
	err = c.service.DeleteCategory(category)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "There is nothing to delete"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Category was deleted!"})

}

func (c *Controller) GetCategory(ctx *gin.Context) {
	var category model.Category
	CategoryId := ctx.Param("id")
	CategoryIdInt, err := strconv.Atoi(CategoryId)
	if err != nil {
		ctx.JSON(404, gin.H{"message": "Used invalid ID"})
		return
	}
	ctx.ShouldBindJSON(&category)
	category, err = c.service.GetCategory(CategoryIdInt)
	if err != nil {
		ctx.JSON(404, gin.H{"message": "There is no object with this ID"})
		return
	}
	ctx.JSON(200, gin.H{"message": category})
}
func (c *Controller) FindProductsByCategory(ctx *gin.Context) {
	category_id := ctx.Param("id")
	CategoryIdInt, err := strconv.Atoi(category_id)
	if err != nil {
		ctx.JSON(404, gin.H{"message": "Used invalid ID"})
		return
	}
	products, err := c.service.FindProductsByCategory(CategoryIdInt)
	if err != nil {
		ctx.JSON(404, gin.H{"message": "There is no object with this ID"})
		return
	}
	ctx.JSON(200, products)

}
