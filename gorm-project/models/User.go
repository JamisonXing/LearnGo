package models

type User struct {
	Id     int    `gorm:"primaryKey"`
	Name   string `gorm:"index:idx_name;not null"`
	Age    int
	Addr   string
	Pic    string //头像地址
	Phone  string
	Ignore string `gorm:"-"`
}
