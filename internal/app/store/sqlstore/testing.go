package sqlstore

import (
	"database/sql"
	"fmt"
	"strings"
	"testing"
)

//TestDB checks if the database is working fine
func TestDB(t *testing.T, databaseURL string) (*sql.DB, func(...string)) {
	t.Helper()

	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		t.Fatal()
	}

	if err := db.Ping(); err != nil {
		t.Fatal()
	}

	return db, func(tables ...string) {
		if len(tables) > 0 {
			db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", ")))
		}

		db.Close()
	}
}
