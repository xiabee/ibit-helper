package main

import (
	"fmt"
	"os"
	"xiabee/pre"
)

type Student struct {
	Name  string
	Id    string
	Class string
}

func main() {
	argc := len(os.Args)
	argv := os.Args

	comp_str := "学工号"
	to_str := "学号"
	str1 := "姓名"
	str2 := "班级标注"
	var stu []Student

	if argc != 3 {
		fmt.Println("usage: " + argv[0] + "  已完成核酸名单.xlsx  全体学生名单.xlsx")
		os.Exit(0)
	}
	// 获取命令行argv

	completed := pre.Init(string(argv[1]))
	// 读取已完成核酸人员名单

	total := pre.Init(string(argv[2]))
	// 读取总人员名单

	comp := pre.Load_data(completed, comp_str)
	to := pre.Load_data(total, to_str)
	// 获取学号
	name := pre.Load_data(total, str1)
	class := pre.Load_data(total, str2)
	// 获取相关信息

	for i := 0; i < len(to); i++ {
		tmp := Student{Id: to[i], Name: name[i], Class: class[i]}
		stu = append(stu, tmp)
	}
	// 构建结构体列表

	incomplete := pre.Diff(to, comp) //对总名单求差集，得到未完成核酸的学号列表

	for i := 0; i < len(to); i++ {
		for j := 0; j < len(incomplete); j++ {
			if incomplete[j] == stu[i].Id {
				fmt.Println(stu[i].Id, stu[i].Name, stu[i].Class)
			}
		}
	}
}
