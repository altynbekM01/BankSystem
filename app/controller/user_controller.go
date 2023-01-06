package Controller

import (
	_ "GolangwithFrame/src/app/service"
	"GolangwithFrame/src/domain/model"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserController interface {
	//FindAllUsers(ctx *gin.Context)
	//CreateUser(ctx *gin.Context)
	//UpdateUser(ctx *gin.Context)
	//DeleteUser(ctx *gin.Context)
	//GetUser(ctx *gin.Context)
	SignUp(ctx *gin.Context)
	Login(ctx *gin.Context)
	Validate(ctx *gin.Context)
}

//func (c *Controller) FindAllUsers(ctx *gin.Context) {
//	ctx.JSON(200, c.service.FindAllUsers())
//}
//
//func (c *Controller) CreateUser(ctx *gin.Context) {
//	var user model.User
//	err := ctx.ShouldBindJSON(&user)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//	c.service.CreateUser(user)
//	ctx.JSON(http.StatusOK, gin.H{"message": "User was created"})
//}
//
//func (c *Controller) UpdateUser(ctx *gin.Context) {
//	var user model.User
//	err := ctx.ShouldBindJSON(&user)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	login := ctx.Param("login")
//	user.Login = login
//	err = c.service.UpdateUser(user)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": "There is no such a user!"})
//		return
//	}
//	ctx.JSON(http.StatusOK, gin.H{"message": "User was updated"})
//}
//
//func (c *Controller) DeleteUser(ctx *gin.Context) {
//	var user model.User
//	login := ctx.Param("login")
//	user.Login = login
//	err := c.service.DeleteUser(user)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": "There is nothing to delete"})
//		return
//	}
//	ctx.JSON(http.StatusOK, gin.H{"message": "User was deleted!"})
//}

func (c *Controller) GetUser(ctx *gin.Context) {
	var user model.User
	UserLogin := ctx.Param("login")
	log, _ := strconv.Atoi(UserLogin)
	ctx.ShouldBindJSON(&user)
	prod, err := c.service.GetUser(uint(log))
	if err != nil {
		ctx.JSON(404, gin.H{"message": "There is no object with this Login"})
		return
	}
	ctx.JSON(200, gin.H{"message": prod})
}

func (c *Controller) GetUser1(ctx *gin.Context) {
	var user model.User
	UserLogin := ctx.Param("login")
	log, _ := strconv.Atoi(UserLogin)
	ctx.ShouldBindJSON(&user)
	prod, err := c.service.GetUser1(uint(log))
	if err != nil {
		ctx.JSON(404, gin.H{"message": "There is no object with this Login"})
		return
	}
	ctx.JSON(200, gin.H{"message": prod})
}

func (c *Controller) SignUp(ctx *gin.Context) {
	var cur_user model.User
	err := ctx.ShouldBindJSON(&cur_user)
	//fmt.Println(body)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body!"})
		return
	}
	fmt.Println(cur_user.Login)
	fmt.Println(cur_user.Password)

	hash, err := bcrypt.GenerateFromPassword([]byte(cur_user.Password), 10)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password!"})
		return
	}
	user := model.User{Login: cur_user.Login, Password: string(hash)}
	c.service.CreateUser(user)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "user created"})

}

func (c *Controller) Login(ctx *gin.Context) {
	var cur_user model.User
	err := ctx.ShouldBindJSON(&cur_user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body!"})
		return
	}

	var user model.User
	user, err = c.service.GetUser(cur_user.Login)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "didnt find"})
		return

	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(cur_user.Password))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "compare"})
		return

	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"subject": user.Login,
		"exp":     time.Now().Add(time.Hour).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return

	}
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", tokenString, 3600, "", "", false, true)
	ctx.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

func (c *Controller) Validate(ctx *gin.Context) {
	fmt.Println("OKAY")
	user, _ := ctx.Get("user")
	userlogin, _ := ctx.Get("userlogin")
	str_login := userlogin.(uint)
	fmt.Println(str_login)
	ctx.JSON(http.StatusAccepted, gin.H{
		"message": user,
	})
}
