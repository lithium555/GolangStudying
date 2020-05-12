package teststore

import (
	"github.com/lithium555/GolangStudying/http-rest-api/internal/app/errors"
	"github.com/lithium555/GolangStudying/http-rest-api/internal/app/model"
)

// UserRepository ...
type UserRepository struct {
	store *Store
	users map[int]*model.User
}

// Create ...
func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}
	if err := u.BeforeCreate(); err != nil {
		return err
	}

	r.users[u.ID] = u
	u.ID = len(r.users)

	return nil
}

// FindByEmail ...
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	for _, u := range r.users {
		if u.Email == email {
			return u, nil
		}
	}

	return nil, errors.ErrRecordNotFound
}

// Find ...
func (r *UserRepository) Find(id int) (*model.User, error) {
	u, ok := r.users[id]
	if !ok {
		return nil, errors.ErrRecordNotFound
	}

	return u, nil
}
