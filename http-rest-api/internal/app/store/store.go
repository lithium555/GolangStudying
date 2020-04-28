package store

import (
	"database/sql"

	_ "github.com/lib/pq" // database package
)

type Store struct {
	config         *Config
	db             *sql.DB
	userRepository *UserRepository
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// Open represents for initialise db, connection to db, handle any error
func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

// Close represents for decline connection to db and so on...
func (s *Store) Close() {
	s.db.Close()
}

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
