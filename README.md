# ibit-helper

i北理疫情防控打卡 表格比对

### 更新

* 增加GUI

### 

### 编译

```bash
go mod tidy
go build main.go
```

### 运行

#### GUI运行

* Windows：双击`.exe`文件

* Linux：暂未测试



#### 命令行运行

* Linux:

```bash
./main  已完成核酸名单.xlsx  全体学生名单.xlsx
```

* Windows:

```bash
main.exe 已完成核酸名单.xlsx  全体学生名单.xlsx
```

或者将对应文件改名为`已完成核酸名单.xlsx`，` 全体学生名单.xlsx`，与`main.exe`和`windows.bat`放置于同一目录下，运行`windows.bat`
