// https://dev.twsiyuan.com/2017/04/calculate-file-chunks-md5-checksum-golang.html
// https://cloud.tencent.com/developer/article/1061238
// https://jingyan.baidu.com/article/fc07f9891bd26353fee51920.html  Windows 10 如何取得MD5值
// https://hackmd.io/@Not/rkh6ff5nd

package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	fmt.Println("test")

	SrcUrl := "https://www.dwsamplefiles.com/?dl_id=353"
	DstFile := "D://AppTemp//download_file.mp4"
	exceptedMD5 := "8038c46439162f7cf65a8f28e2813a40"

	err := DownloadFile(DstFile, SrcUrl)
	if err != nil {
		fmt.Println("檔案下載錯誤：", err)
		return
	}
	fmt.Println("檔案下載完成")

	isCorrect, err := checkFile(DstFile, exceptedMD5)
	if err != nil {
		fmt.Println("檔案檢查時錯誤：", err)
		return
	}

	if isCorrect {
		fmt.Println("檔案校驗正確")
	} else {
		fmt.Println("檔案校驗錯誤")
	}

	realMD5, err := FileMD5(DstFile)
	fmt.Println("realMD5 = ", realMD5)
}

func FileMD5(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	hash := md5.New()
	_, _ = io.Copy(hash, file)
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func checkFile(filepath string, expectedMD5 string) (bool, error) {
	// 打開文件
	f, err := os.Open(filepath)
	if err != nil {
		return false, err
	}
	defer f.Close()

	// 創建MD5 hash
	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return false, err
	}

	// 獲取MD5校驗和
	actualMD5 := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println("MD5 =", actualMD5)

	// 比較MD5校驗和
	return actualMD5 == expectedMD5, nil
}

func DownloadFile(filepath string, url string) error {
	// 創建文件
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// 發送GET請求
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 檢查響應狀態
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// 寫入文件
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
