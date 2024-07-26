// https://www.bilibili.com/video/BV1XY4y1t76G?p=52

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
特別注意以下三個 struct。
UserData1 的每個element name 都是大寫字母：Name、Pw、Desc
在 aboutUser1() 回傳 UserData1 的時候正常。

但UserData2 的每個element name 都是小寫字母：name、pw、desc
在 aboutUser2() 回傳 UserData2 的時候是空字串。

UserData3 的element name 大小寫混合：Name、pw、desc
在 aboutUser3() 回傳 UserData3 的時候，Name 可正常出現，其餘是空字串。

該情況在編譯時不會報錯，要特別小心！！
*/

/*
UserData1 struct 宣告type 之後，接在“內的，是回傳給client 的鍵值（key）的名稱
以 UserData1 struct 為例，回傳後在browser 看到的是: {"姓名":"西瓜","password":"hidden","Desc.描述":"高富帥"}
*/
type UserData1 struct {
	Name string `json:"姓名"`       // name
	Pw   string `json:"password"` // password
	Desc string `json:"Desc.描述"`  // note
}

type UserData2 struct {
	name string // name
	pw   string // password
	desc string // note
}

type UserData3 struct {
	Name string // name
	pw   string // password
	desc string // note
}

func main() {
	r := gin.Default()
	// 配置template 文件的路徑。一定要先寫這行
	// 這個路徑表示模板文件都在相對的 "template/*" 路徑底下
	r.LoadHTMLGlob("template/*")

	r.GET("/", IsServerRunning)
	r.GET("/index", IndexPage)
	r.GET("/json", getJSON)

	r.GET("/aboutUser1", aboutUser1)
	r.GET("/aboutUser2", aboutUser2)
	r.GET("/aboutUser3", aboutUser3)

	r.GET("/xml", getXML)

	r.Run()
}

func IsServerRunning(c *gin.Context) {
	c.String(http.StatusOK, "Server is runnig...")
}

func getJSON(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "你好",
		"pong":    "我也很好",
	})
}

func aboutUser1(c *gin.Context) {
	user := &UserData1{
		Name: "西瓜",
		Pw:   "hidden",
		Desc: "高富帥",
	}
	c.JSON(http.StatusOK, user)
}

func aboutUser2(c *gin.Context) {
	user := &UserData2{
		name: "西瓜",
		pw:   "hidden",
		desc: "高富帥",
	}
	c.JSON(http.StatusOK, user)
}

func aboutUser3(c *gin.Context) {
	user := &UserData3{
		Name: "西瓜",
		pw:   "hidden",
		desc: "高富帥",
	}
	c.JSON(http.StatusOK, user)
}

/*
// 貌似不能這樣用
// 出現 “Error #01: xml: unsupported type: map[string]interface {}”
// 但仍能執行，可是browser 出現“This XML file does not appear to have any style information associated with it. The document tree is shown below.” 提示

func getXML(c *gin.Context) {
	c.XML(http.StatusOK,  map[string]interface{} {
		"success": true,
		"msg":     "hi, 你好",
	})
}
*/

func getXML(c *gin.Context) {
	c.XML(http.StatusOK, gin.H{
		"success": true,
		"msg":     "hi, 你好",
	})
}

func IndexPage(c *gin.Context) {
	// 因為已經有先利用“LoadHTMLGlob()” 設置template 路徑，所以可以直接load folder 裡的文件
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "我是後台數據",
	})
}
