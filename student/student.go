package student

import (
	structerr "CountStud/structerr"

	"github.com/google/uuid"
)

type Student struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Gender    string    `json:"gender"`
	Address   string    `json:"address"`
	IIN       int       `json:"iin"`
}

// func NewUser(name, firstName, lastName, gender, address string, iin int) *User {
// 	return &User{
// 		Id:        uuid.New(),
// 		Name:      name,
// 		FirstName: firstName,
// 		LastName:  lastName,
// 		Gender:    gender,
// 		Address:   address,
// 		IIN:       iin,
// 	}
// }

func NewUser() *User {
	return &User{}
}

// func NewEmptyUser() *User {
// 	return &User{
// 		Id: uuid.New(),
// 	}
// }

// Methods for obtaining data

func (n *User) GetName() (name, firstName, lastName string) {
	return n.Name, n.FirstName, n.LastName
}

func (n *User) GetAddress() (address string) {
	return n.Address
}

func (n *User) GetGender() (gender string) {
	return n.Gender
}

func (n *User) GetIIN() (iin int) {
	return n.IIN
}

// Methods for changing user data
func (n *User) ValidateID(id uuid.UUID) error {
	if n.Id != id {
		e := structerr.NewErr("Validation error id")
		return e
	}
	return nil
}

func (n *User) ChangeName(newName string) error {
	if newName == "" {
		return structerr.NewErr("Name cannot be empty")
	}
	n.Name = newName
	return nil
}

func (n *User) ChangeLastName(newLastName string) error {
	if newLastName == "" {
		return structerr.NewErr("LastName cannot be empty")
	}
	n.LastName = newLastName
	return nil
}

func (n *User) ChangeAddress(newAddress string) error {
	if newAddress == "" {
		return structerr.NewErr("Address cannot be empty")
	}
	n.Address = newAddress
	return nil
}
