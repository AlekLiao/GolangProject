package main

import (
	"fmt"
	"image/color"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type CustomTheme struct {
	fyne.Theme
}

func (t CustomTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if name == theme.ColorNameDisabled {
		return color.Black
	}
	return t.Theme.Color(name, variant)
}

type CustomEntry struct {
	widget.Entry
}

func NewCustomEntry() *CustomEntry {
	entry := &CustomEntry{}
	entry.ExtendBaseWidget(entry)
	return entry
}

func (e *CustomEntry) Disabled() bool {
	return false // Always return false to keep text color
}

func main() {
	myApp := app.New()
	myApp.Settings().SetTheme(CustomTheme{Theme: theme.DefaultTheme()})
	myWindow := myApp.NewWindow("App.自動下載測試器")

	// 創建第一組組件
	lbl_DstPath := widget.NewLabel("App.本地路徑：")
	edt_DstPath := NewCustomEntry()
	edt_DstPath.SetPlaceHolder("請選擇文件夾")
	edt_DstPath.Disable() // 設置為只讀
	btn_SelLocalPath := widget.NewButton("Open...", func() {
		folder := openFolderDialog()
		if folder != "" {
			edt_DstPath.SetText(folder)
			edt_DstPath.Refresh()
		}
	})

	// 創建第二組組件
	lbl_ipaSrcUrl := widget.NewLabel("ipa下載路徑：")
	edt_ipaSrcUrl := widget.NewEntry()
	edt_ipaSrcUrl.SetPlaceHolder("填入運維發佈的下載鏈接")
	btn_ClearIpaSrcUrl := widget.NewButton("Clear", func() {
		edt_ipaSrcUrl.SetText("")
		edt_ipaSrcUrl.SetPlaceHolder("填入運維發佈的下載鏈接")
	})

	// 創建第三組組件
	lbl_apkSrcUrl := widget.NewLabel("apk下載路徑：")
	edt_apkSrcUrl := widget.NewEntry()
	edt_apkSrcUrl.SetPlaceHolder("填入運維發佈的下載鏈接")
	btn_ClearApkSrcUrl := widget.NewButton("Clear", func() {
		edt_apkSrcUrl.SetText("")
		edt_apkSrcUrl.SetPlaceHolder("填入運維發佈的下載鏈接")
	})

	// 創建第四組組件
	lbl_Interval := widget.NewLabel("下載測試間隔：")
	options := []string{}
	for i := 5; i <= 60; i += 5 {
		options = append(options, strconv.Itoa(i))
	}
	select4 := widget.NewSelect(options, func(value string) {
		fmt.Println("選擇了:", value)
	})
	button4 := widget.NewButton("確定", func() {
		// 按鈕點擊事件處理
	})

	// 添加複選框和選擇音頻文件的按鈕
	checkbox := widget.NewCheck("異常提示音", func(checked bool) {
		fmt.Println("異常提示音:", checked)
	})
	var selectedFile string
	selectFileButton := widget.NewButton("選擇 .mp3", func() {
		file := openFileDialog()
		if file != "" {
			selectedFile = file
			fmt.Println("選擇的mp3文件:", selectedFile)
		}
	})

	// 設置組件的固定大小
	buttonWidth := float32(100)
	windowWidth := float32(630)
	setSize := func(label *widget.Label, input fyne.CanvasObject, button *widget.Button) {
		label.Resize(fyne.NewSize(60, label.MinSize().Height))
		entryWidth := windowWidth - label.Size().Width - buttonWidth - 35 // 35 = 10(左邊距) + 10(label和entry間距) + 10(entry和button間距) + 5(額外間距)
		input.Resize(fyne.NewSize(entryWidth, input.MinSize().Height))
		button.Resize(fyne.NewSize(buttonWidth, button.MinSize().Height))
	}

	setSize(lbl_DstPath, edt_DstPath, btn_SelLocalPath)
	setSize(lbl_ipaSrcUrl, edt_ipaSrcUrl, btn_ClearIpaSrcUrl)
	setSize(lbl_apkSrcUrl, edt_apkSrcUrl, btn_ClearApkSrcUrl)
	setSize(lbl_Interval, select4, button4)

	// 創建自定義容器
	customContainer := container.NewWithoutLayout(
		lbl_DstPath, edt_DstPath, btn_SelLocalPath,
		lbl_ipaSrcUrl, edt_ipaSrcUrl, btn_ClearIpaSrcUrl,
		lbl_apkSrcUrl, edt_apkSrcUrl, btn_ClearApkSrcUrl,
		lbl_Interval, select4, button4,
		checkbox, selectFileButton,
	)
	customContainer.Resize(fyne.NewSize(windowWidth, 220))

	// 設置組件位置的函數
	setPositions := func(label *widget.Label, input fyne.CanvasObject, button *widget.Button, y float32) {
		label.Move(fyne.NewPos(10, y))
		inputX := label.Position().X + label.Size().Width + 10
		input.Move(fyne.NewPos(inputX, y))
		buttonX := windowWidth - buttonWidth - 10 // 右對齊
		button.Move(fyne.NewPos(buttonX, y))
	}

	// 設置四組組件位置
	setPositions(lbl_DstPath, edt_DstPath, btn_SelLocalPath, 10)
	setPositions(lbl_ipaSrcUrl, edt_ipaSrcUrl, btn_ClearIpaSrcUrl, 50)
	setPositions(lbl_apkSrcUrl, edt_apkSrcUrl, btn_ClearApkSrcUrl, 90)
	setPositions(lbl_Interval, select4, button4, 130)

	// 設置複選框和選擇文件按鈕的位置
	checkbox.Resize(fyne.NewSize(120, checkbox.MinSize().Height))
	checkbox.Move(fyne.NewPos(10, 170))
	selectFileButton.Resize(fyne.NewSize(buttonWidth, selectFileButton.MinSize().Height))
	selectFileButton.Move(fyne.NewPos(140, 170)) // 直接放在複選框右側

	// 创建状态栏
	statusBar := widget.NewLabel("")
	updateTime := func() {
		currentTime := time.Now().Format("2006-01-02 15:04:05")
		statusBar.SetText(currentTime)
	}
	updateTime() // 初始化时间

	// 创建一个计时器，每秒更新一次时间
	go func() {
		ticker := time.NewTicker(time.Second)
		for range ticker.C {
			updateTime()
		}
	}()

	// 使用 BorderLayout 将状态栏添加到底部
	borderContainer := container.NewBorder(nil, statusBar, nil, nil, customContainer)

	myWindow.SetContent(borderContainer)

	// 設置固定窗口大小
	windowSize := fyne.NewSize(windowWidth, 240) // 增加高度以容纳状态栏
	myWindow.Resize(windowSize)
	myWindow.SetFixedSize(true) // 禁用窗口調整大小

	myWindow.ShowAndRun()
}

func openFileDialog() string {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("osascript", "-e", `choose file of type {"mp3"} with prompt "請選擇一個MP3文件"`)
	case "windows":
		cmd = exec.Command("powershell", "-Command", "Add-Type -AssemblyName System.Windows.Forms; $f = New-Object System.Windows.Forms.OpenFileDialog; $f.Filter = 'MP3 Files (*.mp3)|*.mp3'; $f.ShowDialog(); $f.FileName")
	default: // Linux and others
		cmd = exec.Command("zenity", "--file-selection", "--file-filter=*.mp3")
	}

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}

	return strings.TrimSpace(filepath.Clean(string(output)))
}

func openFolderDialog() string {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("osascript", "-e", `choose folder with prompt "請選擇一個文件夾"`)
	case "windows":
		cmd = exec.Command("powershell", "-Command", "Add-Type -AssemblyName System.Windows.Forms; $f = New-Object System.Windows.Forms.FolderBrowserDialog; $f.ShowDialog(); $f.SelectedPath")
	default: // Linux and others
		cmd = exec.Command("zenity", "--file-selection", "--directory")
	}

	//output, err := cmd.CombinedOutput()
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	fmt.Println(string(output))

	// 移除可能的 "OK" 後綴
	folder := strings.TrimSpace(string(output))
	folder = strings.TrimSuffix(folder, "OK")
	folder = strings.TrimPrefix(folder, "OK")
	folder = strings.TrimSuffix(folder, "Cancel")
	folder = strings.TrimPrefix(folder, "Cancel")
	folder = strings.TrimSpace(folder) // 再次去除可能的空白

	return filepath.Clean(folder)
}
