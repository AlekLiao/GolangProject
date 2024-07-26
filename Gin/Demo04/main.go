// https://www.bilibili.com/video/BV1XY4y1t76G?p=54&vd_source=a7e3b225ce17b0febfcfa4203e1d833c
// 沒看完

package main

import (
	"net/http"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
)

func UnixToTime(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

func Printline(str1 string, str2 string) string {
	return str1 + str2
}

func main() {
	r := gin.Default()

	// 自定義模版函數一定要寫在 gin.Default() 之後，LoadHTMLGlob() 之前
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": UnixToTime,
		"Printline":  Printline,
	})

	r.LoadHTMLGlob("templates/*")
	r.GET("/", defaultPage)

	r.Run()

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
