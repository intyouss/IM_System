package main

import (
	"IM_System/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	// 迁移 schema
	DB, err := gorm.Open(mysql.Open("root:xxxxx@tcp(127.0.0.1:3306)/ginchat?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	DB.AutoMigrate(&models.UserContact{})
	DB.AutoMigrate(&models.UserMessage{})
	DB.AutoMigrate(&models.UserGroup{})
	//if err != nil {
	//	panic("failed to migrate database")
	//}
	//
	//// Create
	//user := &mysql.UserBasic{}
	//user.Name = "haha"
	//mysql.Create(user)

	// Read
	//fmt.Println(mysql.First(user, 1)) // 根据整型主键查找
	//mysql.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	// Update - 将 product 的 price 更新为 200
	//mysql.Model(user).Update("Password", "12345")
	// Update - 更新多个字段
	//mysql.Model(user).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	//mysql.Model(user).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - 删除 product
	//mysql.Delete(&product, 1)
}
