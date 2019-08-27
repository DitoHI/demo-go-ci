package models

type Todo struct {
	Id          string
	Title       string
	Description *string
	User        *User
}
