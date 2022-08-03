package postgresql

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getDB(t *testing.T) *gorm.DB {
	dsn := "host=localhost user=postgres password=hui123456 dbname=dbtest port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 必须添加这个 在migrate的时候才不会创建外键
	})
	require.Nil(t, err)
	return db
}

func TestRelation(t *testing.T) {
	db := getDB(t)

	db.AutoMigrate(&Db1{})
	db.AutoMigrate(&Db2{})

	var db1column = &Db1{
		Name: "123",
	}
	require.Nil(t, db.Create(db1column).Error)

	var db2column = []*Db2{
		{Db1Id: db1column.ID, Name: "第一条"},
		{Db1Id: db1column.ID, Name: "第二条"},
		{Db1Id: db1column.ID, Name: "第三条"},
		{Db1Id: db1column.ID, Name: "第四条"},
		{Db1Id: db1column.ID, Name: "第五条"},
	}
	require.Nil(t, db.Create(&db2column).Error)

	var info Db1
	db.Model(Db1{}).Where("id=1").Preload("Db2Info").First(&info) // 获取级联关系
	fmt.Println(info.Db2Info)
}

func TestExpandMenu(t *testing.T) {
	db := getDB(t)
	defer func() {
		if err := db.Exec("drop table menu").Error; err != nil {
			require.Nil(t, err)
		}
	}()
	db.AutoMigrate(&Menu{})

	menu := []Menu{
		{Name: "1", ParentId: 0},
		{Name: "2", ParentId: 0},
		{Name: "3", ParentId: 0},
		{Name: "4", ParentId: 0},
		{Name: "1-1", ParentId: 1},
		{Name: "1-2", ParentId: 1},
		{Name: "1-3", ParentId: 1},
		{Name: "1-4", ParentId: 1},
		{Name: "1-1-1", ParentId: 5},
		{Name: "2-1", ParentId: 2},
	}
	require.Nil(t, db.Create(&menu).Error)

	var menus []Menu
	if err := db.Model(Menu{}).Where("id = ?", 6).Preload("ParentMenu", expandMenu).Find(&menus).Error; err != nil {
		require.Nil(t, err)
	}
	fmt.Println(menus)

	var menus2 []Menu
	if err := db.Model(Menu{}).Where("id = ?", 1).Preload("Children", expandChildren).Find(&menus2).Error; err != nil {
		require.Nil(t, err)
	}
	fmt.Println(menus2)
}
