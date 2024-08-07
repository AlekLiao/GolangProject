package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

const componentHInterval = 5  // 橫向間距
const componentVInterval = 10 // 直向間距
const componentStartX = 10    // X 軸起點
const componentStartY = 10    // Y 軸起點
const entryHeight = 10

func main() {
	a := app.New()
	w := a.NewWindow("App. 自動下載測試器")
	// w.Resize(fyne.NewSize(800, 450))
	//w.SetFixedSize(true)

	label_localPath := widget.NewLabel("App. 本地路徑")
	//label_localPath.Move(fyne.NewPos(componentStartX, componentStartY))
	label_localPath.Resize(fyne.NewSize(60, label_localPath.MinSize().Height))

	entry_localPath := widget.NewEntry()
	entry_localPath.SetPlaceHolder("填入app. 存儲路徑")
	//entry_localPath.Resize(fyne.NewSize(500, 30))
	//entry_localPath.Move(fyne.NewPos(150, 10))
	entry_localPath.Resize(fyne.NewSize(500, label_localPath.MinSize().Height))

	btn_OpenLocalPath := widget.NewButton("Open...", nil)
	//btn_OpenLocalPath.Resize(fyne.NewSize(100, 10))
	//btn_OpenLocalPath.Move(fyne.NewPos(655, 10))
	btn_OpenLocalPath.Resize(fyne.NewSize(80, label_localPath.MinSize().Height))

	// 創建一個容器，使用 Center 佈局
	content := container.New(layout.NewCenterLayout(),
		container.New(layout.NewHBoxLayout(),
			layout.NewSpacer(),
			label_localPath,
			layout.NewSpacer(),
			entry_localPath,
			layout.NewSpacer(),
			btn_OpenLocalPath,
			layout.NewSpacer(),
		),
	)

	w.SetContent(content)
	/*
		localPathContainer := (container.NewHBox(
			label_localPath,
			entry_localPath,
			btn_OpenLocalPath,
		))
		localPathContainer.Move(fyne.NewPos(10, 10))
	*/
	/*
		w.SetContent(
			container.NewWithoutLayout(
				label_localPath,
				entry_localPath,
				btn_OpenLocalPath,

			),
		)
	*/

	/*
		entry_ipaUrl := widget.NewEntry()
		entry_ipaUrl.SetPlaceHolder("填入ipa 下載鏈接")
		entry_ipaUrl.Resize(fyne.NewSize(500, 30))
		entry_ipaUrl.Move(fyne.NewPos(40, 100))

		entry_apkUrl := widget.NewEntry()
		entry_apkUrl.SetPlaceHolder("填入ipa 下載鏈接")
		entry_apkUrl.Resize(fyne.NewSize(500, 30))
		entry_apkUrl.Move(fyne.NewPos(40, 150))

		w.SetContent(
			container.NewWithoutLayout(
				localPathContainer,
				entry_ipaUrl,
				entry_apkUrl,
			),
		)
	*/
	/*
		w.SetContent(container.NewVBox(
			localPath,
			entry_localPath,
			entry_ipaUrl,
			entry_apkUrl,
			widget.NewButton("Open...", func() {

			}),
		))
	*/
	/*

		w.SetContent(container.NewVBox(
			localPath,
			widget.NewButton("Hi!", func() {
				localPath.SetText("Welcome :)")
			}),
		))
	*/

	w.ShowAndRun()
}
