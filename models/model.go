package models

import (
	"gorm.io/gorm"
)

type User struct {
	//model of user
	//gorm.Model is automatically create a

	//ID        uint `gorm:"primarykey"`
	//	CreatedAt time.Time
	//	UpdatedAt time.Time
	//	DeletedAt `gorm:"index"`'

	gorm.Model
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `gorm:"unique"`
	Password string `json:"password"`
	Role     bool   `json:"role"`
}

type Course struct {
	Name  string `json:"name"`
	Level string `json:"level"`
}
