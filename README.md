# Go语言学习笔记
[toc]

## 相关概念

* Go 语言被设计成一门应用于搭载 Web 服务器，存储集群或类似用途的巨型中央服务器的系统编程语言。
* 对于高性能分布式系统领域而言，Go 语言无疑比大多数其它语言有着更高的开发效率。它提供了海量并行的支持
* 简洁、快速、安全并行、有趣、开源内存管理、数组安全、编译迅速

## 相关命令
| 命令 | 解释 |
| --- | --- |
| go run | 执行go语言代码 |
| go build | 生成二进制文件 |

## 相关语法

* Go 语言的基础组成部分：包声明、引入包、函数、变量、语句 & 表达式、注释
* 每一个go应用程序都有一个main的包，`package main`表示一个可独立执行的程序
* 必须在源文件中第一行指明这个文件属于哪个包，如：package main
* 程序启动时最先执行main函数，如果有init()函数则先执行
* `Println("hello tom")`等同于`Print(hello tom\n)`
* 当标识符（包括常量、变量、类型、函数名、结构字段等等）以一个大写字母开头，如：Group1，那么该对象就可以被外部包的代码所使用（客户端程序需要先导入这个包）
* 标识符如果以小写字母开头，则对包外是不可见的，但是他们在整个包的内部是可见并且可用的
* 括号 { 不能单独放在一行
* 一行代表一个语句结束，不需要向C那样以分号结尾

### 字符串连接
go语言的字符串可以通过+实现，如`fmt.Println("tom"+"is"+"dog")`，输出`tomisdog`

### 数据类型
| 类型 | 描述 |
| --- | --- |
| 布尔型 bool | 布尔型的值只可以是常量 true 或者 false。如：`var b bool = true` |
| 数字类型 | 整数型int、uint8、uint16、uint32、uint64、int8、int16、int32、int64 <br>浮点型 float32、float64、complex64、complex128<br>其他类型：byte、rune、uint、int、uintptr|
| 字符串类型 | Go语言的字符串使用UTF-8编码表示Unicode文本 |
| 派生型 | 包括指针类型Pointer、数组类型、结构化类型struct、Channel类型、函数类型、切片类型、接口类型interface、Map类型 |

### 变量
* Go 语言中变量的声明必须使用空格隔开，如：`var age int;`
* 声明变量的一般形式是使用 var 关键字：`var a string = "Runoob"`,也可以一次声明多个变量：`var a,b int =1,2`
#### 变量声明的三种形式
1. 指定变量类型，如果没有初始化，默认0值（最常用）
2. 根据值自行判断变量类型，如：`var a = true`
3. 省略 var, 如：`intVal,intVal1 := 1,2`，注意 := 左侧必须是没有声明的新的变量，否则会编译错误

# 接下来学习值类型和引用类型