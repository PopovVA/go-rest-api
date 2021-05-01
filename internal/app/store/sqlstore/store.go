package sqlstore

import (
	"database/sql"

	"github.com/gopherschool/http-rest-api/internal/app/store"
	_ "github.com/lib/pq" //...
)

//Store is ...
type Store struct {
	db             *sql.DB
	userRepository *UserRepository
}

//New is ...
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

//User is ...
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}
