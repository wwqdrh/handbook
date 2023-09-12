package postgresql

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/sharding"
)

type Order struct {
	ID        int64 `gorm:"primaryKey"`
	UserID    int64
	ProductID int64
}

func NewSharding(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn}))
	if err != nil {
		panic(err)
	}

	// 初始化分表
	for i := 0; i < 64; i += 1 {
		table := fmt.Sprintf("orders_%02d", i)
		// db.Exec(`DROP TABLE IF EXISTS ` + table)
		db.Exec(`CREATE TABLE IF NOT EXISTS ` + table + ` (
			id BIGSERIAL PRIMARY KEY,
			user_id bigint,
			product_id bigint
		)`)
	}

	// 初始阿虎sharding
	middleware := sharding.Register(sharding.Config{
		ShardingKey:         "user_id",
		NumberOfShards:      64,
		PrimaryKeyGenerator: sharding.PKSnowflake,
	}, "orders")
	db.Use(middleware)
	return db
}
