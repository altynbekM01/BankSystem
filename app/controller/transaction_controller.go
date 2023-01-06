package Controller

import (
	"GolangwithFrame/src/domain/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Valid(number int) bool {
	return (number%10+checksum(number/10))%10 == 0
}

func checksum(number int) int {
	var luhn int

	for i := 0; number > 0; i++ {
		cur := number % 10

		if i%2 == 0 { // even
			cur = cur * 2
			if cur > 9 {
				cur = cur%10 + cur/10
			}
		}

		luhn += cur
		number = number / 10
	}
	return luhn % 10
}

type TransactionController interface {
	Transfer(ctx *gin.Context)
}

var body struct {
	receiver int
	amount   int
}

func (c *Controller) Transfer(ctx *gin.Context) {

	var ans model.TransferModel
	err := ctx.ShouldBindJSON(&ans)
	fmt.Println(ans)
	sender, _ := ctx.Get("userlogin")
	sender1 := sender.(uint)
	//fmt.Println(cart)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	} else if !Valid(ans.Receiver) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Card is wrong. Try another one!"})
	} else {
		total := c.service.Transfer(int(sender1), ans.Receiver, ans.Amount)
		if total == 3 {
			ctx.JSON(http.StatusOK, gin.H{"message": "You don't have enough money!"})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Your successfully sent money!"})
		}
	}

}

func (c *Controller) SeeHistorySender(ctx *gin.Context) {
	sender, _ := ctx.Get("userlogin")
	sender1 := sender.(uint)

	transactions, err := c.service.SeeHistorySender(sender1)
	if err != nil {
		ctx.JSON(404, gin.H{"message": "History is empty"})
		return
	}
	ctx.JSON(200, transactions)

}

func (c *Controller) SeeHistoryRececiver(ctx *gin.Context) {
	sender, _ := ctx.Get("userlogin")
	sender1 := sender.(uint)

	transactions, err := c.service.SeeHistoryReceiver(sender1)
	if err != nil {
		ctx.JSON(404, gin.H{"message": "History is empty"})
		return
	}
	ctx.JSON(200, transactions)

}
