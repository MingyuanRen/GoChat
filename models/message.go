package models

import (
	"gorm.io/gorm"
)

// 消息
type Message struct {
	gorm.Model
	UserId     int64  //sender
	TargetId   int64  //Receiver
	Type       int    //Send type 1 private chat 2 group chat 3 heartbeat
	Media      int    //Message type 1 text 2 sticker 3 voice 4 picture/meme
	Content    string //message content
	CreateTime uint64 //create time
	ReadTime   uint64 //read time
	pic        string
	Url        string
	Desc       string
	Amount     int // other digital statistics
}

func (table *Message) TableName() string {
	return "message"
}
