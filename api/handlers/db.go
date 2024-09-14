package handlers

import (
	conf "github.com/wikixen/blogapp/database/config"
	"gorm.io/gorm"
)

var db *gorm.DB

func dbInit() {
	db = conf.CreateDB()
}
