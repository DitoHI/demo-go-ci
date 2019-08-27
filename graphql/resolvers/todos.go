package graphql

import (
	"gitlab.com/ditohafizh/demo-go-ci/api/models"
	"gitlab.com/ditohafizh/demo-go-ci/common"
)

func CreateTodo(title string, description *string, user *models.User) *models.Todo {
	todo := &models.Todo{
		Id:          common.GenerateRandomStr(10),
		Title:       title,
		Description: description,
		User:        user,
	}

	return todo
}
