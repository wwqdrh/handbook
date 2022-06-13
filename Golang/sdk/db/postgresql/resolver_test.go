package postgresql

import (
	"fmt"
	"testing"

	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

var (
	// 测试环境配置 user、address、product三个垂直分库的数据库
	userdsn    = "postgres://postgres:hui123456@localhost:5432/sharding-user?sslmode=disable"
	addressdsn = "postgres://postgres:hui123456@localhost:5432/sharding-address?sslmode=disable"
	productdsn = "postgres://postgres:hui123456@localhost:5432/sharding-product?sslmode=disable"
)

func initResolverDB(t *testing.T) *gorm.DB {
	db, err := NewResolver()
	if err != nil {
		t.Fatal(err)
	}
	return db
}

func TestResolverSQL(t *testing.T) {
	DB := initResolverDB(t)
	// `User` Resolver Examples
	DB.Table("users").Rows()                       // replicas `db5`
	DB.Model(&User{}).Find(&AdvancedUser{})        // replicas `db5`
	DB.Exec("update users set name = ?", "jinzhu") // sources `db1`

	var name string
	DB.Raw("select name from users").Row().Scan(&name) // replicas `db5`
	DB.Create(&User{Name: "jinzhu"})                   // sources `db1`
	DB.Delete(&User{}, "name = ?", "jinzhu")           // sources `db1`
	DB.Table("users").Update("name", "jinzhu")         // sources `db1`

	// Global Resolver Examples
	DB.Find(&Pet{}) // replicas `db3`/`db4`
	DB.Save(&Pet{}) // sources `db2`

	// Orders Resolver Examples
	DB.Find(&Order{})                  // replicas `db8`
	DB.Table("orders").Find(&Report{}) // replicas `db8`
}

// 手动切换连接数据库
func TestResolverSwitch(t *testing.T) {
	DB := initResolverDB(t)

	var user User
	// Use Write Mode: read user from sources `db1`
	DB.Clauses(dbresolver.Write).First(&user)

	// Specify Resolver: read user from `secondary`'s replicas: db8
	DB.Clauses(dbresolver.Use("secondary")).First(&user)

	// Specify Resolver and Write Mode: read user from `secondary`'s sources: db6 or db7
	DB.Clauses(dbresolver.Use("secondary"), dbresolver.Write).First(&user)
}

// 管理事务
func TestResolverTransaction(t *testing.T) {
	DB := initResolverDB(t)

	// Start transaction based on default replicas db
	tx := DB.Clauses(dbresolver.Read).Begin()

	// Start transaction based on default sources db
	tx = DB.Clauses(dbresolver.Write).Begin()

	// Start transaction based on `secondary`'s sources
	tx = DB.Clauses(dbresolver.Use("secondary"), dbresolver.Write).Begin()

	fmt.Println(tx)
}
