package main

import (
	"fmt"
	"gorm-project/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	//username:password@tcp(ip:port)/dbname?charset=utf8&parseTime=True&loc=Local
	dsn := "root:xxj753896@tcp(localhost:3306)/gorm_project?charset=utf8mb3&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	//自动迁移
	db.AutoMigrate(&models.User{})

	//增加
	db.Create(&models.User{
		Id:    2,
		Name:  "zhangSan",
		Age:   18,
		Addr:  "xxx",
		Pic:   "/static/upload/pic111.jpg",
		Phone: "150xxx",
	})

	//查询
	var user models.User
	//db.First(&user, 2) // 根据整型主键查找
	db.First(&user, "name=?", "zhangSan")
	fmt.Println(user)

	//db.Migrator().DropTable("users")

}
