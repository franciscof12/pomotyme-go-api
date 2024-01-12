package initializers

import (
	models "github.com/franciscof12/pomotyme-go-api/v1/models/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root@tcp(localhost:3306)/golangpomotyme?charset=utf8mb4&parseTime=True&loc=Local"

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}
	err = database.AutoMigrate(&models.User{})
	if err != nil {
		panic("Failed to migrate database!")
	}

	DB = database
}
