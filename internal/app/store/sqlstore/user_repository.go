package sqlstore

import (
	"database/sql"

	"github.com/gopherschool/http-rest-api/internal/app/model"
	"github.com/gopherschool/http-rest-api/internal/app/store"
)

//UserRepository contains store ref and crud methods
type UserRepository struct {
	store *Store
}

//Create user in the postgres database
func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
		u.Email, u.EncryptedPassword,
	).Scan(&u.ID)
}

//FindByEmail returns User by email
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE email = $1",
		email,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecorcNotFound
		}

		return nil, err
	}
	return u, nil
}

//Find returns User by ID
func (r *UserRepository) Find(id int) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE id = $1",
		id,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecorcNotFound
		}

		return nil, err
	}
	return u, nil
}
