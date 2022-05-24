# ibit-helper

i北理疫情防控打卡 表格比对

### 更新

* 增加GUI

<img src="https://tva1.sinaimg.cn/large/0084b03xly1h2jnzk0aqoj30ln0modic.jpg" title="" alt="1653387638(1).jpg" width="348">

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

### 注意事项

* 已打卡文件中，学号索引为“学工号”

* 全体学生文件中，学号索引为“学号”
