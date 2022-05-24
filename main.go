package main

import (
	"fmt"
	"os"
	"xiabee/gui"
	"xiabee/handle"
)

func main() {
	argc := len(os.Args)
	argv := os.Args
	// 获取命令行argv

	if argc == 1 {
		gui.Mainform()
		// 启动GUI
	} else if argc == 3 {
		handle.Handle(argv[1], argv[2])
		// 直接执行命令行
	} else {
		fmt.Println("usage: " + argv[0] + "  已完成核酸名单.xlsx  全体学生名单.xlsx")
		os.Exit(0)
	}
}
