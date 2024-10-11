package database

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
}

type Blogs struct {
	gorm.Model
	Author   string
	Content  string
	Likes    uint
	Dislikes uint
}
