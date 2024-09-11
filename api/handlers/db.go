package handlers

import (
	"github.com/wikixen/blogapp/config"
	conf "github.com/wikixen/blogapp/database/config"
	"gorm.io/gorm"
)

var db *gorm.DB
var env = config.GetConfig()

func dbInit(){
	db = conf.CreateDB()
}