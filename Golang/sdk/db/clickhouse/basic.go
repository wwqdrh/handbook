package clickhouse

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/google/uuid"
)

const ddl = `
CREATE TABLE example (
	  Col1 UInt64
	, Col2 String
	, Col3 Array(UInt8)
	, Col4 DateTime
) ENGINE = Memory`

func OpenDB() driver.Conn {
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{"127.0.0.1:9000"},
		Auth: clickhouse.Auth{
			Database: "default",
			Username: "default",
			Password: "",
		},
		DialTimeout:     time.Second,
		MaxOpenConns:    10,
		MaxIdleConns:    5,
		ConnMaxLifetime: time.Hour,
	})
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

func InsertSimple(conn driver.Conn) {
	ctx := context.Background()
	if err := conn.Exec(ctx, `DROP TABLE IF EXISTS example`); err != nil {
		log.Fatal(err)
	}
	if err := conn.Exec(ctx, ddl); err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 100; i++ {
		err := conn.AsyncInsert(ctx, fmt.Sprintf(`INSERT INTO example VALUES (
			%d, '%s', [1, 2, 3, 4, 5, 6, 7, 8, 9], now()
		)`, i, "Golang SQL database driver"), false)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func example(conn clickhouse.Conn) error {
	batch, err := conn.PrepareBatch(context.Background(), "INSERT INTO example")
	if err != nil {
		return err
	}
	var (
		col1 []uint64
		col2 []string
		col3 [][]uint8
		col4 []time.Time
	)
	for i := 0; i < 1_000; i++ {
		col1 = append(col1, uint64(i))
		col2 = append(col2, "Golang SQL database driver")
		col3 = append(col3, []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9})
		col4 = append(col4, time.Now())
	}
	if err := batch.Column(0).Append(col1); err != nil {
		return err
	}
	if err := batch.Column(1).Append(col2); err != nil {
		return err
	}
	if err := batch.Column(2).Append(col3); err != nil {
		return err
	}
	if err := batch.Column(3).Append(col4); err != nil {
		return err
	}
	return batch.Send()
}

func ScanStruct(conn clickhouse.Conn) error {
	ctx := clickhouse.Context(context.Background(), clickhouse.WithSettings(clickhouse.Settings{
		"max_block_size": 10,
	}), clickhouse.WithProgress(func(p *clickhouse.Progress) {
		fmt.Println("progress: ", p)
	}))
	if err := conn.Ping(ctx); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("Catch exception [%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		}
		return err
	}
	if err := conn.Exec(ctx, `DROP TABLE IF EXISTS example`); err != nil {
		return err
	}
	err := conn.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS example (
			Col1 UInt8,
			Col2 String,
			Col3 DateTime
		) engine=Memory
	`)
	if err != nil {
		return err
	}
	batch, err := conn.PrepareBatch(ctx, "INSERT INTO example (Col1, Col2, Col3)")
	if err != nil {
		return err
	}
	for i := 0; i < 10; i++ {
		if err := batch.Append(uint8(i), fmt.Sprintf("value_%d", i), time.Now()); err != nil {
			return err
		}
	}
	if err := batch.Send(); err != nil {
		return err
	}

	var result []struct {
		Col1           uint8
		Col2           string
		ColumnWithName time.Time `ch:"Col3"`
	}

	if err = conn.Select(ctx, &result, "SELECT Col1, Col2, Col3 FROM example"); err != nil {
		return err
	}

	for _, v := range result {
		fmt.Printf("row: col1=%d, col2=%s, col3=%s\n", v.Col1, v.Col2, v.ColumnWithName)
	}

	return nil
}

func Batch(conn clickhouse.Conn) error {
	ctx := context.Background()
	if err := conn.Exec(ctx, `DROP TABLE IF EXISTS example`); err != nil {
		return err
	}
	err := conn.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS example (
			  Col1 UInt8
			, Col2 String
			, Col3 FixedString(3)
			, Col4 UUID
			, Col5 Map(String, UInt8)
			, Col6 Array(String)
			, Col7 Tuple(String, UInt8, Array(Map(String, String)))
			, Col8 DateTime
		) Engine = Memory
	`)
	if err != nil {
		return err
	}

	batch, err := conn.PrepareBatch(ctx, "INSERT INTO example")
	if err != nil {
		return err
	}
	for i := 0; i < 500_000; i++ {
		err := batch.Append(
			uint8(42),
			"ClickHouse", "Inc",
			uuid.New(),
			map[string]uint8{"key": 1},             // Map(String, UInt8)
			[]string{"Q", "W", "E", "R", "T", "Y"}, // Array(String)
			[]interface{}{ // Tuple(String, UInt8, Array(Map(String, String)))
				"String Value", uint8(5), []map[string]string{
					map[string]string{"key": "value"},
					map[string]string{"key": "value"},
					map[string]string{"key": "value"},
				},
			},
			time.Now(),
		)
		if err != nil {
			return err
		}
	}
	return batch.Send()
}

