// https://www.bilibili.com/video/BV1XY4y1t76G?p=55&vd_source=a7e3b225ce17b0febfcfa4203e1d833c

package main

import (
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Username string `form:"username" json:"username"` // json:"username" 這段非必要
	Password string `form:"password" json:"password"`
}

type Article struct {
	Title   string `xml:"title" json:"title"`
	Content string `xml:"content" json:"content"`
}

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")
	r.GET("/", defaultPage)
	r.GET("/user", userPage)
	r.GET("/getuser", getUser)

	// 動態路由。例如 list/123; list/abc
	r.GET("/list/:cid", List)

	r.POST("/doAddUser1", AddUser1)
	r.POST("/doAddUser2", AddUser2)
	r.POST("/xml", getXML)

	r.Run()

}

// 動態路由
func List(c *gin.Context) {
	cid := c.Param("cid")
	c.String(200, "%v", cid) // 輸出到瀏覽器
}

func getXML(c *gin.Context) {
	article := &Article{}
	xmlSliceData, _ := c.GetRawData()
	fmt.Println(xmlSliceData)

	if err := xml.Unmarshal(xmlSliceData, &article); err == nil {
		// 成功
		c.JSON(http.StatusOK, article)
	} else {
		// 失敗
		c.JSON(http.StatusBadRequest, article)
	}
}

func defaultPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		// golang 裡邊，因為是UTF-8編碼，所以一個中文字是3個bytes
		"title":    "這是title",
		"subtitle": "subtitle",
		"this":     "這是",
		"date":     1721383029,
	})
}

// load html file
func userPage(c *gin.Context) {
	c.HTML(http.StatusOK, "userPage.html", gin.H{})
}

func AddUser1(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	//age := c.PostForm("age")
	age := c.DefaultPostForm("age", "25")

	c.JSON(http.StatusOK, gin.H{
		"username": username,
		"password": password,
		"age":      age,
	})
}

// 利用POST method 獲取數據，傳入struct
func AddUser2(c *gin.Context) {
	user := &UserInfo{}
	if err := c.ShouldBind(&user); err == nil {
		fmt.Printf("%#v", user)
		c.JSON(http.StatusOK, user)

	} else {
		c.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
	}
}

// http://localhost:8080/getuser?username=zhang&password=123456
func getUser(c *gin.Context) {
	user := &UserInfo{}

	if err := c.ShouldBind(&user); err == nil {
		fmt.Printf("%#v", user)
		c.JSON(http.StatusOK, user)

	} else {
		c.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
	}
}
