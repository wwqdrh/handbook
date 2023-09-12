package postgresql

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

// DBResolver 会根据工作表/结构自动切换连接 对于 RAW SQL，DBResolver 将从 SQL 中提取表名以匹配解析器
// Multiple sources, replicas
// Read/Write Splitting
// Automatic connection switching based on the working table/struct
// Manual connection switching
// Sources/Replicas load balancing
// Works for RAW SQL
// Transaction

type Resolver struct {
	Name string
	Dsn  string
}

type User struct {
	Name string
}

type Address struct {
}

type Product struct {
}

type AdvancedUser struct {
}

type Pet struct {
}

type Report struct {
}

func NewResolver(resolvers ...Resolver) (*gorm.DB, error) {
	DB, err := gorm.Open(postgres.Open("db1_dsn"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB.Use(dbresolver.Register(dbresolver.Config{
		// use `db2` as sources, `db3`, `db4` as replicas
		Sources:  []gorm.Dialector{postgres.Open("db2_dsn")},
		Replicas: []gorm.Dialector{postgres.Open("db3_dsn"), postgres.Open("db4_dsn")},
		// sources/replicas load balancing policy
		Policy: dbresolver.RandomPolicy{},
	}).Register(dbresolver.Config{
		// use `db1` as sources (DB's default connection), `db5` as replicas for `User`, `Address`
		Replicas: []gorm.Dialector{postgres.Open("db5_dsn")},
	}, &User{}, &Address{}).Register(dbresolver.Config{
		// use `db6`, `db7` as sources, `db8` as replicas for `orders`, `Product`
		Sources:  []gorm.Dialector{postgres.Open("db6_dsn"), postgres.Open("db7_dsn")},
		Replicas: []gorm.Dialector{postgres.Open("db8_dsn")},
	}, "orders", &Product{}, "secondary"))

	return DB, nil
}
