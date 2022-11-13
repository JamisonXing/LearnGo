package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type User struct {
	Name string
	Age  int
	Addr string
	Pic  string //头像地址
	gorm.Model
}

// OpenDB 连接数据库
func OpenDB() *gorm.DB {
	//username:password@tcp(ip:port)/dbname?charset=utf8&parseTime=True&loc=Local
	dsn := "root:xxj753896@tcp(localhost:3306)/gorm_project?charset=utf8mb3&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("connect success: %v\n", db)
	return db
}

// CreateTable 创建表
func CreateTable(db *gorm.DB, Any any) {
	err := db.AutoMigrate(Any) //any是模型
	if err != nil {
		fmt.Printf("fail to create table:%v\n", err)
	} else {
		fmt.Printf("create table success:%v\n", Any)
	}
}

// InsertInfo 创建记录
func InsertInfo(db *gorm.DB, Any any) {
	//db := OpenDB()
	create := db.Create(Any)
	if create.Error != nil {
		fmt.Printf("fail insert:%v\n", create.Error)
	} else {
		fmt.Printf("insert success:%v\n", Any)
	}
}

// QueryByPrimaryKey 简单记录查询，通过主键查找
func QueryByPrimaryKey(db *gorm.DB, Any any, Id int) {
	first := db.Find(Any, Id)
	if first.Error != nil {
		fmt.Printf("fail query:%v\n", first.Error)
	} else {
		fmt.Printf("query success:%v\n", Any)
	}
}

// QueryByName 通过姓名查找
func QueryByName(db *gorm.DB, Any any, name string) {
	first := db.Find(Any, "name=?", name)
	if first.Error != nil {
		fmt.Printf("fail query:%v\n", first.Error)
	} else {
		fmt.Printf("query success:%v\n", Any)
	}
}

// UpdateInfo 修改记录
func UpdateInfo(db *gorm.DB, Any *User, Id int, Name string) {
	first := db.Find(Any, Id)
	if first.Error != nil {
		fmt.Printf("search fail:%v\n", first.Error)
	} else {
		Any.Name = Name
		save := db.Save(Any)
		if save.Error != nil {
			fmt.Printf("modify success:after value:%v\n", Any)
		}
	}
}

// DeleteInfo 删除记录
func DeleteInfo(db *gorm.DB, Any any, Id int) {
	first := db.Find(Any, Id)
	if first.Error != nil {
		fmt.Printf("search fail:%v\n", first.Error)
	} else {
		del := db.Delete(Any, Id)
		if del.Error != nil {
			fmt.Printf("del fail:%v\n", del.Error)
		} else {
			fmt.Printf("del success:%v\n", del.RowsAffected)
		}
	}
}

func main() {
	user := &User{
		Name:  "jamison",
		Age:   0,
		Addr:  "xc",
		Pic:   "www.baidu.com",
		Model: gorm.Model{},
	}
	db := OpenDB()
	CreateTable(db, user)
	InsertInfo(db, user)
	QueryByPrimaryKey(db, user, 2)
	UpdateInfo(db, user, 2, "xxj")
	DeleteInfo(db, user, 8)
}
