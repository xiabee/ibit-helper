package handle

import (
	"fmt"
	"log"
	"os"
	"xiabee/pre"
)

type Student struct {
	Name string
	Id   string
}

func Handle(file1, file2 string) bool {
	flag := false
	// 标记是否panic

	comp_str := "学工号"
	// “学工号”为已做核酸名单中的索引

	title := Student{Id: "学号", Name: "姓名"}
	// 此二者为总名单中的索引

	var stu []Student

	completed := pre.Init(file1)
	// 读取已完成核酸人员名单

	total := pre.Init(file2)
	// 读取总人员名单

	comp := pre.Load_data(completed, comp_str)
	stu_id := pre.Load_data(total, title.Id)
	// 获取学号
	name := pre.Load_data(total, title.Name)
	// 获取相关信息

	if len(name) != len(stu_id) {
		flag = true
		fmt.Println("File Error! Please check your inputs!")
		return flag
	}
	// 读取文件不对，导致不到正确的信息

	for i := 0; i < len(stu_id); i++ {
		tmp := Student{Id: stu_id[i], Name: name[i]}
		stu = append(stu, tmp)
		defer func() bool {
			if err := recover(); err != nil {
				fmt.Printf("recover: %v\n", err)
				flag = true
				// 判断是否溢出
			}
			return flag
		}()
	}

	incomplete := pre.Diff(stu_id, comp) //对总名单求差集，得到未完成核酸的学号列表

	outputfile, err := os.OpenFile("./未完成.csv", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		log.Println(err)
		flag = true
	}
	defer outputfile.Close()
	// 写入文件
	outputfile.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM，防止中文乱码

	fmt.Fprintf(outputfile, "序号,学号,姓名\n")
	count := 0
	for i := 0; i < len(stu_id); i++ {
		for j := 0; j < len(incomplete); j++ {
			if incomplete[j] == stu[i].Id {
				count++
				fmt.Fprintf(outputfile, "%d,%s,%s\n", count, stu[i].Id, stu[i].Name)

			}
		}
	}
	fmt.Println("Process Finished!")
	return flag
}
