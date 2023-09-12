package postgresql

import (
	"time"

	"gorm.io/gorm"
)

type Db1 struct {
	gorm.Model
	Name    string `json:"name"`
	Db2Info []Db2  `json:"db_2_info"`
}

type Db2 struct {
	gorm.Model
	Name    string `json:"name"`
	Db1Id   uint   `json:"db_1_id"`
	Db1Info *Db1   `gorm:"foreignKey:Db1Id" json:"db_1_info"`
}

type Db3 struct {
	// ID         uint32 `json:"id" form:"id"`
	ID       int64  `gorm:"primarykey"`
	Name     string `json:"name"`
	ParentId uint   `gorm:"column:parent_id;default:0;" json:"parent_id"` //父ID
	Db3Info  *Db3   `gorm:"foreignKey:ParentId"  json:"db_3_info,omitempty"`
	Db3List  []*Db3 `gorm:"foreignKey:ParentId;"  json:"db_3_list,omitempty"`
}

type SearchRank struct {
	ID              int64     `gorm:"primaryKey"`
	MenuID          int64     `gorm:"column:menu_id;uniqueIndex:search_menu_date;not null"`
	Date            time.Time `gorm:"type:date;uniqueIndex:search_menu_date;default:current_date;not null"`
	SearchTerm      string    `gorm:"column:search_term;uniqueIndex:search_menu_date;not null"`
	Asin            string    `gorm:"column:asin;not null"`
	Rank            int64     `gorm:"column:rank;not null"`
	ClickShare      string    `gorm:"column:click_share"`
	ConversionShare string    `gorm:"column:conversion_shares"`
}

type SearchRank2 struct {
	ID              int64     `gorm:"primaryKey"`
	MenuID          int64     `gorm:"column:menu_id;uniqueIndex:search_menu_date;not null"`
	Date            time.Time `gorm:"type:date;uniqueIndex:search_menu_date;default:current_date;not null"`
	SearchTerm      string    `gorm:"column:search_term;uniqueIndex:search_menu_date;not null"`
	Asin            string    `gorm:"column:asin;uniqueIndex:search_menu_date;not null"`
	Rank            int64     `gorm:"column:rank;not null"`
	ClickShare      string    `gorm:"column:click_share"`
	ConversionShare string    `gorm:"column:conversion_shares"`
}

func (SearchRank) TableName() string {
	return "search_rank"
}

func (SearchRank2) TableName() string {
	return "search_rank"
}

func (Db3) TableName() string {
	return "db3"
}

func (d Db3) expandChild1(tx *gorm.DB) *gorm.DB {
	return tx.Preload("Db3List", d.expandChild1)
}

type Menu struct {
	ID       int64  `gorm:"primarykey"`
	Name     string `gorm:"uniqueIndex" json:"name"`
	ParentId int64  `gorm:"column:parentid"`

	ParentMenu *Menu  `gorm:"foreignKey:ParentId"`
	Children   []Menu `gorm:"foreignKey:ParentId"`
}

func (Menu) TableName() string {
	return "menu"
}

// 巨坑 不能这么写 留着警示
func (Menu) expandMenu(db *gorm.DB) *gorm.DB {
	var fn func() *gorm.DB
	fn = func() *gorm.DB {
		return db.Preload("ParentMenu", fn)
	}
	return fn()
}

// 巨坑 不能这么写 留着警示
func (Menu) expandChildren(db *gorm.DB) *gorm.DB {
	var fn func() *gorm.DB
	fn = func() *gorm.DB {
		return db.Preload("Children", fn)
	}
	return fn()
}

func expandMenu(db *gorm.DB) *gorm.DB {
	return db.Preload("ParentMenu", expandMenu)
}

func expandChildren(db *gorm.DB) *gorm.DB {
	return db.Preload("Children", expandChildren)
}
