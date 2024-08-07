/*
這個程序做了以下幾件事:

設置了要下載的URL和保存的文件名。
使用time.NewTicker創建一個定時器,每小時觸發一次。
在主循環中,每當定時器觸發時,就調用downloadFile函數。
downloadFile函數使用http.Get發送GET請求下載文件,然後將內容保存到本地文件中。

要使用這個程序,你需要:

將url變量修改為你想下載的文件的實際URL。
如果需要,修改filename變量為你想保存的文件名。
如果你想更改下載頻率,可以調整time.NewTicker(1 * time.Hour)中的時間間隔。

請注意,這個程序會一直運行,直到你手動停止它。在實際使用中,你可能需要添加錯誤處理、日誌記錄、重試機制等功能來增強其可靠性。
*/

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type mobileFile struct {
	ApkURL string `json:"APKURL"`
	IpaURL string `json:"IPAURL"`
}

func main() {
	// 設置下載的URL和保存的文件名
	//apkUrl := "https://apk.ccj0qjz8os.com/cai8.apk"
	apkUrl := "https://health-hp.tncghb.gov.tw/nutrition_zone/download/7de18763-a6ac-4fe5-9cae-454a3886200a"
	apkLocalName := "D:\\AppTemp\\cai8.pdf"
	//ipaUrl := "iOS"
	//ipaLocalName := "https://update.etm4ucefpo.com/enterprise_cp8_x_btwo_240730_113000_resign.ipa"

	/*
		url := "https://example.com/file.zip"
		filename := "downloaded_file.zip"
	*/

	// 設置定時器,每小時執行一次
	//ticker := time.NewTicker(1 * time.Hour)
	ticker := time.NewTicker(10 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			err := downloadFile(apkUrl, apkLocalName)
			if err != nil {
				fmt.Printf("下載失敗: %v\n", err)
			} else {
				fmt.Printf("文件成功下載於 %v\n", time.Now())
			}
		}
	}

}

// 從外部讀取 apk 及 ipa 的URL
func getSrcPath() error {
	configFile := "cfg.json"
	cfgFile, err := os.Open(configFile)
	if err != nil {
		log.Fatal(err)
	}
	defer cfgFile.Close()
}

func downloadFile(url string, filepath string) error {
	// 發送GET請求
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 創建文件
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// 將響應內容寫入文件
	_, err = io.Copy(out, resp.Body)
	return err
}
