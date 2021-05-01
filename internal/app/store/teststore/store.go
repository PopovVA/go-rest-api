package teststore

import (
	"github.com/gopherschool/http-rest-api/internal/app/model"
	"github.com/gopherschool/http-rest-api/internal/app/store"
)

//Store is ...
type Store struct {
	userRepository *UserRepository
}

//New is ...
func New() *Store {
	return &Store{}
}

//User is ...
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[string]*model.User),
	}

	return s.userRepository
}
