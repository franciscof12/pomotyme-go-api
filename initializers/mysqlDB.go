package initializers

import (
	models "github.com/franciscof12/pomotyme-go-api/v1/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type MySQLDatabase struct {
	DB *gorm.DB
}

func ConnectMysqlDatabase() *MySQLDatabase {
	dsn := "root@tcp(localhost:3306)/golangpomotyme?charset=utf8mb4&parseTime=True&loc=Local"

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}
	err = database.AutoMigrate(&models.User{}, &models.Task{}, &models.Pomodoro{})
	if err != nil {
		panic("Failed to migrate database!")
	}

	return &MySQLDatabase{
		DB: database,
	}
}

func (m *MySQLDatabase) GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := m.DB.Find(&users)
	return users, result.Error
}

func (m *MySQLDatabase) GetUserByID(id string) (models.User, error) {
	var user models.User
	result := m.DB.First(&user, id)
	return user, result.Error
}

func (m *MySQLDatabase) CreateUser() (models.User, error) {
	var user models.User
	result := m.DB.Create(&user)
	return user, result.Error
}

func (m *MySQLDatabase) UpdateUser() (models.User, error) {
	var user models.User
	result := m.DB.Save(&user)
	return user, result.Error
}

func (m *MySQLDatabase) DeleteUser() (models.User, error) {
	var user models.User
	result := m.DB.Delete(&user)
	return user, result.Error
}
