package users

import (
	"CountStud/structerr"
	"unicode"

	"github.com/google/uuid"
)

type Users struct {
	Id       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Password string    `json:"_"`
	Name     string    `json:"name"`
	SurName  string    `json:"surname"`
	LastName string    `json:"lastname"`
	Role     string    `json:"role"`
}

func NewUser() *Users {
	return &Users{}
}

//GETTER

func (u *Users) GetFullName() string {
	fullName := u.LastName + u.Name + u.SurName
	return fullName
}

func (u *Users) GetDataUser() (login, pass string) {
	return u.Email, u.Password
}

//SETTER

func (u *Users) ChangeLastName(newLastName string) error {
	if newLastName == "" {
		return structerr.NewErr("The new surname is empty")
	}
	u.LastName = newLastName
	return nil
}

func (u *Users) ChangeName(newName string) error {
	if newName == "" {
		return structerr.NewErr("Name cannot be empty")
	}
	u.Name = newName
	return nil
}

func (u *Users) ChangePassword(newPass string) error {
	for _, r := range newPass {
		if unicode.Is(unicode.Cyrillic, r) {
			return structerr.NewErr("The password contains Cyrillic characters, it should only be in Latin.")
		}
	}
	if newPass == "" {
		return structerr.NewErr("New password is empty")
	} else if len(newPass) >= 10 {
		return structerr.NewErr("The new password is longer than 10 characters")
	}

	u.Password = newPass
	return nil
}