func BatchStruct(conn clickhouse.Conn) error {
	ctx := context.Background()
	const ddl = `
CREATE TABLE example (
	  Col1 UInt64
	, Col2 String
	, Col3 Array(UInt8)
	, Col4 DateTime
) Engine = Memory
`

	type row struct {
		Col1 uint64
		Col4 time.Time
		Col2 string
		Col3 []uint8
	}
	if err := conn.Exec(ctx, `DROP TABLE IF EXISTS example`); err != nil {
		log.Fatal(err)
	}
	if err := conn.Exec(ctx, ddl); err != nil {
		log.Fatal(err)
	}
	batch, err := conn.PrepareBatch(context.Background(), "INSERT INTO example")
	if err != nil {
		return err
	}
	for i := 0; i < 1_000; i++ {
		err := batch.AppendStruct(&row{
			Col1: uint64(i),
			Col2: "Golang SQL database driver",
			Col3: []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9},
			Col4: time.Now(),
		})
		if err != nil {
			return err
		}
	}
	return batch.Send()
}

func Columer(conn clickhouse.Conn) error {
	const ddl = `
CREATE TABLE example (
	  Col1 UInt64
	, Col2 String
	, Col3 Array(UInt8)
	, Col4 DateTime
) ENGINE = Memory`

	ctx := context.Background()
	if err := conn.Exec(ctx, `DROP TABLE IF EXISTS example`); err != nil {
		log.Fatal(err)
	}
	if err := conn.Exec(ctx, ddl); err != nil {
		log.Fatal(err)
	}

	batch, err := conn.PrepareBatch(context.Background(), "INSERT INTO example")
	if err != nil {
		return err
	}
	var (
		col1 []uint64
		col2 []string
		col3 [][]uint8
		col4 []time.Time
	)
	for i := 0; i < 1_000; i++ {
		col1 = append(col1, uint64(i))
		col2 = append(col2, "Golang SQL database driver")
		col3 = append(col3, []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9})
		col4 = append(col4, time.Now())
	}
	if err := batch.Column(0).Append(col1); err != nil {
		return err
	}
	if err := batch.Column(1).Append(col2); err != nil {
		return err
	}
	if err := batch.Column(2).Append(col3); err != nil {
		return err
	}
	if err := batch.Column(3).Append(col4); err != nil {
		return err
	}
	return batch.Send()
}

func Bind(conn clickhouse.Conn) error {
	ctx := context.Background()
	if err := conn.Exec(ctx, `DROP TABLE IF EXISTS example`); err != nil {
		return err
	}
	const ddl = `
	CREATE TABLE example (
		  Col1 UInt8
		, Col2 String
		, Col3 DateTime
	) ENGINE = Memory
	`
	if err := conn.Exec(ctx, ddl); err != nil {
		return err
	}
	datetime := time.Now()
	{
		batch, err := conn.PrepareBatch(ctx, "INSERT INTO example")
		if err != nil {
			return err
		}
		for i := 0; i < 10; i++ {
			if err := batch.Append(uint8(i), "ClickHouse Inc.", datetime); err != nil {
				return err
			}
		}
		if err := batch.Send(); err != nil {
			return err
		}
	}

	var result struct {
		Col1 uint8
		Col2 string
		Col3 time.Time
	}
	{
		if err := conn.QueryRow(ctx, `SELECT * FROM example WHERE Col1 = $1 AND Col3 = $2`, 2, datetime).ScanStruct(&result); err != nil {
			return err
		}
		fmt.Println(result)
	}
	{
		if err := conn.QueryRow(ctx, `SELECT * FROM example WHERE Col1 = @Col1 AND Col3 = @Col2`,
			clickhouse.Named("Col1", 4),
			clickhouse.Named("Col2", datetime),
		).ScanStruct(&result); err != nil {
			return err
		}
		fmt.Println(result)
	}
	return nil
}
