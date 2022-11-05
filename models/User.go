package models

import "gopkg.in/validator.v2"

type User struct {
	ID    int    `json:"id" gorm:"primary_key, AUTO_INCREMENT"`
	Name  string `json:"name" validate:"nonzero, regexp=^\\D*$"`
	Email string `json:"email" validate:"nonzero, regexp=^.+@[a-zA-Z]+\\.[a-zA-Z]+$"`
}

func (user *User) ValidateUser() error {
	if err := validator.Validate(user); err != nil {
		return err
	}
	return nil
}
