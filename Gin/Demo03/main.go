// https://www.bilibili.com/video/BV1XY4y1t76G?p=53

package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Title   string
	Content string
}

type RequestInfo struct {
	URL string `json:"URL路徑參數"`
	IP  string `json:"User IP"`
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("template/*") // LoadHTMLGlob() 指定某個folder
	//r.LoadHTMLFiles("template/index.html") //LoadHTMLFiles() 要指定檔案

	r.GET("/", IsServerRunning)
	r.GET("/news", getNewsPage)
	r.GET("/betting", BettingPage)

	r.Run()
}

func IsServerRunning(c *gin.Context) {
	Info := RequestInfo{}

	// get URL
	// 如果是127.0.0.1，IP會顯示 127.0.0.1
	// 如果是localhost，IP會顯示 ::1
	Info.IP = c.ClientIP()
	Info.URL = c.Param("id")
	log.Println(Info.IP)
	log.Println(Info.URL)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":   "Welcome to here.",
		"userIP":  Info.IP,
		"userURL": Info.URL,
	})

}

func getNewsPage(c *gin.Context) {
	news := &Article{
		Title:   "新聞標題",
		Content: "新聞內容",
	}
	c.HTML(http.StatusOK, "news.html", gin.H{
		"news": news,
	})
}

func BettingPage(c *gin.Context) {
	c.HTML(http.StatusOK, "betting.html", gin.H{
		"betting": 350,
		"people":  []string{"Mary", "John", "haha", "Tom"},
		"newsList": []interface{}{
			&Article{
				Title:   "新聞標題1",
				Content: "新聞內容1",
			},
			&Article{
				Title:   "新聞標題2",
				Content: "新聞內容2",
			},
		},
		"newPeople": []string{},
	})

}
