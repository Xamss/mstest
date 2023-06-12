package entities

import "gopkg.in/go-playground/validator.v9"

type Group struct {
	ID      int     `json:"id"`
	Name    string  `json:"name" validate:"max=250"`
	Contact Contact `json:"contact"`
}

func (g *Group) Validate() error {
	validate := validator.New()
	err := validate.Struct(g)
	if err != nil {
		return err
	}
	return nil
}
