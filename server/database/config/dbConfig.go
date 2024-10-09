package database

import (
	"log"

	"github.com/glebarez/sqlite"
	"github.com/wikixen/blogapp/config"
	"gorm.io/gorm"
	models "github.com/wikixen/blogapp/database/models"
)

func CreateDB() *gorm.DB {
	errText := "Failure to connect to database:\n"
	env := config.GetConfig()
	db, err := gorm.Open(sqlite.Open(env.Database))
	if err != nil {
		log.Fatalln(errText, err)
	}

	db.AutoMigrate(&models.Blogs{},&models.User{})

	return db
}
