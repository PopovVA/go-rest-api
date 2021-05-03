package store

import (
	"github.com/gopherschool/http-rest-api/internal/app/model"
)

//UserRepository is ...
type UserRepository interface {
	Create(*model.User) error
	Find(int) (*model.User, error)
	FindByEmail(string) (*model.User, error)
}
