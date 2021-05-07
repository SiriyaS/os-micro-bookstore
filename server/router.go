package server

import (
	"os-micro-bookstore/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(cors.Default())

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
	router.POST("/register", userController.CreateUser)
	router.POST("/login", userController.Login)
	router.PUT("/editProfile", userController.EditProfileByEmail)

	return router
}
