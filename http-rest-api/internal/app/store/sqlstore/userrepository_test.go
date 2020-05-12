package sqlstore

import (
	"testing"

	"github.com/lithium555/GolangStudying/http-rest-api/internal/app/errors"
	"github.com/lithium555/GolangStudying/http-rest-api/internal/app/model"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := TestDB(t, databaseURL)
	defer teardown("users")

	s := New(db)
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u.ID)
}

func TestUserRepository_Find(t *testing.T) {
	db, teardown := TestDB(t, databaseURL)
	defer teardown("users")

	s := New(db)
	u1 := model.TestUser(t)
	s.User().Create(u1)
	u2, err := s.User().Find(u1.ID)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := TestDB(t, databaseURL)
	defer teardown("users")

	s := New(db)
	u1 := model.TestUser(t)
	_, err := s.User().FindByEmail(u1.Email)
	assert.EqualError(t, err, errors.ErrRecordNotFound.Error())

	s.User().Create(u1)
	u2, err := s.User().FindByEmail(u1.Email)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
}
