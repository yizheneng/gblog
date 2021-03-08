package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yizheneng/gblog/api"
	"github.com/yizheneng/gblog/config"
	"github.com/yizheneng/gblog/middleware"
	_ "github.com/yizheneng/gblog/model"
)

func initRoute() {
	gin.SetMode("debug")

	r := gin.New()

	r.Static("/static", "./static/")

	// r.GET("/", func(c *gin.Context) {
	// 	c.HTML(200, "front", nil)
	// })

	loginR := r.Group("login")
	{
		loginR.POST("/", api.Login)
		loginR.POST("/backend", api.LoginToBackEnd)
	}

	userR1 := r.Group("user")
	{
		userR1.POST("add", api.AddUser)
		userR1.POST("check_token", api.CheckToken)
	}

	userR2 := r.Group("user")
	userR2.Use(middleware.JwtToken())
	{
		userR2.POST("change_password", api.ChangePassword)
		userR2.PUT("update_info", api.UpdateUserInfo)
		userR2.POST("get_userinfo", api.GetUserInfo)
	}

	r.Run(config.ServerSettings.ServerPort)
}

func main() {
	initRoute()
}
