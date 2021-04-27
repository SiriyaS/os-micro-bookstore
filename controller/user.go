package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"os-micro-bookstore/form"
	"os-micro-bookstore/model"
)

type UserController struct{}

func (uc UserController) CreateUser(c *gin.Context) {
	log.Println("[User: CreateUser]")

	userModel := model.UserModel{}

	var request form.UserInfo
	err := c.BindJSON(&request)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error while binding to model.",
		})
		return
	}

	// encode for security?
	password := request.Password

	err = userModel.Add(request.Name, request.Email, request.Address, request.Telephone, request.UserName, password)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while creating user",
		})
	}

	log.Println("Create user successfully")
	c.JSON(http.StatusOK, gin.H{
		"message": "Create user successfully",
	})
}
