package service

import (
	"fmt"
	"gochat/models"
	"gochat/utils"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// GetUserList
// @Summary list all users
// @Tags userpage
// @Success 200 {string} json{"code", "message"}
// @Router /user/getUserList [post]
func GetUserList(c *gin.Context) {
	data := make([]*models.UserBasic, 10)
	data = models.GetUserList()
	c.JSON(200, gin.H{
		"code":    0, //  0 success -1 fail
		"message": "User List",
		"data":    data,
	})
}

// CreateUser
// @Summary create user
// @Tags userpage
// @param name query string false "username"
// @param password query string false "password"
// @param repassword query string false "confirm password"
// @Success 200 {string} json{"code", "message"}
// @Router /user/createUser [post]
func CreateUser(c *gin.Context) {
	fmt.Println("call create user")

	user := models.UserBasic{}
	// user.Name = c.Query("name")
	// password := c.Query("password")
	// repassword := c.Query("repassword")
	user.Name = c.Request.FormValue("name")
	password := c.Request.FormValue("password")
	repassword := c.Request.FormValue("Identity")
	fmt.Println(user.Name, "  >>>>>>>>>>>  ", password, "wdadwadaw", repassword)
	salt := fmt.Sprintf("%06d", rand.Int31())
	data := models.FindUserByName(user.Name)
	if user.Name == "" {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "User name can not be empty!",
			"data":    user,
		})
		return
	}
	if data.Name != "" {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "This user name has been registered",
			"data":    user,
		})
		return
	}
	if password != repassword {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "The two passwords are inconsistent!",
			"data":    user,
		})
		return
	}
	//user.PassWord = password
	user.PassWord = utils.MakePassword(password, salt)
	user.Salt = salt
	fmt.Println(user.PassWord)
	user.LoginTime = time.Now()
	user.LoginOutTime = time.Now()
	user.HeartbeatTime = time.Now()
	models.CreateUser(user)
	c.JSON(200, gin.H{
		"code":    0,
		"message": "New User has been created, welcome to GoChat!",
		"data":    user,
	})
}

// DeleteUser
// @Summary delete user
// @Tags userpage
// @param id query string false "id"
// @Success 200 {string} json{"code", "message"}
// @Router /user/deleteUser [post]
func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)
	models.DeleteUser(user)
	c.JSON(200, gin.H{
		"code":    0,
		"message": "Delete User Successfully!",
		"data":    user,
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
// @Router /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	// fmt.Println("call!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	user.Name = c.PostForm("name")
	user.PassWord = c.PostForm("password")
	user.Phone = c.PostForm("phone")
	user.Avatar = c.PostForm("icon")
	user.Email = c.PostForm("email")
	fmt.Println("update :", user)
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "Modify parameters do not match!",
			"data":    user,
		})
	} else {
		models.UpdateUser(user)
		c.JSON(200, gin.H{
			"code":    0,
			"message": "Update User Info Sucessfully",
			"data":    user,
		})
	}
}

// FindUserByNameAndPwd
// @Summary login
// @Tags userpage
// @param name query string false "name"
// @param password query string false "pasword"
// @Success 200 {string} json{"code","message"}
// @Router /user/findUserByNameAndPwd [post]
func FindUserByNameAndPwd(c *gin.Context) {
	data := models.UserBasic{}

	//name := c.Query("name")
	//password := c.Query("password")
	name := c.Request.FormValue("name")
	password := c.Request.FormValue("password")
	fmt.Println(name, password)
	user := models.FindUserByName(name)
	if user.Name == "" {
		c.JSON(200, gin.H{
			"code":    -1, //  0 success    -1 fail
			"message": "This User Does Not Exist",
			"data":    data,
		})
		return
	}

	flag := utils.ValidPassword(password, user.Salt, user.PassWord)
	if !flag {
		c.JSON(200, gin.H{
			"code":    -1, //  0 success    -1 fail
			"message": "Wrong PassWord",
			"data":    data,
		})
		return
	}
	pwd := utils.MakePassword(password, user.Salt)
	data = models.FindUserByNameAndPwd(name, pwd)

	c.JSON(200, gin.H{
		"code":    0, //  0 success    -1 fail
		"message": "Login Successfully",
		"data":    data,
	})
}

// Prevent Cross-Origin Site Forgery Requests
// to be updated
var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func SendMsg(c *gin.Context) {
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(ws *websocket.Conn) {
		err = ws.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(ws)
	MsgHandler(c, ws)
}

func MsgHandler(c *gin.Context, ws *websocket.Conn) {
	for {
		msg, err := utils.Subscribe(c, utils.PublishKey)
		if err != nil {
			fmt.Println(" MsgHandler Sending Failed", err)
		}
		tm := time.Now().Format("2006-01-02 15:04:05")
		m := fmt.Sprintf("[ws][%s]:%s", tm, msg)
		err = ws.WriteMessage(1, []byte(m))
		if err != nil {
			fmt.Println(err)
		}
	}
}

func SendUserMsg(c *gin.Context) {
	models.Chat(c.Writer, c.Request)
}

func SearchFriends(c *gin.Context) {
	id, _ := strconv.Atoi(c.Request.FormValue("userId"))
	users := models.SearchFriend(uint(id))
	utils.RespOKList(c.Writer, users, len(users))
}

func AddFriend(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Request.FormValue("userId"))
	targetName := c.Request.FormValue("targetName")
	//targetId, _ := strconv.Atoi(c.Request.FormValue("targetId"))
	code, msg := models.AddFriend(uint(userId), targetName)
	if code == 0 {
		utils.RespOK(c.Writer, code, msg)
	} else {
		utils.RespFail(c.Writer, msg)
	}
}

// //新建群
// func CreateCommunity(c *gin.Context) {
// 	ownerId, _ := strconv.Atoi(c.Request.FormValue("ownerId"))
// 	name := c.Request.FormValue("name")
// 	icon := c.Request.FormValue("icon")
// 	desc := c.Request.FormValue("desc")
// 	community := models.Community{}
// 	community.OwnerId = uint(ownerId)
// 	community.Name = name
// 	community.Img = icon
// 	community.Desc = desc
// 	code, msg := models.CreateCommunity(community)
// 	if code == 0 {
// 		utils.RespOK(c.Writer, code, msg)
// 	} else {
// 		utils.RespFail(c.Writer, msg)
// 	}
// }

// //加载群列表
// func LoadCommunity(c *gin.Context) {
// 	ownerId, _ := strconv.Atoi(c.Request.FormValue("ownerId"))
// 	//	name := c.Request.FormValue("name")
// 	data, msg := models.LoadCommunity(uint(ownerId))
// 	if len(data) != 0 {
// 		utils.RespList(c.Writer, 0, data, msg)
// 	} else {
// 		utils.RespFail(c.Writer, msg)
// 	}
// }

// //加入群 userId uint, comId uint
// func JoinGroups(c *gin.Context) {
// 	userId, _ := strconv.Atoi(c.Request.FormValue("userId"))
// 	comId := c.Request.FormValue("comId")

// 	//	name := c.Request.FormValue("name")
// 	data, msg := models.JoinGroup(uint(userId), comId)
// 	if data == 0 {
// 		utils.RespOK(c.Writer, data, msg)
// 	} else {
// 		utils.RespFail(c.Writer, msg)
// 	}
// }

// func FindByID(c *gin.Context) {
// 	userId, _ := strconv.Atoi(c.Request.FormValue("userId"))

// 	//	name := c.Request.FormValue("name")
// 	data := models.FindByID(uint(userId))
// 	utils.RespOK(c.Writer, data, "ok")
// }
