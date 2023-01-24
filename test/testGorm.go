package main

import (
	"fmt"
	"gochat/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:zimablue@tcp(127.0.0.1:3306)/gochat?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// schema
	db.AutoMigrate(&models.UserBasic{})

	// Create
	user := &models.UserBasic{}
	user.Name = "Mingyuan"
	db.Create(user)

	// Read
	fmt.Println(db.First(user, 1))

	// Update
	db.Model(user).Update("PassWord", "1234")
}
