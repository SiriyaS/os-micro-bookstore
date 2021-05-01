package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"os-micro-bookstore/form"
	"os-micro-bookstore/model"

	"golang.org/x/crypto/bcrypt"
)

type UserController struct{}

func (uc UserController) CreateUser(c *gin.Context) {
	log.Println("[User: CreateUser]")

	userModel := model.UserModel{}

	var request form.UserInfoRequest
	err := c.BindJSON(&request)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error while binding to model.",
		})
		return
	}

	// encode for security
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	request.Password = string(hashPassword)

	err = userModel.Add(request)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while creating user",
		})
		return
	}

	log.Println("Create user successfully")
	c.JSON(http.StatusOK, gin.H{
		"message": "Create user successfully",
	})
}

func (uc UserController) Login(c *gin.Context) {
	log.Println("[User: Login]")

	userModel := model.UserModel{}

	var request form.UserLoginRequest
	err := c.BindJSON(&request)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error while binding to model.",
		})
		return
	}

	// get userInfo to check password
	user, err := userModel.ReadByEmail(request.Email)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error while logging in.",
		})
		return
	}

	// verify password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Cannot login. Wrong email or password.",
		})
		return
	}

	log.Println("Log in successfully")
	c.JSON(http.StatusOK, gin.H{
		"message": "Log in successfully",
	})
}

func (uc UserController) EditProfile(c *gin.Context) {
	log.Println("[User: EditProfile]")

	userModel := model.UserModel{}

	var request form.UserInfoRequest
	err := c.BindJSON(&request)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error while binding to model.",
		})
		return
	}

	user, err := userModel.ReadByEmail(request.Email)
	if (user == form.UserInfo{}) || err != nil {
		log.Println("No account belong to this email.")
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No account belong to this email.",
		})
		return
	}

	err = userModel.UpdateByEmail(request.Email, request)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while update user.",
		})
		return
	}

	log.Println("Update user successfully")
	c.JSON(http.StatusOK, gin.H{
		"message": "Update user successfully.",
	})
}
