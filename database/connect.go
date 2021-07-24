package database

import (
	"github.com/directoryxx/go-fiber/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	// if cant reach db try add "()"
	conn, err := gorm.Open(mysql.Open("root:password@(mysql-server)/dev"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	conn.AutoMigrate(models.User{})
}
