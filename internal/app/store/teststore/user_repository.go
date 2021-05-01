package teststore

import (
	"github.com/gopherschool/http-rest-api/internal/app/store"

	"github.com/gopherschool/http-rest-api/internal/app/model"
)

//UserRepository is ...
type UserRepository struct {
	store *Store
	users map[string]*model.User
}

//Create is ...
func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	r.users[u.Email] = u
	u.ID = len(r.users)

	return nil
}

//FindByEmail is ...
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u, ok := r.users[email]
	if !ok {
		return nil, store.ErrRecorcNotFound
	}

	return u, nil
}
