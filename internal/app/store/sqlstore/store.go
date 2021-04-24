package sqlstore

import (
	"database/sql"

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
func (s *Store) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}
