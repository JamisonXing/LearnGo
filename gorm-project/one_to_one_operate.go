package main

import (
	"gorm-project/models/relate_tables"
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
	db.AutoMigrate(&relate_tables.UserProfile{}, &relate_tables.User{})

	//增加
	/*	userProfile := relate_tables.UserProfile{
			Pic:   "1.jpg",
			CPic:  "2.jpg",
			Phone: "150xxx",
			User: relate_tables.User{
				Id:   0,
				Name: "jamison",
				Age:  20,
				Addr: "xc",
			},
		}

		db.Create(&userProfile)*/

	/*	//关联查询
		//方式一
		var userProfile relate_tables.UserProfile
		db.Debug().First(&userProfile, 1)
		db.Debug().Model(&userProfile).Association("User").Find(&userProfile.User)
		fmt.Println(userProfile)

		//方式二
		var userProfile2 relate_tables.UserProfile
		db.Preload("User").Find(&userProfile2, 2)
		fmt.Println(userProfile2)*/

	//更新
	var userProfile3 relate_tables.UserProfile
	//先查询
	db.Preload("User").Find(&userProfile3, 1)

	//db.Model(&userProfile3.User).Update("name", "xxj") //单字段
	db.Model(&userProfile3.User).Updates(relate_tables.User{ //多字段
		Name: "lll",
		Age:  45,
		Addr: "gg",
	})

	//删除
	var userProfile4 relate_tables.UserProfile
	db.Preload("User").Find(&userProfile4, 2)
	db.Debug().Delete(&userProfile4.User)

}
