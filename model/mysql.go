package model

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/yizheneng/gblog/config"
)

var db *gorm.DB

func init() {
	dataBaseLink := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.ServerSettings.MySqlUserNamr,
		config.ServerSettings.MySqlPassword,
		config.ServerSettings.MySqlAddr,
		config.ServerSettings.MySqlDB)
	fmt.Printf(dataBaseLink)

	var err error
	db, err = gorm.Open("mysql", dataBaseLink)

	if err != nil {
		fmt.Printf("Connect database error:%s", err.Error())
		os.Exit(1)
	}
	// defer db.Close()

	db.AutoMigrate(&Article{}, &Category{}, &User{})
}
