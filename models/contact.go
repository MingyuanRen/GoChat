package models

import (
	"gochat/utils"

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

func SearchFriend(userId uint) []UserBasic {
	contacts := make([]Contact, 0)
	objIds := make([]uint64, 0)
	utils.DB.Where("owner_id = ? and type=1", userId).Find(&contacts)
	for _, v := range contacts {
		objIds = append(objIds, uint64(v.TargetId))
	}
	users := make([]UserBasic, 0)
	utils.DB.Where("id in ?", objIds).Find(&users)
	return users
}

// add friend 1. own id 2. friend id
func AddFriend(userId uint, targetName string) (int, string) {
	//user := UserBasic{}

	if targetName != "" {
		targetUser := FindUserByName(targetName)
		//fmt.Println(targetUser, " userId        ", )
		if targetUser.Salt != "" {
			if targetUser.ID == userId {
				return -1, "Bro, you can not add yourself"
			}
			contact0 := Contact{}
			utils.DB.Where("owner_id =?  and target_id =? and type=1", userId, targetUser.ID).Find(&contact0)
			if contact0.ID != 0 {
				return -1, "Can not add duplicate friends"
			}
			tx := utils.DB.Begin()
			// rollback
			defer func() {
				if r := recover(); r != nil {
					tx.Rollback()
				}
			}()
			contact := Contact{}
			contact.OwnerId = userId
			contact.TargetId = targetUser.ID
			contact.Type = 1
			if err := utils.DB.Create(&contact).Error; err != nil {
				tx.Rollback()
				return -1, "Fail to add friend"
			}
			contact1 := Contact{}
			contact1.OwnerId = targetUser.ID
			contact1.TargetId = userId
			contact1.Type = 1
			if err := utils.DB.Create(&contact1).Error; err != nil {
				tx.Rollback()
				return -1, "Fail to add friend"
			}
			tx.Commit()
			return 0, "Successfully added friend"
		}
		return -1, "User Not Found"
	}
	return -1, "Friend ID can not be empty"
}

func SearchUserByGroupId(communityId uint) []uint {
	contacts := make([]Contact, 0)
	objIds := make([]uint, 0)
	utils.DB.Where("target_id = ? and type=2", communityId).Find(&contacts)
	for _, v := range contacts {
		objIds = append(objIds, uint(v.OwnerId))
	}
	return objIds
}
