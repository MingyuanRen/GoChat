package service

import (
	"gochat/models"

	"github.com/gin-gonic/gin"
)

// GetUserList
// @Summary list all users
// @Tags mainpage
// @Success 200 {string} json{"code", "message"}
// @Router /user/GetUserList [get]
func GetUserList(c *gin.Context) {
	data := make([]*models.UserBasic, 10)
	data = models.GetUserList()
	c.JSON(200, gin.H{
		"message": data,
	})
}

// CreateUser
// @Summary create user
// @Tags userpage
// @param name query string false "username"
// @param password query string false "password"
// @param repassword query string false "confirm password"
// @Success 200 {string} json{"code", "message"}
// @Router /user/CreateUser [get]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Query("name")
	password := c.Query("password")
	repassword := c.Query("repassword")
	if password != repassword {
		c.JSON(-1, gin.H{
			"message": "Invalid password",
		})
	}
	user.PassWord = password
	models.CreateUser(user)
	c.JSON(200, gin.H{
		"message": "New User Created!",
	})
}
