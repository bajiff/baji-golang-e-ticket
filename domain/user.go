package domain

import (
	"errors"
	"strings"
)

// ! Value Object Password
type Password struct{
	hashedValue string
}

// ! Value Object Email
type Email struct{
	value string
}

// ! Entity Users
type User struct {
	ID string
	Name string
	EmailAddress Email
	PasswordUser Password
}

func NewPassword(inputPassword string) (Password, error) {
	if len(inputPassword) < 6 {
		return Password{}, errors.New("Password minimal 6 karakter")
	}
	return Password{hashedValue: inputPassword}, nil
}

func NewEmail(inputEmail string) (Email, error) {
	if inputEmail == "" {
		return Email{}, errors.New("Error gaboleh kosong coy")
	}
	if !strings.Contains(inputEmail,"@") {
	return  Email{value: ""}, errors.New("Harus mengandung @")	
	}
	return Email{value: inputEmail}, nil
}

func domain() {
	
}