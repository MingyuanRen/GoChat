package models

import (
	"fmt"
	"gochat/utils"

	"gorm.io/gorm"
)

type Community struct {
	gorm.Model
	Name    string
	OwnerId uint
	Img     string
	Desc    string
}

func CreateCommunity(community Community) (int, string) {
	tx := utils.DB.Begin()
	// Once the transaction starts, no matter what exception will eventually rollback
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if len(community.Name) == 0 {
		return -1, "Group Name Can not be Empty"
	}
	if community.OwnerId == 0 {
		return -1, "Please login first"
	}
	if err := utils.DB.Create(&community).Error; err != nil {
		fmt.Println(err)
		tx.Rollback()
		return -1, "Fail to create Group"
	}
	contact := Contact{}
	contact.OwnerId = community.OwnerId
	contact.TargetId = community.ID
	contact.Type = 2 //gruop conn
	if err := utils.DB.Create(&contact).Error; err != nil {
		tx.Rollback()
		return -1, "Fail to join Group"
	}

	tx.Commit()
	return 0, "Created Group Successfully"

}

func LoadCommunity(ownerId uint) ([]*Community, string) {
	contacts := make([]Contact, 0)
	objIds := make([]uint64, 0)
	utils.DB.Where("owner_id = ? and type=2", ownerId).Find(&contacts)
	for _, v := range contacts {
		objIds = append(objIds, uint64(v.TargetId))
	}

	data := make([]*Community, 10)
	utils.DB.Where("id in ?", objIds).Find(&data)
	for _, v := range data {
		fmt.Println(v)
	}
	//utils.DB.Where()
	return data, "search successful"
}
