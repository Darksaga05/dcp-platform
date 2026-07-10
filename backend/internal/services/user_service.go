package services

import "github.com/dcp-project/backend/internal/models"

func CreateUser(username string, email string) models.User {

	user := models.User{
		Username: username,
		Email:    email,
	}

	return user
}
