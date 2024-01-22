package model

import "testing"

func TestUser(t *testing.T) *User {
	return &User{
		Name:       "Николай",
		Surname:    "Чернышевский",
		Patronymic: "Гаврилович",
	}
}
