package mysql

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
