package database

import (
	"golang-auth/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:@/auth_apps"), &gorm.Config{})

	if err != nil {
		panic("Tidak konek ke database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}
