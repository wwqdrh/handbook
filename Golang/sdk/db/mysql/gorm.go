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

func InitGorm() {
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

////////////////////
// migrate
////////////////////

// 根据结构体新建数据表
func InitTable() {
	mysqlDB.AutoMigrate(&User{})
	// db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{}) // 设置引擎
}

// 删除数据表
func DeleteTable() {
	mysqlDB.Migrator().DropTable(&User{})
}

////////////////////
// curd
////////////////////

func CreateCord() {
	mysqlDB.Create(
		&User{
			Name: "张三",
		},
	)
}

func SelectCord() *User {
	var user User
	// mysqlDB.Where("name = ?", name).First(&user)
	mysqlDB.First(&user)
	return &user
}

func UpdateCord() {
	// .Save(data)  data.Update
	var user User
	mysqlDB.Where("name = ?", "张三").First(&user)
	user.Name = "张四"
	mysqlDB.Save(user)

	mysqlDB.Model(&User{}).Where("name = ?", "张四").Update("name", "张五")
}

func DeleteCord() {
	var user User
	mysqlDB.Where("name = ?", "张五").Delete(&user)
}
