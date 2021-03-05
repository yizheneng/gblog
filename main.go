package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yizheneng/gblog/api"
	"github.com/yizheneng/gblog/config"
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

	userR := r.Group("user")
	{
		userR.POST("add", api.AddUser)
		userR.POST("check_token", api.CheckToken)
	}

	r.Run(config.ServerSettings.ServerPort)
}

func main() {
	initRoute()
}
