package main

import (
	"fmt"
	"log"
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
	// “学工号”为已做核酸名单中的索引

	to_str := "学号"
	str1 := "姓名"
	str2 := "班级标注"
	// 此三者为总名单中的索引

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

	outputfile, err := os.OpenFile("./未完成.csv", os.O_WRONLY|os.O_CREATE, 0)
	if err != nil {
		log.Println(err)
	}
	defer outputfile.Close()
	// 写入文件
	outputfile.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM，防止中文乱码

	fmt.Fprintf(outputfile, "序号,学号,姓名,班级\n")
	for i := 0; i < len(to); i++ {
		for j := 0; j < len(incomplete); j++ {
			if incomplete[j] == stu[i].Id {
				// fmt.Println(stu[i].Id, stu[i].Name, stu[i].Class)
				fmt.Fprintf(outputfile, "%d,%s,%s,%s\n", i+1, stu[i].Id, stu[i].Name, stu[i].Class)
			}
		}
	}
	fmt.Println("Success!")
	fmt.Println("生成文件：未完成.csv")
}
