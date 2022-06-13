package postgresql

import (
	"fmt"
	"os"
	"testing"

	"gorm.io/gorm"
	"gorm.io/sharding"
)

var db *gorm.DB

func initDB(t *testing.T) {
	dsn := os.Getenv("db")
	if dsn == "" {
		t.Skip()
	}

	db = NewSharding(dsn)
}

func TestShardingBasicInsert(t *testing.T) {
	initDB(t)

	// insert to orders_02
	err := db.Create(&Order{UserID: 2}).Error
	if err != nil {
		t.Error(err)
	}

	// insert to otders_03
	err = db.Exec("INSERT INTO orders(user_id) VALUES(?)", int64(3)).Error
	if err != nil {
		t.Error(err)
	}

	// this will redirect query to orders_02
	var orders []Order
	err = db.Model(&Order{}).Where("user_id", int64(2)).Find(&orders).Error
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%#v\n", orders)
}

func TestShardingErrMissShardingKey(t *testing.T) {
	initDB(t)

	// this will throw ErrMissingShardingKey error
	err := db.Exec("INSERT INTO orders(product_id) VALUES(1)").Error
	if err != sharding.ErrMissingShardingKey {
		t.Error(err)
	}

	// this will throw ErrMissingShardingKey error
	var orders []Order
	err = db.Model(&Order{}).Where("product_id", "1").Find(&orders).Error
	if err != sharding.ErrMissingShardingKey {
		t.Error(err)
	}
}

func TestShardingUpdate(t *testing.T) {
	initDB(t)

	// Update and Delete are similar to create and query
	err := db.Exec("UPDATE orders SET product_id = ? WHERE user_id = ?", 2, int64(3)).Error
	if err != nil {
		t.Error()
	}
	err = db.Exec("DELETE FROM orders WHERE product_id = 3").Error
	if err != sharding.ErrMissingShardingKey {
		t.Error(err)
	}
}
