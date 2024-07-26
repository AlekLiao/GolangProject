// https://www.bilibili.com/video/BV1XY4y1t76G?p=51

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// gin.SetMode(gin.DebugMode) // debug 模式
	// gin.SetMode(gin.ReleaseMode)// release 模式

	// 初始化一个http服务对象
	r := gin.Default()

	//配置各路由
	r.GET("/", IsServerRunning)
	r.GET("/login", doLogin)

	r.POST("/add", addUser)

	r.PUT("/edit", editData)

	r.DELETE("/delete", delData)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func doLogin(c *gin.Context) {
	c.String(http.StatusOK, "do login job")
}

func IsServerRunning(c *gin.Context) {
	c.String(http.StatusOK, "Server is runnig...")
}

func addUser(c *gin.Context) {
	c.String(http.StatusOK, "add user")
}

func editData(c *gin.Context) {
	c.String(http.StatusOK, "edit data")
}

func delData(c *gin.Context) {
	c.String(http.StatusOK, "delete data")
}
