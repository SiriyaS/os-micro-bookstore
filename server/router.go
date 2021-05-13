package server

import (
	"os-micro-bookstore/controller"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(CORSMiddleware())

	router.GET("/ping", controller.Ping)

	bookController := controller.BookController{}
	router.GET("/books", bookController.GetAllBooks)
	router.GET("/book", bookController.GetBookByID) // pass product_id through query string
	router.GET("/book/category", bookController.GetBookByCategory)
	router.GET("/book/author", bookController.GetBookByAuthor)
	router.GET("/book/publisher", bookController.GetBookByPublisher)
	router.POST("/book/add", bookController.InsertBook)
	router.PUT("/book/update", bookController.UpdateBookByID)    // pass product_id through query string
	router.DELETE("/book/delete", bookController.DeleteBookByID) // pass product_id through query string

	orderController := controller.OrderController{}
	router.POST("/makeOrder", orderController.CreateOrder)
	router.GET("/getOrder", orderController.GetOrderByOrderNo)

	userController := controller.UserController{}
	router.POST("/profile", userController.VerifyTokenGitHub)
	// router.POST("/register", userController.CreateUser)
	// router.POST("/login", userController.Login)
	// router.PUT("/editProfile", userController.EditProfileByEmail)

	return router
}
