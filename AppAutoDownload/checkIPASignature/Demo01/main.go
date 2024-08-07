package main

import (
	"archive/zip"
	"fmt"
	"io"
	"time"
)

// 定義一個簡單的 ASN.1 結構
type MobileProvision struct {
	Name           string
	ExpirationDate time.Time
	Entitlements   map[string]interface{}
}

func checkIPASignature(ipaPath string) (bool, error) {
	// 1. 打開 IPA 文件
	reader, err := zip.OpenReader(ipaPath)
	if err != nil {
		fmt.Println("打開 IPA 文件 fail")
		return false, err
	}
	defer reader.Close()

	
	// 2. 尋找並讀取 embedded.mobileprovision 文件
	// var provisionData []byte
	for _, file := range reader.File {
		if file.Name == "D://AppTemp//彩票_signed//Payload//DemoOC.app//embedded.mobileprovision" {
			rc, err := file.Open()
			if err != nil {
				fmt.Println("打開 Iembedded.mobileprovision 文件 fail")
				return false, err
			}

			//provisionData, err = ioutil.ReadAll(rc)
			provisionData, err := io.ReadAll(rc)
			rc.Close()
			if err != nil {
				fmt.Println("讀取 Iembedded.mobileprovision 文件 fail")
				return false, err
			}
			fmt.Println(provisionData)
			break
		}
	}

	// 3. 解析 mobileprovision 文件
	// 這裡需要實現 ASN.1 解碼來提取過期時間
	var expirationDate time.Time
	// ... 解析 ASN.1 結構，提取 ExpirationDate ...
	/*
		rawData := []byte{}
		var provision MobileProvision
		_, err := asn1.Unmarshal(rawData, &provision)
		if err != nil {
			fmt.Println("解析錯誤:", err)
			return false, err
		}

		fmt.Println("名稱:", provision.Name)
		fmt.Println("過期日期:", provision.ExpirationDate)

		fmt.Println("權限:", provision.Entitlements)
	*/
	// 4. 檢查是否過期
	if time.Now().After(expirationDate) {
		fmt.Println("檢查是否過期 fail")
		return false, nil // 已過期
	}

	return true, nil // 未過期
}

func main() {
	valid, err := checkIPASignature("D:/AppTemp/彩票_signed.ipa")
	if err != nil {
		fmt.Println("檢查錯誤:", err)
		return
	}
	if valid {
		fmt.Println("IPA 簽名有效")
	} else {
		fmt.Println("IPA 簽名已失效")
	}
}
