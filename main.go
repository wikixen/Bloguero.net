package main

import (
	"github.com/glebarez/sqlite"
	_ "github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Blogs struct {
	gorm.Model
	Title   string
	Author  string
	Content string
}

var db, err = gorm.Open(sqlite.Open("blogs.db"), &gorm.Config{})

func main() {
	db.AutoMigrate(&Blogs{})

}
