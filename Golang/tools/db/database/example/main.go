package main

import (
	"cookbook/chapter05/database"

	_ "github.com/go-sql-driver/mysql" //we import supported libraries for database/sql
)

func main() {
	db, err := database.Setup()
	if err != nil {
		panic(err)
	}

	if err := database.Exec(db); err != nil {
		panic(err)
	}
}
