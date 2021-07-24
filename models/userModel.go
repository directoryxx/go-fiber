package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       int64
	Name     string
	Username string
	Password string
}
