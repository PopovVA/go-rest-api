package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

//User is ...
type User struct {
	ID                int    `json:"id"`
	Email             string `json:"email"`
	Password          string `json:"password,omitempty"`
	EncryptedPassword string `json:"-"`
}

//Validate is ...
func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(
			&u.Email,
			validation.Required,
			is.Email,
		),
		validation.Field(
			&u.Password,
			validation.By(requiredIf(u.EncryptedPassword == "")),
			validation.Length(6, 100),
		),
	)
}

//Sanitaze is ...
func (u *User) Sanitaze() {
	u.Password = ""
}

//BeforeCreate ...
func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		enc, err := ecnryptString(u.Password)
		if err != nil {
			return err
		}

		u.EncryptedPassword = enc
	}

	return nil
}

func ecnryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
