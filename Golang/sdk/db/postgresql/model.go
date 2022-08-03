package postgresql

import "gorm.io/gorm"

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
