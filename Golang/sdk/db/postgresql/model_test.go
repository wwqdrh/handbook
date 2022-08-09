package postgresql

import (
	"fmt"
	"testing"

	gorm1 "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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

func TestDB3Menu(t *testing.T) {
	db := getDB(t)
	defer func() {
		if err := db.Exec("drop table db3").Error; err != nil {
			require.Nil(t, err)
		}
	}()
	db.AutoMigrate(&Db3{})
	menu := []Db3{
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
	var menus2 []Db3
	if err := db.Model(Db3{}).Where("id = ?", 1).Preload("Db3List", Db3{}.expandChild1).Find(&menus2).Error; err != nil {
		require.Nil(t, err)
	}
	fmt.Println(menus2)
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

type TeUser struct {
	ID   int64  `gorm:"primaryKey"`
	Name string `gorm:"name"`
}

func (TeUser) TableName() string {
	return "te_user"
}

func TestMulti(t *testing.T) {
	db := getDB(t)
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		if err := db.Exec("drop table te_user").Error; err != nil {
			require.Nil(t, err)
		}
	}()
	db.AutoMigrate(TeUser{})

	users := []TeUser{
		{Name: "user1"},
		{Name: "user2"},
	}
	for _, item := range users {
		// 如果使用 item则 reflect.Value.SetInt using unaddressable value
		fmt.Println(db.Model(item).Create(&item).Error)
	}
}

func TestMulti2(t *testing.T) {
	db, err := gorm1.Open("postgres", "host=localhost port=5432 user=postgres dbname=dbtest password=hui123456 sslmode=disable TimeZone=Asia/Shanghai")
	require.Nil(t, err)
	defer db.Close()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		if err := db.Exec("drop table te_user").Error; err != nil {
			require.Nil(t, err)
		}
	}()
	db.AutoMigrate(TeUser{})

	users := []TeUser{
		{Name: "user1"},
		{Name: "user2"},
	}
	for _, item := range users {
		// 如果使用 item则 reflect.Value.SetInt using unaddressable value
		fmt.Println(db.Model(item).Create(&item).Error)
	}
}

func TestUniqueIndexModify(t *testing.T) {
	db := getDB(t)
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		if err := db.Exec("drop table search_rank").Error; err != nil {
			require.Nil(t, err)
		}
	}()

	db.AutoMigrate(&SearchRank{})
	// 太坑了，必须要先手动dropindex才能重建索引
	require.Nil(t, db.Migrator().DropIndex(&SearchRank{}, "search_menu_date"))
	db.AutoMigrate(&SearchRank2{})
	fmt.Println("here")
}
