// https://www.bilibili.com/video/BV1XY4y1t76G?p=56
package main

import (
	"fmt"
	"gitTest/Gin/Demo06/routers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type loginInfo struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
	//captcha  string `JSON:"captcha"`
}

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("*.*")

	routers.AdminRoutersInit(router)

	router.GET("/", homePage)
	router.POST("/VerifyLogin", VerifyLogin)

	router.Run()

}

func homePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})

	//c.String(http.StatusOK, "this should be index page")
}

func VerifyLogin(c *gin.Context) {
	/*
		username := c.PostForm("username")
		password := c.PostForm("password")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	*/

	userLoginInfo := &loginInfo{}
	if err := c.ShouldBind(&userLoginInfo); err == nil {
		fmt.Printf("%#v", userLoginInfo)
		c.JSON(http.StatusOK, userLoginInfo)

	} else {
		c.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
	}
}
