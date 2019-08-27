package graphql

import (
	"gitlab.com/ditohafizh/demo-go-ci/api/models"
	"gitlab.com/ditohafizh/demo-go-ci/common"
)

func FindUserById(users []*models.User, userId string) *models.User {
	var userFound *models.User
	for _, user := range users {
		if user.Id == userId {
			userFound = user
			break
		}
	}

	return userFound
}

func CreateUser(username string, email string) *models.User {
	user := &models.User{
		Id:       common.GenerateRandomStr(10),
		Name:     &username,
		Username: username,
		Email:    email,
	}

	return user
}
