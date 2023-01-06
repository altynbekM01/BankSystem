package Controller

import (
	"GolangwithFrame/src/domain/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CreditController interface {
	FindAllCredits(ctx *gin.Context)
	CreateCredit(ctx *gin.Context)
	FindCreditsById(ctx *gin.Context)
}

func (c *Controller) FindAllCredits(ctx *gin.Context) {
	ctx.JSON(200, c.service.FindAllCredits())
	//fmt.Printf("ClientIP: %s\n", ctx.ClientIP())
}

func (c *Controller) CreateCredit(ctx *gin.Context) {
	var category model.Credit
	err := ctx.ShouldBindJSON(&category)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.service.CreateCredit(category)
	ctx.JSON(http.StatusOK, gin.H{"message": "Category was created"})

}

func (c *Controller) FindCreditById(ctx *gin.Context) {
	user_id := ctx.Param("id")
	CreditIdInt, err := strconv.Atoi(user_id)
	if err != nil {
		ctx.JSON(404, gin.H{"message": "Used invalid ID"})
		return
	}
	credits, err := c.service.FindCreditsById(uint(CreditIdInt))
	if err != nil {
		ctx.JSON(404, gin.H{"message": "There is no object with this ID"})
		return
	}
	ctx.JSON(200, credits)

}
