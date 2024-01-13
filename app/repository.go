package app

import models "github.com/franciscof12/pomotyme-go-api/v1/models"

type Repository interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(id string) (models.User, error)
	CreateUser() (models.User, error)
	UpdateUser() (models.User, error)
	DeleteUser() (models.User, error)
}
