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
	db.AutoMigrate(&relate_tables.Tag{}, &relate_tables.Article2{})

	//增加
	/*	tag1 := relate_tables.Tag{
			Name: "tag1",
			Desc: "描1",
		}
		tag2 := relate_tables.Tag{
			Name: "tag2",
			Desc: "描2",
		}
		//模拟tag database已经存在
		db.Create(&tag1)
		db.Create(&tag2)

		article := relate_tables.Article2{
			Title:   "文章标题",
			Content: "文章内容",
			Desc:    "描述",
			Tags: []relate_tables.Tag{
				tag1,
				tag2,
			},
		}
		db.Create(&article)*/

	//查询
	var article2 relate_tables.Article2
	db.Preload("Tags").Find(&article2, 1)
	fmt.Println(article2)

	/*	//更新
		var article3 relate_tables.Article2
		db.Preload("Tags").Find(&article3, 1)
		fmt.Println(article3)
		db.Model(&article3.Tags).Where("id=?", 1).Update("name", "gin标签")
	*/
	//删除
	var article4 relate_tables.Article2
	db.Preload("Tags").Find(&article4, 5)
	fmt.Println(article4)
	db.Where("id = ?", 5).Delete(&article4.Tags)
}
