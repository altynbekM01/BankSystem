package Controller

import (
	_ "GolangwithFrame/src/app/service"
	"GolangwithFrame/src/domain/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CartController interface {
	CreateCart(ctx *gin.Context)
	DeleteCategory(ctx *gin.Context)
	FindAllCarts(ctx *gin.Context)
	GetUserCart(ctx *gin.Context)
}

func (c *Controller) FindAllCarts(ctx *gin.Context) {
	ctx.JSON(200, c.service.FindAllCarts())
	//fmt.Printf("ClientIP: %s\n", ctx.ClientIP())
}

func (c *Controller) CreateCart(ctx *gin.Context) {
	var cart model.Cart
	err := ctx.ShouldBindJSON(&cart)
	strlogin, _ := ctx.Get("userlogin")
	strlogin1 := strlogin.(uint)
	cart.UserLogin = strlogin1
	//fmt.Println(cart)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.service.CreateCart(cart)
	ctx.JSON(http.StatusOK, gin.H{"message": "Product was added to your cart!"})

}

func (c *Controller) DeleteCart(ctx *gin.Context) {
	var cart model.Cart
	login := ctx.Param("user_login")
	loginINT, _ := strconv.Atoi(login)
	cart.UserLogin = uint(loginINT)
	err := c.service.DeleteCart(uint(loginINT))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Your cart is empty!"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Cart was deleted!"})

}

func (c *Controller) GetUserCart(ctx *gin.Context) {
	login := ctx.Param("user_login")
	loginINT, _ := strconv.Atoi(login)
	carts, err := c.service.GetUserCart(uint(loginINT))
	if err != nil {
		ctx.JSON(404, gin.H{"message": "Your cart is empty. Add some product!"})
		return
	}
	ctx.JSON(200, carts)
}
