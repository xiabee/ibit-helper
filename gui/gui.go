package gui

import (
	"fmt"
	"os"
	"xiabee/handle"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

func fileopen(w fyne.Window) {
	var file1, file2 string
	lblMsg := widget.NewLabel("")
	lblMsg1 := widget.NewLabel("")
	lblMsg2 := widget.NewLabel("")
	dl1 := dialog.NewFileOpen(func(uc fyne.URIReadCloser, e error) {
		if uc == nil {
			return
		}
		lblMsg1.SetText(uc.URI().Path())
		file1 = uc.URI().Path()
	}, w)

	dl2 := dialog.NewFileOpen(func(uc fyne.URIReadCloser, e error) {
		if uc == nil {
			return
		}
		lblMsg2.SetText(uc.URI().Path())
		file2 = uc.URI().Path()
	}, w)
	// 输入文件按钮

	bt1 := widget.NewButton("运行", func() {
		lblMsg.SetText("运行中")
		handle.Handle(file1, file2)
		lblMsg.SetText("已完成")
	})
	// 启动按钮
	uri := storage.NewFileURI(".")
	luri, _ := storage.ListerForURI(uri)
	dl1.SetLocation(luri)
	dl2.SetLocation(luri)
	// 默认打开当前目录
	btn1 := widget.NewButton("已打卡名单", func() {
		dl1.Resize(fyne.NewSize(800, 400))
		dl1.SetFilter(storage.NewExtensionFileFilter([]string{".xlsx"}))
		dl1.Show()
	})

	btn2 := widget.NewButton("全体成员名单", func() {
		dl2.Resize(fyne.NewSize(800, 400))
		dl2.SetFilter(storage.NewExtensionFileFilter([]string{".xlsx"}))
		dl2.Show()
	})

	c := container.NewVBox(lblMsg1, btn1, lblMsg2, btn2, lblMsg, bt1)
	w.SetContent(c)
}

// 打开文件窗口

func Mainform() {
	a := app.NewWithID("xiabee")
	w := a.NewWindow("核酸检测打卡筛选")
	fileopen(w)
	w.Resize(fyne.NewSize(600, 600))
	w.ShowAndRun()
	fmt.Println("End")
	os.Unsetenv("FYNE_FONT")
	// 设置中文字体
}

// 主程序窗口
