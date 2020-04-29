package model

import "testing"

func TestUser(t *testing.T) *User { // we need it for creating user with valid parameters in tests
	return &User{
		Email:    "user@example.org",
		Password: "password",
	}
}
