package entities

import "gopkg.in/go-playground/validator.v9"

type Fullname struct {
	Firstname  string `json:"firstname"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

type Contact struct {
	ID       int      `json:"id"`
	fullname Fullname `json:"fullname"`
	Phone    string   `json:"phone" validate:"number"`
}

func (c *Contact) Value() Fullname {
	return c.fullname
}

var contactCounter int

func NewContact(firstname, surname, patronymic, phone string) *Contact {
	contactCounter++
	name := &Fullname{
		Firstname:  firstname,
		Surname:    surname,
		Patronymic: patronymic,
	}

	return &Contact{
		ID:       contactCounter,
		fullname: *name,
		Phone:    phone,
	}
}

func (c *Contact) Validate() error {
	validate := validator.New()
	err := validate.Struct(c)
	if err != nil {
		return err
	}
	return nil
}
