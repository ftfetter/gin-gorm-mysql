package models

import "gopkg.in/validator.v2"

type User struct {
	ID    int    `json:"id" gorm:"primary_key, AUTO_INCREMENT"`
	Name  string `json:"name" validate:"nonzero, regexp=^\\D*$"`
	Email string `json:"email" validate:"nonzero, regexp=^.+@[a-zA-Z]+\\.[a-zA-Z]+$"`
}

func (u1 *User) CopyFrom(u2 User) {
	if u2.Email != "" {
		u1.Email = u2.Email
	}
	if u2.Name != "" {
		u1.Name = u2.Name
	}
}

func (u *User) ValidateUser() error {
	if err := validator.Validate(u); err != nil {
		return err
	}
	return nil
}
