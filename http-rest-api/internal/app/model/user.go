package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

// User represents user data struct
type User struct {
	ID                int    `json:"id"`
	Email             string `json:"email"`
	Password          string `json:"password,omitempty"` // omitempty ->>>> if pass is empty - we wouldn`t return it
	EncryptedPassword string `json:"-"`
}

// BeforeCreate ... here we will hash password before adding into db
func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		enc, err := encryptString(u.Password) // передаем изначальный пароль
		if err != nil {
			return err
		}

		u.EncryptedPassword = enc
	}
	return nil
}

// Sanitize will overwrite attributes, which which we consider private.
// Also we dont what to render them into external world.
func (u *User) Sanitize() {
	u.Password = ""
}

// ComparePassword will compare current password of user 'EncryptedPassword' with pass, which we got as parameter
func (u *User) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(password)) == nil
}

func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

// Validate
func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.By(requiredIf(u.Password == "")), validation.Length(6, 100)),
	)
}
