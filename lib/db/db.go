package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"../../conf"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open("mysql", conf.Conf["betradar-db"]+"?parseTime=true")
	if nil != err {
		log.Fatalln(err)
	}
	err = DB.DB().Ping()
	if nil != err {
		log.Fatalln(err)
	}
	DB.LogMode(false)
}
