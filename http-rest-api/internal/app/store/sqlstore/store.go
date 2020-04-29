package sqlstore

import (
	"database/sql"

	_ "github.com/lib/pq" // database package
)

// Store ...
type Store struct {
	db             *sql.DB
	userRepository *UserRepository
}

// New ...
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// User ...
// User represents method, which can use our repositories in the internet
//
// So users can call store.User().Create() and so on...
func (s *Store) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}
