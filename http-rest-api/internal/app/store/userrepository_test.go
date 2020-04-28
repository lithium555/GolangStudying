package store

import (
	"github.com/lithium555/GolangStudying/http-rest-api/internal/app/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := TestStore(t, databaseURL)
	defer teardown("users")

	u, err := s.User().Create(&model.User{
		Email: "user@example.org",
	})

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := TestStore(t, databaseURL)
	defer teardown("users")

	// Example_1 : user doesnt exist in db
	email := "user@example.org"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	// Example_2: user exists in db
	u, err := s.User().Create(&model.User{
		Email: "user@example.org",
	})

	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
