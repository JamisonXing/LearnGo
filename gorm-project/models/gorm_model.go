package models

import "gorm.io/gorm"

type GormModel struct {
	gorm.Model
	Name string
}

// TableName 重命名表名
func (GormModel) TableName() string {
	return "test_gorm_model"
}
