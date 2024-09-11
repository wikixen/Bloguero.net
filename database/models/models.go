package database

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
}

type Blogs struct {
	gorm.Model
	Title    string
	Author   string
	Content  string
	Likes    uint
	Dislikes uint
}
