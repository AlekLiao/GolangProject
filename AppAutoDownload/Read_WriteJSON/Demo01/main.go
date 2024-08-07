package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type config struct {
	ApkSrc    string `json:"ApkSrc"`
	IpaSrc    string `json:"IpaSrc"`
	ApkDst    string `json:"ApkDst"`
	IpaDst    string `json:"IpaDst"`
	DownTimer string `json:"Interval"`
}

var configFile = "cfg.json"

func main() {
	fmt.Print("read and write JSON file")

	readJSON()
	//writeJSON()
}

// https://tutorialedge.net/golang/parsing-json-with-golang/
func readJSON() {
	jsonFile, err := os.Open(configFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully Opened users.json")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var cfg config
	json.Unmarshal(byteValue, &cfg)

	fmt.Println("Src. of .apk:" + cfg.ApkSrc)
	fmt.Println("Src. of .ipa:" + cfg.IpaSrc)
	fmt.Println("Dst. of .apk:" + cfg.ApkDst)
	fmt.Println("Dst. of .ipk:" + cfg.IpaDst)
	fmt.Println("download interval:" + cfg.DownTimer)
}

// https://www.developer.com/languages/json-files-golang/
// https://www.cnblogs.com/cenjw/p/go-ioutil-writefile-perm.html
func writeJSON() {
	cfg := config{}

	cfg.ApkDst = "12324"
	cfg.ApkSrc = "abcdw"
	cfg.IpaSrc = "Peter"
	cfg.IpaDst = "asdasaasdas"
	cfg.DownTimer = "25"

	content, err := json.Marshal(cfg)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ioutil.WriteFile(configFile, content, 0644)
	if err != nil {
		log.Fatal(err)
	}

}
