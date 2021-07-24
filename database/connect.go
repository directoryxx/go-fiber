package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	// if cant reach db try add "()"
	_, err := gorm.Open(mysql.Open("root:password@(mysql-server)/dev"), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}
