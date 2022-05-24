package pre

import (
	"fmt"
	"os"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func Init(filename string) *excelize.File {
	file, err := excelize.OpenFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return file
}

// 读取excel文件

func Load_data(file *excelize.File, str string) []string {
	sheetname := file.GetSheetName(file.GetActiveSheetIndex())
	// 默认获取第一个表单
	rows := file.GetRows(sheetname)

	var index int
	var storage []string
	for _, row := range rows {
		for tmp := range row {
			if row[tmp] == str {
				index = tmp
				break
			}
		}
	}
	// 记录索引位置

	for _, row := range rows {
		if row[index] != str {
			storage = append(storage, row[index])
		}
	}
	return storage
}

// 读取表格表单数据，返回特定索引下的列
