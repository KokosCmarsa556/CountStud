package student

import (
	structerr "CountStud/structerr"

	"github.com/google/uuid"
)

type Student struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Gender    string    `json:"gender"`
	Address   string    `json:"address"`
	IIN       int       `json:"iin"`
}

func NewStudent() *Student {
	return &Student{}
}

// Methods for obtaining data

func (n *Student) GetName() (name, firstName, lastName string) {
	return n.Name, n.FirstName, n.LastName
}

func (n *Student) GetAddress() (address string) {
	return n.Address
}

func (n *Student) GetGender() (gender string) {
	return n.Gender
}

func (n *Student) GetIIN() (iin int) {
	return n.IIN
}

// Methods for changing user data
func (n *Student) ValidateID(id uuid.UUID) error {
	if n.Id != id {
		e := structerr.NewErr("Validation error id")
		return e
	}
	return nil
}

func (n *Student) ChangeName(newName string) error {
	if newName == "" {
		return structerr.NewErr("Name cannot be empty")
	}
	n.Name = newName
	return nil
}

func (n *Student) ChangeLastName(newLastName string) error {
	if newLastName == "" {
		return structerr.NewErr("LastName cannot be empty")
	}
	n.LastName = newLastName
	return nil
}

func (n *Student) ChangeAddress(newAddress string) error {
	if newAddress == "" {
		return structerr.NewErr("Address cannot be empty")
	}
	n.Address = newAddress
	return nil
}
