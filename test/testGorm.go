package main

import (
	"fmt"
	models "gochat/test2"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:zimablue@tcp(127.0.0.1:3306)/gochat?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("Workkkkkkkkk")
	}
	fmt.Println("??????")
	db.AutoMigrate(&models.Community{})
	// user := &models.UserBasic{}
	// user.Name = "Mingyuan"
	// db.Create(user)
	// fmt.Println(db.First(user, 1))
	// db.Model(user).Update("PassWord", "1234")
}
