package mysql

import "testing"

func TestInitTable(t *testing.T) {
	InitTable()
	if !mysqlDB.Migrator().HasTable(&User{}) {
		t.Error("发生错误")
	}
	DeleteTable()
	if mysqlDB.Migrator().HasTable(&User{}) {
		t.Error("发生错误")
	}
}

func TestCURD(t *testing.T) {
	InitTable()
	if !mysqlDB.Migrator().HasTable(&User{}) {
		t.Error("发生错误")
	}

	CreateCord()
	if SelectCord().Name != "张三" {
		t.Error("发生错误")
	}
	UpdateCord()
	if SelectCord().Name != "张五" {
		t.Error("发生错误")
	}
	DeleteCord()

	DeleteTable()
	if mysqlDB.Migrator().HasTable(&User{}) {
		t.Error("发生错误")
	}
}
