package main

func main() {

	// 迁移 schema
	//err := utils.DB.AutoMigrate(&db.UserBasic{})
	//if err != nil {
	//	panic("failed to migrate database")
	//}
	//
	//// Create
	//user := &db.UserBasic{}
	//user.Name = "haha"
	//db.Create(user)

	// Read
	//fmt.Println(db.First(user, 1)) // 根据整型主键查找
	//db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	// Update - 将 product 的 price 更新为 200
	//db.Model(user).Update("Password", "12345")
	// Update - 更新多个字段
	//db.Model(user).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	//db.Model(user).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - 删除 product
	//db.Delete(&product, 1)
}
