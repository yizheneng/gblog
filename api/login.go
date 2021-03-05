package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yizheneng/gblog/middleware"
	"github.com/yizheneng/gblog/model"
)

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	err := model.CheckPassword(username, password)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "Error",
			"message": err.Error(),
		})
		return
	}

	var token string
	token, err = middleware.CreateToken(username)

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
		"token":   token,
	})
}

func LoginToBackEnd(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	err := model.CheckPasswordBackEnd(username, password)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "Error",
			"message": err.Error(),
		})
		return
	}

	var token string
	token, err = middleware.CreateToken(username)

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
		"token":   token,
	})
}
