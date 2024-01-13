package app

import models "github.com/franciscof12/pomotyme-go-api/v1/models/mysql"

type Repository interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(id string) (models.User, error)
}
