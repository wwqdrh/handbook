package main

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

var DB *gorm.DB

type User struct {
	ID   int    `gorm:"primaryKey;not null"`
	Name string `gorm:"not null"`
}

func init() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3307)/user?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	db.Use(dbresolver.Register(
		dbresolver.Config{
			Replicas: []gorm.Dialector{mysql.Open("root:123456@tcp(127.0.0.1:3308)/user?charset=utf8mb4&parseTime=True&loc=Local"), mysql.Open("root:123456@tcp(127.0.0.1:3309)/user?charset=utf8mb4&parseTime=True&loc=Local")},
			Policy:   dbresolver.RandomPolicy{},
		},
		"secondary",
	))
	if err := db.AutoMigrate(User{}); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	DB = db
}

func main() {
	user := User{
		Name: "张三",
	}
	DB.Create(&user)
	fmt.Println(user.ID, user.Name)
	fmt.Println("wait sync")
	time.Sleep(3 * time.Second)
	fmt.Println("test sync")

	var secuser User
	if err := DB.Clauses(dbresolver.Use("secondary")).Model(User{}).Where("id = ?", user.ID).Find(&secuser).Error; err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(secuser.ID, secuser.Name)
	}
}
