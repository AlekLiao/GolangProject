package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/fullsailor/pkcs7"
	"howett.net/plist"
)

type ProvisioningProfile struct {
	ExpirationDate              time.Time
	Name                        string
	TeamIdentifier              []string
	AppIDName                   string
	ApplicationIdentifierPrefix []string
}

func extractProvisioningProfileInfo(data []byte) (*ProvisioningProfile, error) {
	// 解析 CMS 結構
	p7, err := pkcs7.Parse(data)
	if err != nil {
		return nil, fmt.Errorf("解析 CMS 結構失敗: %v", err)
	}

	// 獲取 CMS 結構中的內容
	content := p7.Content

	// 解析 plist 數據
	var provisioningProfile map[string]interface{}
	decoder := plist.NewDecoder(bytes.NewReader(content))
	err = decoder.Decode(&provisioningProfile)
	if err != nil {
		return nil, fmt.Errorf("解析 plist 失敗: %v", err)
	}

	// 提取所需信息
	profile := &ProvisioningProfile{}

	if expirationDate, ok := provisioningProfile["ExpirationDate"].(time.Time); ok {
		profile.ExpirationDate = expirationDate
	}

	if name, ok := provisioningProfile["Name"].(string); ok {
		profile.Name = name
	}

	if teamIdentifier, ok := provisioningProfile["TeamIdentifier"].([]string); ok {
		profile.TeamIdentifier = teamIdentifier
	}

	if appIDName, ok := provisioningProfile["AppIDName"].(string); ok {
		profile.AppIDName = appIDName
	}

	if applicationIdentifierPrefix, ok := provisioningProfile["ApplicationIdentifierPrefix"].([]string); ok {
		profile.ApplicationIdentifierPrefix = applicationIdentifierPrefix
	}

	return profile, nil
}

func readMobileProvision(ipaPath string) ([]byte, error) {
	reader, err := zip.OpenReader(ipaPath)
	if err != nil {
		return nil, fmt.Errorf("Unable to open ipa file :%v", err)
	}
	defer reader.Close()

	var provisionFile *zip.File
	for _, file := range reader.File {
		if strings.HasSuffix(file.Name, "embedded.mobileprovision") {
			provisionFile = file
			break
		}
	}

	if provisionFile == nil {
		return nil, fmt.Errorf("Unable to find embedded.mobileprovision file")
	}

	// 打開 embedded.mobileprovision 文件
	rc, err := provisionFile.Open()
	if err != nil {
		return nil, fmt.Errorf("Unable to open embedded.mobileprovision 文件: %v", err)
	}
	defer rc.Close()

	// 讀取文件內容
	content, err := io.ReadAll(rc)
	if err != nil {
		return nil, fmt.Errorf("Unable to read embedded.mobileprovision file content: %v", err)
	}

	return content, nil
}

func main() {
	ipaPath := "D:/AppTemp/彩票_signed.ipa"
	content, err := readMobileProvision(ipaPath)
	if err != nil {

		fmt.Println("readMobileProvision() failed with err = ", err)
		return
	}
	fmt.Printf("read embedded.mobileprovision file successful. file size = : %d bytes\n", len(content))

	profile, err := extractProvisioningProfileInfo(content)
	if err != nil {
		fmt.Println("Unable to get embedded.mobileprovision data with error = ", err)
		return
	}

	fmt.Printf("過期日期: %v\n", profile.ExpirationDate)
	fmt.Printf("名稱: %s\n", profile.Name)
	fmt.Printf("團隊標識符: %v\n", profile.TeamIdentifier)
	fmt.Printf("App ID 名稱: %s\n", profile.AppIDName)
	fmt.Printf("應用標識符前綴: %v\n", profile.ApplicationIdentifierPrefix)

	// 檢查是否過期
    if time.Now().After(profile.ExpirationDate) {
        fmt.Println("配置文件已過期")
    } else {
        fmt.Println("配置文件有效")
    }
}
