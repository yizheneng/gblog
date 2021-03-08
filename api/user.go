package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/yizheneng/gblog/middleware"
	"github.com/yizheneng/gblog/model"
)

// 添加用户
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

// 检查Token
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

// 获取用户列表
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.PostForm("pageSize"))
	pageIndex, _ := strconv.Atoi(c.PostForm("pageIndex"))

	if pageSize <= 0 || pageSize > 1000 {
		c.JSON(http.StatusOK, gin.H{
			"status":  "Error",
			"message": "pageSize should > 0 and < 1000",
		})
		return
	}

	users, total, err := model.GetUsers(pageSize, pageIndex)

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
		"data":    users,
		"total":   total,
	})
}

// 删除用户
func Delete(c *gin.Context) {

}

// 获取用户信息
func GetUserInfo(c *gin.Context) {
	tokenUsername := c.GetString("token_username")
	userInfo, err := model.GetUserInfo(tokenUsername)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "Error",
			"message": err.Error(),
		})
		return
	}

	jsonStr, _ := json.Marshal(userInfo)

	c.JSON(http.StatusOK, gin.H{
		"status":  "succeed",
		"message": "",
		"data":    string(jsonStr),
	})
}

// 更新用户信息
func UpdateUserInfo(c *gin.Context) {
	email := c.PostForm("email")
	phone := c.PostForm("phone")
	tokenUsername := c.GetString("token_username")

	err := model.UpdateUserInfo(tokenUsername, email, phone)
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

// 修改密码
func ChangePassword(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	newPassword := c.PostForm("new_password")
	tokenUsername := c.GetString("token_username")

	if username != tokenUsername {
		c.JSON(http.StatusOK, gin.H{
			"status":  "Error",
			"message": "token error",
		})
		return
	}

	err := model.ChangePassword(username, password, newPassword)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "Error",
			"message": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  "Succeed",
			"message": "Change password succeed",
		})
		return
	}
}
