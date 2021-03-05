package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/yizheneng/gblog/middleware"
	"github.com/yizheneng/gblog/model"
)

func AddUser(c *gin.Context) {
	var user model.User
	_ = c.ShouldBindJSON(&user)

	err := validator.New().Struct(&user)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "Error",
			"message": err.Error(),
		})
		return
	}

	if model.CheckUser(user.UserName) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "Error",
			"message": "Username used",
		})
		return
	}

	err = model.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "Error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "Succeed",
		"message": "",
	})
}

func CheckToken(c *gin.Context) {
	token := c.PostForm("token")

	claims, err := middleware.CheckToken(token)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "Error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   "Succeed",
		"message":  "",
		"username": claims.UserName,
	})
}

func GetUsers() {

}

func Delete() {

}

func EditUser() {

}

func ChangePassword() {

}
