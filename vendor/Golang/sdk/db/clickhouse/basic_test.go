package clickhouse

import "testing"

func TestSimpleInsert(t *testing.T) {
	db := OpenDB()
	InsertSimple(db)
}
