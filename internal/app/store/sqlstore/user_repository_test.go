package sqlstore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gopherschool/http-rest-api/internal/app/model"
	"github.com/gopherschool/http-rest-api/internal/app/store/sqlstore"
)

func TestUserRepository_Create(t *testing.T) {
	db, tearDown := sqlstore.TestDB(t, databaseURL)
	defer tearDown("users")
	s := sqlstore.New(db)

	u, err := s.User().Create(model.TestUser(t))
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, tearDown := sqlstore.TestDB(t, databaseURL)
	defer tearDown("users")

	s := sqlstore.New(db)
	email := "user@example.org"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	u := model.TestUser(t)
	u.Email = email
	s.User().Create(u)
	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
