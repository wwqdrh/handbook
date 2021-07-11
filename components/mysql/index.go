package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	mysqlDSN = "%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local"
	mysqlDB  *gorm.DB
)

type User struct {
	ID   uint64 `gorm:"column:id;primary_key"`
	Name string `gorm:"column:name;type:varchar(255);unique"`
}

func init() {
	var err error
	mysqlDB, err = gorm.Open(mysql.Open(fmt.Sprintf(mysqlDSN,
		"root",
		"123456",
		"127.0.0.1:3306",
		"handbook")), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Println("mysql初始化失败")
	}
}
