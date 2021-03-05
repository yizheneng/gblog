package config

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

type serverSettings struct {
	ServerPort string

	MySqlAddr     string
	MySqlPort     string
	MySqlDB       string
	MySqlUserNamr string
	MySqlPassword string

	JwtKey []byte
}

var ServerSettings serverSettings

func init() {
	setting, err := ini.Load("server.ini")

	if err != nil {
		fmt.Printf("Open server settings file error:", err)
		os.Exit(1)
	}

	ServerSettings.ServerPort = setting.Section("server").Key("port").String()

	ServerSettings.MySqlAddr = setting.Section("mysql").Key("addr").String()
	ServerSettings.MySqlPort = setting.Section("mysql").Key("port").String()
	ServerSettings.MySqlDB = setting.Section("mysql").Key("db").String()
	ServerSettings.MySqlUserNamr = setting.Section("mysql").Key("user").String()
	ServerSettings.MySqlPassword = setting.Section("mysql").Key("password").String()

	ServerSettings.JwtKey = []byte(setting.Section("jwt").Key("key").String())
}
