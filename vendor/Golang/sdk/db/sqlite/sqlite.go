package sqlite

import (
	"database/sql"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strings"

	"github.com/mattn/go-sqlite3"
	_ "github.com/mattn/go-sqlite3"
	sqlite "github.com/mattn/go-sqlite3"
)

func SimpleDB() {
	os.Remove("./foo.db")

	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	create table foo (id integer not null primary key, name text);
	delete from foo;
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into foo(id, name) values(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	for i := 0; i < 100; i++ {
		_, err = stmt.Exec(i, fmt.Sprintf("こんにちわ世界%03d", i))
		if err != nil {
			log.Fatal(err)
		}
	}
	tx.Commit()

	rows, err := db.Query("select id, name from foo")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err = db.Prepare("select name from foo where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var name string
	err = stmt.QueryRow("3").Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)

	_, err = db.Exec("delete from foo")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("insert into foo(id, name) values(1, 'foo'), (2, 'bar'), (3, 'baz')")
	if err != nil {
		log.Fatal(err)
	}

	rows, err = db.Query("select id, name from foo")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func HookExample() {
	sqlite3conn := []*sqlite3.SQLiteConn{}
	sql.Register("sqlite3_with_hook_example",
		&sqlite3.SQLiteDriver{
			ConnectHook: func(conn *sqlite3.SQLiteConn) error {
				sqlite3conn = append(sqlite3conn, conn)
				conn.RegisterUpdateHook(func(op int, db string, table string, rowid int64) {
					switch op {
					case sqlite3.SQLITE_INSERT:
						log.Println("Notified of insert on db", db, "table", table, "rowid", rowid)
					}
				})
				return nil
			},
		})
	os.Remove("./foo.db")
	os.Remove("./bar.db")

	srcDb, err := sql.Open("sqlite3_with_hook_example", "./foo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer srcDb.Close()
	srcDb.Ping()

	_, err = srcDb.Exec("create table foo(id int, value text)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = srcDb.Exec("insert into foo values(1, 'foo')")
	if err != nil {
		log.Fatal(err)
	}
	_, err = srcDb.Exec("insert into foo values(2, 'bar')")
	if err != nil {
		log.Fatal(err)
	}
	_, err = srcDb.Query("select * from foo")
	if err != nil {
		log.Fatal(err)
	}
	destDb, err := sql.Open("sqlite3_with_hook_example", "./bar.db")
	if err != nil {
		log.Fatal(err)
	}
	defer destDb.Close()
	destDb.Ping()

	bk, err := sqlite3conn[1].Backup("main", sqlite3conn[0], "main")
	if err != nil {
		log.Fatal(err)
	}

	_, err = bk.Step(-1)
	if err != nil {
		log.Fatal(err)
	}
	_, err = destDb.Query("select * from foo")
	if err != nil {
		log.Fatal(err)
	}
	_, err = destDb.Exec("insert into foo values(3, 'bar')")
	if err != nil {
		log.Fatal(err)
	}

	bk.Finish()
}

// Computes x^y
func pow(x, y int64) int64 {
	return int64(math.Pow(float64(x), float64(y)))
}

// Computes the bitwise exclusive-or of all its arguments
func xor(xs ...int64) int64 {
	var ret int64
	for _, x := range xs {
		ret ^= x
	}
	return ret
}

// Returns a random number. It's actually deterministic here because
// we don't seed the RNG, but it's an example of a non-pure function
// from SQLite's POV.
func getrand() int64 {
	return rand.Int63()
}

// Computes the standard deviation of a GROUPed BY set of values
type stddev struct {
	xs []int64
	// Running average calculation
	sum int64
	n   int64
}

func newStddev() *stddev { return &stddev{} }

func (s *stddev) Step(x int64) {
	s.xs = append(s.xs, x)
	s.sum += x
	s.n++
}

func (s *stddev) Done() float64 {
	mean := float64(s.sum) / float64(s.n)
	var sqDiff []float64
	for _, x := range s.xs {
		sqDiff = append(sqDiff, math.Pow(float64(x)-mean, 2))
	}
	var dev float64
	for _, x := range sqDiff {
		dev += x
	}
	dev /= float64(len(sqDiff))
	return math.Sqrt(dev)
}

func CustomFunc() {
	sql.Register("sqlite3_custom", &sqlite.SQLiteDriver{
		ConnectHook: func(conn *sqlite.SQLiteConn) error {
			if err := conn.RegisterFunc("pow", pow, true); err != nil {
				return err
			}
			if err := conn.RegisterFunc("xor", xor, true); err != nil {
				return err
			}
			if err := conn.RegisterFunc("rand", getrand, false); err != nil {
				return err
			}
			if err := conn.RegisterAggregator("stddev", newStddev, true); err != nil {
				return err
			}
			return nil
		},
	})

	db, err := sql.Open("sqlite3_custom", ":memory:")
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}
	defer db.Close()

	var i int64
	err = db.QueryRow("SELECT pow(2,3)").Scan(&i)
	if err != nil {
		log.Fatal("POW query error:", err)
	}
	fmt.Println("pow(2,3) =", i) // 8

	err = db.QueryRow("SELECT xor(1,2,3,4,5,6)").Scan(&i)
	if err != nil {
		log.Fatal("XOR query error:", err)
	}
	fmt.Println("xor(1,2,3,4,5) =", i) // 7

	err = db.QueryRow("SELECT rand()").Scan(&i)
	if err != nil {
		log.Fatal("RAND query error:", err)
	}
	fmt.Println("rand() =", i) // pseudorandom

	_, err = db.Exec("create table foo (department integer, profits integer)")
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}
	_, err = db.Exec("insert into foo values (1, 10), (1, 20), (1, 45), (2, 42), (2, 115)")
	if err != nil {
		log.Fatal("Failed to insert records:", err)
	}

	rows, err := db.Query("select department, stddev(profits) from foo group by department")
	if err != nil {
		log.Fatal("STDDEV query error:", err)
	}
	defer rows.Close()
	for rows.Next() {
		var dept int64
		var dev float64
		if err := rows.Scan(&dept, &dev); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("dept=%d stddev=%f\n", dept, dev)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func createBulkInsertQuery(n int, start int) (query string, args []interface{}) {
	values := make([]string, n)
	args = make([]interface{}, n*2)
	pos := 0
	for i := 0; i < n; i++ {
		values[i] = "(?, ?)"
		args[pos] = start + i
		args[pos+1] = fmt.Sprintf("こんにちわ世界%03d", i)
		pos += 2
	}
	query = fmt.Sprintf(
		"insert into foo(id, name) values %s",
		strings.Join(values, ", "),
	)
	return
}

func bulkInsert(db *sql.DB, query string, args []interface{}) (err error) {
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}

	_, err = stmt.Exec(args...)
	if err != nil {
		return
	}

	return
}

func Limit() {
	var sqlite3conn *sqlite3.SQLiteConn
	sql.Register("sqlite3_with_limit", &sqlite3.SQLiteDriver{
		ConnectHook: func(conn *sqlite3.SQLiteConn) error {
			sqlite3conn = conn
			return nil
		},
	})

	os.Remove("./foo.db")
	db, err := sql.Open("sqlite3_with_limit", "./foo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	create table foo (id integer not null primary key, name text);
	delete from foo;
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	if sqlite3conn == nil {
		log.Fatal("not set sqlite3 connection")
	}

	limitVariableNumber := sqlite3conn.GetLimit(sqlite3.SQLITE_LIMIT_VARIABLE_NUMBER)
	log.Printf("default SQLITE_LIMIT_VARIABLE_NUMBER: %d", limitVariableNumber)

	num := 400
	query, args := createBulkInsertQuery(num, 0)
	err = bulkInsert(db, query, args)
	if err != nil {
		log.Fatal(err)
	}

	smallLimitVariableNumber := 100
	sqlite3conn.SetLimit(sqlite3.SQLITE_LIMIT_VARIABLE_NUMBER, smallLimitVariableNumber)

	limitVariableNumber = sqlite3conn.GetLimit(sqlite3.SQLITE_LIMIT_VARIABLE_NUMBER)
	log.Printf("updated SQLITE_LIMIT_VARIABLE_NUMBER: %d", limitVariableNumber)

	query, args = createBulkInsertQuery(num, num)
	err = bulkInsert(db, query, args)
	if err != nil {
		if err != nil {
			log.Printf("expect failed since SQLITE_LIMIT_VARIABLE_NUMBER is too small: %v", err)
		}
	}

	bigLimitVariableNumber := 999999
	sqlite3conn.SetLimit(sqlite3.SQLITE_LIMIT_VARIABLE_NUMBER, bigLimitVariableNumber)
	limitVariableNumber = sqlite3conn.GetLimit(sqlite3.SQLITE_LIMIT_VARIABLE_NUMBER)
	log.Printf("set SQLITE_LIMIT_VARIABLE_NUMBER: %d", bigLimitVariableNumber)
	log.Printf("updated SQLITE_LIMIT_VARIABLE_NUMBER: %d", limitVariableNumber)

	query, args = createBulkInsertQuery(500, num+num)
	err = bulkInsert(db, query, args)
	if err != nil {
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("no error if SQLITE_LIMIT_VARIABLE_NUMBER > 999")
}
