package service

import (
	"text/template"

	"github.com/gin-gonic/gin"
)

// GetIndex
// @Tags mainpage
// @Success 200 {string} welcome
// @Router /index [get]
func GetIndex(c *gin.Context) {
	ind, err := template.ParseFiles("index.html", "views/chat/head.html")
	if err != nil {
		panic(err)
	}
	ind.Execute(c.Writer, "index")
	// c.JSON(200, gin.H{
	// 	"message": "welcome to GoChat!",
	// })
}
