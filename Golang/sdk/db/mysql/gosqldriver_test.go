package mysql

import (
	"testing"

	_ "github.com/go-sql-driver/mysql" //we import supported libraries for database/sql
)

func TestDatabase(t *testing.T) {
	db, err := Setup()
	if err != nil {
		panic(err)
	}

	if err := Exec(db); err != nil {
		panic(err)
	}
}
