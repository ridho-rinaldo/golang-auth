package database

import (
	"golang-auth/models"

	// LIBRARY GORM SEMUA ACTIVITY KE DATABASE
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	// CREATE CONNECTION
	connection, err := gorm.Open(mysql.Open("root:@/auth_apps"), &gorm.Config{})

	if err != nil {
		panic("Tidak konek ke database")
	}

	DB = connection

	// AUTO MIGRATION MODEL(COLUMN)
	connection.AutoMigrate(&models.User{})
}
