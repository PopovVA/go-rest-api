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

//New Store create
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

//User helps get user from repository
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}
