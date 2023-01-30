package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

type Post struct {
	gorm.Model
	Title string
	Body  string
}
