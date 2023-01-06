package main

import (
	Controller "GolangwithFrame/app/controller"
	middleware2 "GolangwithFrame/app/middleware"
	"GolangwithFrame/src/app/service"
	"GolangwithFrame/src/infrastructure/cache"
	"GolangwithFrame/src/infrastructure/repository"
	"net/http"
	"os"
	"text/template"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var (
//Host     = os.Getenv("LOCAL_DB_HOST")
//Port     = os.Getenv("LOCAL_DB_PORT")
//User     = os.Getenv("LOCAL_DB_USER")
//Password = os.Getenv("LOCAL_DB_PASSWORD")
//Dbname   = os.Getenv("LOCAL_DB_DBNAME")

//
//	Cache 			= cache.NewRedisCache("localhost:6379, 0, 15)
//	Repository	 	= repository.NewRepository()
//	Service 		= service.New(Repository)
//	controller	 	= Controller.New(*Service, Cache)

// Repository repository.ProductRepository = repository.New()
// Service    service.SER                  = service.New(Repository)
// controller controller.ProductController = controller.New(Service)
)

func init() {
	middleware2.LoadEnvVariables()

}

func payment_page(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/payments.html")
	tmpl.Execute(w, "hi")

}

func login_page(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/main.html")
	tmpl.Execute(w, "hi")

}

func main() {
	// Loading variables from Environment
	//REDIS_HOST := os.Getenv("REDIS_HOST")
	//REDIS_PORT := os.Getenv("REDIS_PORT")

	Cache := cache.NewRedisCache("localhost:6379", 0, 60)
	Repository := repository.NewRepository()
	Service := service.New(Repository)
	controller := Controller.New(*Service, Cache)

	defer Repository.CloseDB()
	//server := gin.New()
	server := gin.Default()

	//handler := http.FileServer(http.Dir("templates"))

	server.Use(
	//gin.Recovery(),
	//middleware2.Logger(),
	//middleware2.BasicAuth(),
	)

	//http.HandleFunc("/", login_page)
	//http.HandleFunc("/payments", payment_page)

	products := server.Group("/products", controller.RequireAuth)
	//products.Use(middleware2.RequireAuth)
	{
		products.GET("", controller.FindAllProducts)
		products.POST("", controller.CreateProduct)
		products.GET("/:id", controller.GetProduct)
		products.PUT("/:id", controller.UpdateProduct)
		products.DELETE("/:id", controller.DeleteProduct)
	}

	category := server.Group("/category", controller.RequireAuth)
	//category.Use(middleware2.RequireAuth)
	{

		category.GET("", controller.FindAllCategory)
		category.GET("/:id/products", controller.FindProductsByCategory)
		category.POST("", controller.CreateCategory)
		category.GET("/:id", controller.GetCategory)
		category.PUT("/:id", controller.UpdateCategory)
		category.DELETE("/:id", controller.DeleteCategory)
	}

	cart := server.Group("/cart", controller.RequireAuth)
	{

		cart.GET("", controller.FindAllCarts)
		cart.GET("/:user_login", controller.GetUserCart)
		cart.POST("", controller.CreateCart)
		cart.DELETE("/:user_login", controller.DeleteCart)
	}

	users := server.Group("/account")
	{
		users.POST("/signup", controller.SignUp)
		users.POST("/login", controller.Login)
		users.GET("/validate", controller.RequireAuth, controller.Validate)

	}

	transaction := server.Group("/transaction", controller.RequireAuth)
	{

		transaction.POST("/transfer", controller.Transfer)
		transaction.GET("/historySend", controller.SeeHistorySender)
		transaction.GET("/historyReceive", controller.SeeHistoryRececiver)

	}

	credit := server.Group("/credit", controller.RequireAuth)
	//credit.Use(middleware2.RequireAuth)
	{

		credit.GET("", controller.FindAllCredits)
		credit.GET("/:id", controller.FindCreditById)
		credit.POST("", controller.CreateCredit)
	}

	tmp, err := template.New("controller.GetUser1").ParseFiles("./templates/test.html")
	if err != nil {
		panic(err)
	}
	tmp.Execute(os.Stdout, controller.GetUser1)

	server.Static("/static", "./templates")

	server.Run(":8080")

	//cart := server.Group("/cart")
	//{
	//	cart.GET("/:id", controller.SelectFromCartByID)
	//	cart.DELETE("/:id", controller.DeleteFromCartByID)
	//	cart.POST("", controller.AddToCart)
	//	cart.PUT("", controller.UpdateCartControl)
	//}

	//server.GET("/products", func(ctx *gin.Context) {
	//	ctx.JSON(200, controller.FindAll)
	//})
	//server.POST("/products", func(ctx *gin.Context) {
	//	err := controller.Save(ctx)
	//	if err != nil {
	//		ctx.JSON(http.StatusBadRequest, gin.H{"error": erгr.Error()})
	//	} else {
	//		ctx.JSON(http.StatusOK, gin.H{"message": "Product was created"})
	//
	//	}
	//})
	//server.PUT("/products/:id", func(ctx *gin.Context) {
	//	err := controller.Update(ctx)
	//	//fmt.Println(err)
	//	if err != nil {
	//		//ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//		if err.Error() == "record not found" {
	//			ctx.JSON(http.StatusBadRequest, gin.H{"error": "There is nothing to update"})
	//		} else {
	//			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Used invalid ID"})
	//		}
	//	} else {
	//		ctx.JSON(http.StatusOK, gin.H{"message": "Product was updated"})
	//
	//	}
	//})
	//
	//server.DELETE("/products/:id", func(ctx *gin.Context) {
	//	err := controller.Delete(ctx)
	//	//fmt.Println(err)
	//	if err != nil {
	//		if err.Error() == "record not found" {
	//			ctx.JSON(http.StatusBadRequest, gin.H{"error": "There is nothing to delete"})
	//
	//		} else {
	//			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Used invalid ID"})
	//		}
	//		//fmt.Println(err.Error()[0:17])
	//	} else {
	//		ctx.JSON(http.StatusOK, gin.H{"message": "Product was deleted!"})
	//
	//	}
	//})
	//server.GET("/products/:id", func(ctx *gin.Context) {
	//	prod, err := controller.FindById(ctx)
	//	//fmt.Println(err, err.Error())
	//	if err != nil {
	//		if err.Error() == "record not found" {
	//			ctx.JSON(404, gin.H{"message": "There is no object with this ID"})
	//		} else {
	//			ctx.JSON(404, gin.H{"message": "Used invalid ID"})
	//		}
	//	} else {
	//		ctx.JSON(200, gin.H{"message": prod})
	//	}
	//})
}
