package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"

	"os-micro-bookstore/form"
	"os-micro-bookstore/model"
)

type UserController struct{}

func (uc UserController) VerifyTokenGitHub(c *gin.Context) {
	log.Println("[User: Verify GitHub access_token]")

	userModel := model.UserModel{}

	bearerToken := c.Request.Header["Authorization"]
	fmt.Println(bearerToken)
	tokenJoined := strings.Join(bearerToken, "")
	token := strings.Split(tokenJoined, " ")

	// --------------- request to GitHub API get Token Info ---------------------
	clientID := os.Getenv("GitHub_Client_ID")
	tokenUrl := fmt.Sprintf("https://api.github.com/applications/%s/token", clientID)

	reqBody := fmt.Sprintf(`{"access_token": "%s"}`, token[1])

	username := clientID
	passwd := os.Getenv("GitHub_Client_Secret")

	client := &http.Client{}
	req, err := http.NewRequest("POST", tokenUrl, strings.NewReader(reqBody))
	req.SetBasicAuth(username, passwd)
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Println("Invalid token")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid token.",
		})
		return
	}
	tokenBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}

	// map to struct
	tokenClaims := form.GitHubClaim{}
	err = json.Unmarshal(tokenBody, &tokenClaims)
	if err != nil {
		log.Println(err)
	}
	// log.Println("response: ", string(tokenBody))
	// log.Printf("struct: %#v", tokenClaims)

	// --------- check user exists - if does not exist create one in DB -------------
	user, err := userModel.ReadByID(tokenClaims.User.ID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while getting user",
		})
		return
	}
	if (user == form.User{}) {
		log.Println("No user belong to this UserID. Creating user in database...")
		// model create row
		userRequest := form.User{
			ID:       tokenClaims.User.ID,
			Username: tokenClaims.User.Login,
		}

		err = userModel.Add(userRequest)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error while creating user",
			})
			return
		}
		log.Println("Create user successfully.")
	}

	// get user from Database
	user, err = userModel.ReadByID(tokenClaims.User.ID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while getting user.",
		})
		return
	}

	log.Println("Get user successfully.")
	c.JSON(http.StatusOK, user)
}

// func (us UserController) VerifyTokenGoogle(c *gin.Context) {
// 	log.Println("[User: Verify Google access_token]")

// 	clientID := os.Getenv("Google_Client_ID")

// 	userModel := model.UserModel{}

// 	bearerToken := c.Request.Header["Authorization"]
// 	fmt.Println(bearerToken)
// 	tokenJoined := strings.Join(bearerToken, "")
// 	token := strings.Split(tokenJoined, " ")

// 	// --------------- request to Google API get Token Info ---------------------
// 	tokenUrl := fmt.Sprintf("https://oauth2.googleapis.com/tokeninfo?access_token=%s", token[1])

// 	resToken, err := http.Get(tokenUrl)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	defer resToken.Body.Close()

// 	if resToken.StatusCode != 200 {
// 		log.Println("Invalid token")
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": "Invalid token.",
// 		})
// 		return
// 	}
// 	tokenBody, err := ioutil.ReadAll(resToken.Body)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	// map to struct
// 	tokenClaims := form.GoogleTokenClaim{}
// 	err = json.Unmarshal(tokenBody, &tokenClaims)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	log.Println("response: ", string(tokenBody))
// 	// log.Printf("struct: %#v", tokenClaims)

// 	// check ClientID
// 	if tokenClaims.Aud != clientID {
// 		log.Println("Client ID not match")
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": "Does not have access. Please check your Client ID.",
// 		})
// 		return
// 	}

// 	// check expiration time
// 	expire, err := strconv.ParseInt(tokenClaims.Exp, 10, 64)
// 	if err != nil {
// 		log.Println(err)
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": "Does not have access. Please check your Client ID.",
// 		})
// 		return
// 	}

// 	// fmt.Println(time.Unix(expire, 0))
// 	// fmt.Println((time.Now()).After(time.Unix(expire, 0)))
// 	if (time.Now()).After(time.Unix(expire, 0)) {
// 		log.Println("Access Token has expired")
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": "Does not have access. Access Token has expired.",
// 		})
// 		return
// 	}

// 	// ----------- request to Google API get User Info --------------
// 	userUrl := fmt.Sprintf("https://www.googleapis.com/oauth2/v2/userinfo?access_token=%s", token[1])

// 	resUser, err := http.Get(userUrl)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	defer resUser.Body.Close()

