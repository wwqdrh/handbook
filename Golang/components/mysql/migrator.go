package mysql

// 根据结构体新建数据表
func InitTable() {
	mysqlDB.AutoMigrate(&User{})
	// db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{}) // 设置引擎
}

// 删除数据表
func DeleteTable() {
	mysqlDB.Migrator().DropTable(&User{})
}
