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
	file1 := ""
	file2 := ""

	msg_state := widget.NewLabel("")
	msg_file1 := widget.NewLabel("")
	msg_file2 := widget.NewLabel("")
	msg_end := widget.NewLabel("")

	dl1 := dialog.NewFileOpen(func(uc fyne.URIReadCloser, e error) {
		if uc == nil {
			return
		}
		msg_file1.SetText(uc.URI().Path())
		file1 = uc.URI().Path()
	}, w)

	dl2 := dialog.NewFileOpen(func(uc fyne.URIReadCloser, e error) {
		if uc == nil {
			return
		}
		msg_file2.SetText(uc.URI().Path())
		file2 = uc.URI().Path()
	}, w)
	// 输入文件按钮

	bt1 := widget.NewButton("运行", func() {
		if file1 != "" && file2 != "" {
			msg_state.SetText("运行中")
			flag := handle.Handle(file1, file2)
			msg_state.SetText("已完成")
			if flag {
				msg_end.SetText("执行失败，请检查输入文件是否正确！")
			} else {
				msg_end.SetText("执行成功，生成文件：未完成.csv")
			}
		} else {
			msg_end.SetText("请选择文件")
		}

	})
	// 启动按钮

	uri := storage.NewFileURI(".")
	luri, _ := storage.ListerForURI(uri)
	dl1.SetLocation(luri)
	dl2.SetLocation(luri)
	// 默认打开当前目录
	btn1 := widget.NewButton("已完成核酸检测的名单", func() {
		dl1.Resize(fyne.NewSize(800, 400))
		dl1.SetFilter(storage.NewExtensionFileFilter([]string{".xlsx"}))
		dl1.Show()
	})

	btn2 := widget.NewButton("全体成员名单", func() {
		dl2.Resize(fyne.NewSize(800, 400))
		dl2.SetFilter(storage.NewExtensionFileFilter([]string{".xlsx"}))
		dl2.Show()
	})

	c := container.NewVBox(msg_file1, btn1, msg_file2, btn2, msg_state, bt1, msg_end)
	w.SetContent(c)
}

// 打开文件窗口

func Mainform() {
	a := app.NewWithID("xiabee")
	w := a.NewWindow("核酸检测未完成 成员筛查")
	fileopen(w)
	w.Resize(fyne.NewSize(800, 400))
	w.ShowAndRun()
	fmt.Println("End")
	os.Unsetenv("FYNE_FONT")
	// 设置中文字体
}

// 主程序窗口
