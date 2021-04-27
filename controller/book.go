package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"os-micro-bookstore/form"
	"os-micro-bookstore/model"
)

type BookController struct{}

func (bc BookController) GetAllBooks(c *gin.Context) {
	log.Println("[Book: GetAllBooks]")
	bookModel := model.BookModel{}

	books, err := bookModel.ReadAll()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while getting books",
		})
		return
	}

	log.Println("Get all books successfully")
	c.JSON(http.StatusOK, books)
}

func (bc BookController) GetBookByID(c *gin.Context) {
	log.Println("[Book: GetBookByID]")
	bookModel := model.BookModel{}

	bookISBN := c.Query("book_isbn")

	book, err := bookModel.ReadByID(bookISBN)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while getting books",
		})
		return
	}
	if (book == form.Book{}) {
		log.Println("No book belong to this ID.")
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No book belong to this ISBN.",
		})
		return
	}

	log.Println("Get book successfully")
	c.JSON(http.StatusOK, book)
}

func (bc BookController) InsertBook(c *gin.Context) {
	log.Println("[Book: InsertBook]")
	bookModel := model.BookModel{}

	var request form.BookRequest
	err := c.BindJSON(&request)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error while binding to model.",
		})
		return
	}

	err = bookModel.Add(request)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while adding book.",
		})
		return
	}

	log.Println("Insert book successfully")
	c.JSON(http.StatusOK, gin.H{
		"message": "Add book successfully.",
	})
}

func (bc BookController) UpdateBookByID(c *gin.Context) {
	log.Println("[Book: UpdateBookByID]")
	bookModel := model.BookModel{}

	bookISBN := c.Query("book_isbn")

	var request form.BookRequest
	err := c.BindJSON(&request)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error while binding to model.",
		})
		return
	}

	// check if there is product belong to this product_id
	book, err := bookModel.ReadByID(bookISBN)
	if (book == form.Book{}) || err != nil {
		log.Println("No book belong to this ID.")
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No book belong to this ISBN.",
		})
		return
	}

	err = bookModel.UpdateByID(bookISBN, request)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while update book.",
		})
		return
	}

	log.Println("Update book successfully")
	c.JSON(http.StatusOK, gin.H{
		"message": "Update book successfully.",
	})
}

func (bc BookController) DeleteBookByID(c *gin.Context) {
	log.Println("[Book: DelateBookByID]")
	bookModel := model.BookModel{}

	bookISBN := c.Query("book_isbn")

	// check if there is product belong to this product_id
	book, err := bookModel.ReadByID(bookISBN)
	if (book == form.Book{}) || err != nil {
		log.Println("No book belong to this ID.")
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No book belong to this ISBN.",
		})
		return
	}

	err = bookModel.DeleteByID(bookISBN)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while deleting book.",
		})
		return
	}

	log.Println("Delete book successfully")
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete book successfully.",
	})
}
