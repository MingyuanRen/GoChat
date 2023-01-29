package service

import (
	"fmt"
	"gochat/models"
	"gochat/utils"
	"math/rand"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// GetUserList
// @Summary list all users
// @Tags userpage
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

	//
	salt := fmt.Sprintf("%06d", rand.Int31())
	data := models.FindUserByName(user.Name)
	if data.Name != "" {
		c.JSON(-1, gin.H{
			"message": "Try again! The user has been registered",
		})
		return
	}
	if password != repassword {
		c.JSON(-1, gin.H{
			"message": "Invalid password",
		})
		return
	}

	// user.PassWord = password
	user.PassWord = utils.MakePassword(password, salt)
	user.Salt = salt
	models.CreateUser(user)
	c.JSON(200, gin.H{
		"message": "New User has been Created!",
	})
}

// DeleteUser
// @Summary delete user
// @Tags userpage
// @param id query string false "id"
// @Success 200 {string} json{"code", "message"}
// @Router /user/DeleteUser [get]
func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)
	models.DeleteUser(user)
	c.JSON(200, gin.H{
		"message": "Delete User!",
	})
}

// UpdateUser
// @Summary update user
// @Tags userpage
// @param id formData string false "id"
// @param name formData string false "name"
// @param password formData string false "password"
// @param phone formData string false "phone"
// @param email formData string false "email"
// @Success 200 {string} json{"code", "message"}
// @Router /user/UpdateUser [post]
func UpdateUser(c *gin.Context) {
	// fmt.Println("call!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	user.Name = c.PostForm("name")
	user.PassWord = c.PostForm("password")
	user.Phone = c.PostForm("phone")
	user.Email = c.PostForm("email")
	fmt.Println("update :", user)

	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"message": "Not valid Update for user INFO!",
		})
	} else {
		models.UpdateUser(user)
		c.JSON(200, gin.H{
			"message": "Updated User Successfully!",
		})
	}

}

// FindUserByNameAndPwd
// @Summary login
// @Tags userpage
// @param name query string false "name"
// @param password query string false "pasword"
// @Success 200 {string} json{"code","message"}
// @Router /user/FindUserByNameAndPwd [post]
func FindUserByNameAndPwd(c *gin.Context) {
	data := models.UserBasic{}

	name := c.Query("name")
	password := c.Query("password")
	user := models.FindUserByName(name)

	if user.Name == "" {
		c.JSON(200, gin.H{
			"message": "This user doesn't exist!",
		})
		return
	}
	flag := utils.ValidPassword(password, user.Salt, user.PassWord)
	if !flag {
		c.JSON(200, gin.H{
			"message": "Wrong Password",
		})
		return
	}
	pwd := utils.MakePassword(password, user.Salt)
	data = models.FindUserByNameAndPwd(name, pwd)
	c.JSON(200, gin.H{
		"message": data,
	})
}
