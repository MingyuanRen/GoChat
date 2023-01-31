package models

import "gorm.io/gorm"

// Group info
type GroupBasic struct {
	gorm.Model
	Name    string
	OwnerId uint
	Icon    string
	Type    int
	Desc    string
}

func (table *GroupBasic) TableName() string {
	return "group_basic"
}
