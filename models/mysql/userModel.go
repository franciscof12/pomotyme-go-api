package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GetAllUsers(db *gorm.DB) ([]User, error) {
	var users []User
	result := db.Find(&users)
	return users, result.Error
}

func GetUserByID(db *gorm.DB, id string) (User, error) {
	var user User
	result := db.First(&user, id)
	return user, result.Error
}

func CreateUser(db *gorm.DB, user *User) error {
	result := db.Create(user)
	return result.Error
}

func UpdateUser(db *gorm.DB, user *User) error {
	result := db.Save(user)
	return result.Error
}

func DeleteUser(db *gorm.DB, user *User) error {
	result := db.Delete(user)
	return result.Error
}
