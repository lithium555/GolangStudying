package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser_BeforeCreate(t *testing.T) {
	u := TestUser(t)
	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.EncryptedPassword)
}

func TestUser_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		u       func() *User // func which return user
		isValid bool
	}{
		{
			name: "valid",
			u: func() *User {
				return TestUser(t)
			},
			isValid: true,
		},
		{
			name: "with encrypt password",
			u: func() *User {
				u := TestUser(t)
				u.Password = ""
				u.EncryptedPassword = "encryptedpassword"

				return u
			},
			isValid: true,
		},
		{
			name: "empty email",
			u: func() *User {
				u := TestUser(t)
				u.Email = ""

				return u
			},
			isValid: false,
		},
		{
			name: "invalid email",
			u: func() *User {
				u := TestUser(t)
				u.Email = "invalid"

				return u
			},
			isValid: false,
		},
		{
			name: "empty password",
			u: func() *User {
				u := TestUser(t)
				u.Password = ""

				return u
			},
			isValid: false,
		},
		{
			name: "short password",
			u: func() *User {
				u := TestUser(t)
				u.Password = "short"

				return u
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u().Validate())
			} else {
				assert.Error(t, tc.u().Validate())
			}
		})
	}
}
