package teststore

import (
	"github.com/lithium555/GolangStudying/http-rest-api/internal/app/model"
	"github.com/lithium555/GolangStudying/http-rest-api/internal/app/store"
)

type Store struct {
	userRepository *UserRepository
}

func New() *Store {
	return nil
}

// User represents method, which can use our repositories in the internet
//
// So users can call store.User().Create() and so on...
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[int]*model.User),
	}

	return s.userRepository
}
