package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql" //we import supported libraries for database/sql
)

// Example hold the results of our queries
type Example struct {
	Name    string
	Created *time.Time
}

// DB is an interface that is satisfied
// by an sql.DB or an sql.Transaction
type DB interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

// Transaction can do anything a Query can do
// plus Commit, Rollback, or Stmt
type Transaction interface {
	DB
	Commit() error
	Rollback() error
	Stmt(stmt *sql.Stmt) *sql.Stmt
}

// Setup configures and returns our database
// connection poold
func Setup() (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/gocookbook?parseTime=true", os.Getenv("MYSQLUSERNAME"), os.Getenv("MYSQLPASSWORD")))
	if err != nil {
		return nil, err
	}
	// there will only ever be 24 open connections
	db.SetMaxOpenConns(24)

	// MaxIdleConns can never be less than max open SetMaxOpenConns
	// otherwise it'll default to that value
	db.SetMaxIdleConns(24)
	return db, nil
}

// Create makes a table called example
// and populates it
func Create(db *sql.DB) error {
	// create the database
	if _, err := db.Exec("CREATE TABLE example (name VARCHAR(20), created DATETIME)"); err != nil {
		return err
	}

	if _, err := db.Exec(`INSERT INTO example (name, created) values ("Aaron", NOW())`); err != nil {
		return err
	}

	return nil
}

// Exec takes a new connection
// creates tables, and later drops them
// and issues some queries
func Exec(db *sql.DB) error {
	// uncaught error on cleanup, but we always
	// want to cleanup
	defer db.Exec("DROP TABLE example")

	if err := Create(db); err != nil {
		return err
	}

	if err := Query(db); err != nil {
		return err
	}
	return nil
}

// Query grabs a new connection
// creates tables, and later drops them
// and issues some queries
func Query(db *sql.DB) error {
	name := "Aaron"
	rows, err := db.Query("SELECT name, created FROM example where name=?", name)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var e Example
		if err := rows.Scan(&e.Name, &e.Created); err != nil {
			return err
		}
		fmt.Printf("Results:\n\tName: %s\n\tCreated: %v\n", e.Name, e.Created)
	}
	return rows.Err()
}

// ExecWithTimeout will timeout trying
// to get the current time
func ExecWithTimeout() error {
	db, err := Setup()
	if err != nil {
		return err
	}

	ctx := context.Background()

	// we want to timeout immediately
	ctx, can := context.WithDeadline(ctx, time.Now())

	// call cancel after we complete
	defer can()

	// our transaction is context aware
	_, err = db.BeginTx(ctx, nil)
	return err
}
