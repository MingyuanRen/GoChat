package models

import (
	"gorm.io/gorm"
)

// contact info
type Contact struct {
	gorm.Model
	OwnerId  uint //who's relationship information
	TargetId uint //The corresponding person/group ID
	Type     int  //corresponding type 1 friend 2 group 3xx
	Desc     string
}

func (table *Contact) TableName() string {
	return "contact"
}
