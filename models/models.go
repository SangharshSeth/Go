package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	UserProfile UserProfile
}
type Likes []string

type UserProfile struct {
	gorm.Model
	Sex    string  `json:"sex" gorm:"not null"`
	Age    int     `json:"age" gorm:"not null"`
	Height float64 `json:"height"`
	Weight int     `json:"weight"`
	Likes  Likes   `gorm:"serializer:json" json:"likes"`
	UserId uint
}
