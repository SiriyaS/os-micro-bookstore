package server

import (
	"os-micro-bookstore/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	router.GET("/ping", controller.Ping)

	bookController := controller.BookController{}
	router.GET("/books", bookController.GetAllBooks)
	router.GET("/book", bookController.GetBookByID) // pass product_id through query string
	router.POST("/book/add", bookController.InsertBook)
	router.PUT("/book/update", bookController.UpdateBookByID)    // pass product_id through query string
	router.DELETE("/book/delete", bookController.DeleteBookByID) // pass product_id through query string

	userController := controller.UserController{}
	router.POST("/register", userController.CreateUser)

	return router
}