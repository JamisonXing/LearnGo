package main

import (
	"fmt"
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
	db.AutoMigrate(&relate_tables.User2{}, &relate_tables.Article{})

	//添加
	/*	user := relate_tables.User2{
			Name: "jamison",
			Age:  30,
			Addr: "xc",
			Articles: []relate_tables.Article{
				{
					Title:   "beego详解",
					Content: "内容",
					Desc:    "详解",
				},
				{
					Title:   "beego详解2",
					Content: "内容2",
					Desc:    "详解2",
				},
			},
		}
		db.Create(&user)*/

	//查询
	var user2 relate_tables.User2
	db.Preload("Articles").Find(&user2, 1)
	fmt.Println(user2)

	//更新
	var user3 relate_tables.User2
	db.Preload("Articles").Find(&user3, 1)
	fmt.Println(user3)
	//db.Model(&user3.Articles[0]).Update("title", "beego详解xxx")
	db.Model(&user3.Articles).Where("id = ?", 1).Update("title", "beego详解1")

	//删除
	var user4 relate_tables.User2
	db.Preload("Articles").Find(&user4, 1)
	fmt.Println(user3)
	db.Where("id = ?", 1).Delete(&user4.Articles)

}
