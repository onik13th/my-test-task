package model

import "github.com/go-playground/validator/v10"

type User struct {
	ID            int
	Name          string `validate:"required,alphaunicode"`
	Surname       string `validate:"required,alphaunicode"`
	Patronymic    string `validate:"omitempty,alphaunicode"`
	Age           int    `json:"Age"`
	Gender        string `json:"Gender"`
	Nationalities string `json:"Nationalities"`
}

func (u *User) Validate() error {
	validate := validator.New() // создал экземпляр валидатора
	return validate.Struct(u)   // вызвал метод Struct для валидации объекта User
}