// 	if resUser.StatusCode != 200 {
// 		log.Println("Invalid token")
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": "Invalid token.",
// 		})
// 		return
// 	}
// 	userBody, err := ioutil.ReadAll(resUser.Body)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	// map to struct
// 	userClaims := form.GoogleUserClaim{}
// 	err = json.Unmarshal(userBody, &userClaims)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	log.Println("response: ", string(userBody))
// 	// log.Printf("struct: %#v", userClaims)

// 	// --------- check user exists - if does not exist create one in DB -------------
// 	user, err := userModel.ReadBySubID(tokenClaims.Sub)
// 	if err != nil {
// 		log.Println(err)
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"message": "Error while getting user",
// 		})
// 		return
// 	}
// 	if (user == form.User{}) {
// 		log.Println("No user belong to this UserID. Creating user in database...")
// 		// model create row
// 		userRequest := form.UserRequest{
// 			UserSubID: userClaims.ID,
// 			FirstName: userClaims.GivenName,
// 			LastName:  userClaims.FamilyName,
// 			Email:     userClaims.Email,
// 		}

// 		err = userModel.Add(userRequest)
// 		if err != nil {
// 			log.Println(err)
// 			c.JSON(http.StatusInternalServerError, gin.H{
// 				"message": "Error while creating user",
// 			})
// 			return
// 		}
// 		log.Println("Create user successfully.")
// 	}

// 	// get user from Database
// 	user, err = userModel.ReadBySubID(userClaims.ID)
// 	if err != nil {
// 		log.Println(err)
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"message": "Error while getting user.",
// 		})
// 		return
// 	}

// 	log.Println("Get user successfully.")
// 	c.JSON(http.StatusOK, user)
// }

// func (uc UserController) GetUserInfo(c *gin.Context) {
// 	log.Println("[User: GetUserInfo]")
// 	userModel := model.UserModel{}

// 	userSub := c.Query("sub")

// 	user, err := userModel.ReadBySubID(userSub)
// if err != nil {
// 	log.Println(err)
// 	c.JSON(http.StatusInternalServerError, gin.H{
// 		"message": "Error while getting user",
// 	})
// 	return
// }
// 	if (user == form.User{}) {
// 		log.Println("No user belong to this UserID.")
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"message": "No user belong to this UserID.",
// 		})
// 		return
// 	}

// log.Println("Get user successfully")
// c.JSON(http.StatusOK, user)
// }

// func (uc UserController) CreateUser(c *gin.Context) {
// 	log.Println("[User: CreateUser]")

// 	userModel := model.UserModel{}

// 	var request form.UserInfoRequest
// 	err := c.BindJSON(&request)
// 	if err != nil {
// 		log.Println(err)
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": "Error while binding to model.",
// 		})
// 		return
// 	}

// 	// encode for security
// 	hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	request.Password = string(hashPassword)

// 	err = userModel.Add(request)
// if err != nil {
// 	log.Println(err)
// 	c.JSON(http.StatusInternalServerError, gin.H{
// 		"message": "Error while creating user",
// 	})
// 	return
// }

// 	log.Println("Create user successfully")
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Create user successfully",
// 	})
// }

// func (uc UserController) Login(c *gin.Context) {
// 	log.Println("[User: Login]")

// 	userModel := model.UserModel{}

// 	var request form.UserLoginRequest
// 	err := c.BindJSON(&request)
// 	if err != nil {
// 		log.Println(err)
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": "Error while binding to model.",
// 		})
// 		return
// 	}

// 	// get userInfo to check password
// 	user, err := userModel.ReadByEmail(request.Email)
// 	if err != nil {
// 		log.Println(err)
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": "Error while logging in.",
// 		})
// 		return
// 	}

// 	// verify password
// 	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
// 	if err != nil {
// 		log.Println(err)
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"message": "Cannot login. Wrong email or password.",
// 		})
// 		return
// 	}

// 	log.Println("Log in successfully")
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Log in successfully",
// 	})
// }

// func (uc UserController) EditProfileByEmail(c *gin.Context) {
// 	log.Println("[User: EditProfileByEmail]")

// 	userModel := model.UserModel{}

// 	var request form.UserInfoRequest
// 	err := c.BindJSON(&request)
// 	if err != nil {
// 		log.Println(err)
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": "Error while binding to model.",
// 		})
// 		return
// 	}

// 	user, err := userModel.ReadByEmail(request.Email)
// 	if (user == form.UserInfo{}) || err != nil {
// 		log.Println("No account belong to this email.")
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"message": "No account belong to this email.",
// 		})
// 		return
// 	}

// 	err = userModel.UpdateByEmail(request.Email, request)
// 	if err != nil {
// 		log.Println(err)
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"message": "Error while update user.",
// 		})
// 		return
// 	}

// 	log.Println("Update user successfully")
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Update user successfully.",
// 	})
// }
