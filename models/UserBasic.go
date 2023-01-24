package models

import "gorm.io/gorm"

type UserBasic struct {
	gorm.Model
	Name          string
	Password      string
	Phone         string
	Email         string
	Identifier    string
	ClientIP      string
	ClientPort    string
	LoginTime     uint64
	HeartBeatTime uint64
	LogoutTime    uint64
	IsLogout      bool
	DeviceInfo    string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}
