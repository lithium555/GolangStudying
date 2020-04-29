package store

import "github.com/lithium555/GolangStudying/http-rest-api/internal/app/store/sqlstore"

// Store ...
type Store interface {
	User() sqlstore.UserRepository
}
