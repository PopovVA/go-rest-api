package store

import (
	"database/sql"

	_ "github.com/lib/pq" //...
)

//Store is ...
type Store struct {
	config *Config
	db     *sql.DB
}

//New is ...
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

//Open is ...
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

//Close is ...
func (s *Store) Close() {
	s.db.Close()
}
