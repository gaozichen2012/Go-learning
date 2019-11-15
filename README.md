# Go语言学习笔记
本笔记仅适用于有C语言基础的人，由于Go语言与C语言相似度极高，我在做笔记时只记录易错易混淆和未掌握的知识点，信息密集度较高

## golang开发的优秀项目：
| 项目名 | 描述 |
| --- | --- |
| docker | golang头号优秀项目，通过虚拟化技术实现的操作系统与应用的隔离，也称为容器；<bar>无人不知的虚拟华平台，开源的应用容器引擎,借助该引擎，开发者可以打包他们的应用，移植到任何平台上。 | 
| kubernetes | 由google开发，简称k8s，k8s和docker是当前容器化技术的重要基础设施；<bar>Kubernetes基于Docker，其目的是让用户通过Kubernetes集群来进行云端容器集群的管理，而无需用户进行复杂的设置工作。系统会自动选取合适的工作节点来执行具体的容器集群调度处理工作。 |
| etcd | 一种可靠的分布式KV存储系统，有点类似于zookeeper，可用于快速的云配置；<bar>etcd是由CoreOS开发并维护键值存储系统，它使用Go语言编写，并通过Raft一致性算法处理日志复制以保证强一致性。目前，Google的容器集群管理系统Kubernetes、开源PaaS平台Cloud Foundry和CoreOS的Fleet都广泛使用了etcd。 |
| codis | 由国人开发提供的一套优秀的redis分布式解决方案； |
| tidb | 国内PingCAP 团队开发的一个分布式SQL 数据库，国内很多互联网公司在使用； |
| influxdb | 时序型DB，着力于高性能查询与存储时序型数据，常用于系统监控与金融领域； |
| cockroachdb | 云原生分布式数据库，继NoSQL之后出现的新的概念，称为NewSQL数据库； |
| beego | 国人开发的一款及其轻量级、高可伸缩性和高性能的web应用框架； |
| caddy | 类比于nginx，一款开源的，支持HTTP/2的 Web 服务端； |
| flynn | 一款开源的paas平台； |
| consul | HashiCorp公司推出的开源工具，用于实现分布式系统的服务发现与配置； |
| go-kit | Golang相关的微服务框架，这类框架还有go-micro、typthon； |
| go-ethereum | 官方开发的以太坊协议实现； |
| couchbase | 是一个非关系型数据库； |
| nsq | 一款高性能、高可用消息队列系统，每天能处理数十亿条的消息； |
| packer | 一款用来生成不同平台的镜像文件的工具，例如VM、vbox、AWS等； |
| doozer | 高速的分布式数据同步服务，类似ZooKeeper； |
| tsuru | 开源的PAAS平台，和SAE实现的功能一模一样； |
| gor | 一款用Go语言实现的简单的http流量复制工具； |
| lime | Lime，则是一款用Go语言写的桌面编辑器程序，被看做是著名编辑器Sublime Text的开源实现。 |

## 相关概念
* 适合Go做的比较成熟的软件开发方向主要包括服务器开发，云平台开发，微服务实践和重构，区块链开发（主要以以太坊为主导）等，特别是服务器开发，现在用go开展的很多创业公司团队在开始做，第二个是很多大厂的分布式系统，都是使用Go来构建的，这些具体的实践都是案例；
* 是Google开发的一种静态强类型、编译型、并发型，并具有垃圾回收功能的编程语言。
* Go的语法接近C语言，但对于变量的声明有所不同
* Go支持切片型slice、并发goroutine、通道channel、垃圾回收、接口interface等特性的语言级支持
* 通过使用Go内置gofmt工具后，会自动整理代码，使之符合规定的撰写风格。
* goroutine和channel是Go并发的两大基石。interface是Go语言编程中数据类型的关键，几乎所有的数据结构都围绕接口展开，接口是Go语言中所有数据结构的核心。

Go 语言被设计成一门应用于搭载 Web 服务器，存储集群或类似用途的巨型中央服务器的系统编程语言。
* 对于高性能分布式系统领域而言，Go 语言无疑比大多数其它语言有着更高的开发效率。它提供了海量并行的支持
* 简洁、快速、安全并行、有趣、开源内存管理、数组安全、编译迅速

## 相关命令
| 命令 | 解释 |
| --- | --- |
| go run | 执行go语言代码 |
| go build | 生成二进制文件 |

## 相关函数
| 函数名 | 描述 |
| --- | --- |
| delete() | 用于删除集合的元素,参数为map及其对应的key<br> 如定义一个集合`map1 := map[int]string{1:"t",2:"o",3:"m"}`，然后使用`delete(map1,2)`,会删除键值对`2:'0'`，剩下`{1:"t",3:"m"}`|
|  |  |

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
* 全局变量可以声明不使用，局部变量声明必须使用
* 交换2个变量值可以使用`a,b=b,a`，前提是2个变量的类型相同
* 常量声明用const，如`const a string = "abc"`或`const a = "abc"`或`const a,b = 1,2`
* 常量可以用len()、cap()、unsafe().Sizeof()函数计算表达式
* Go的算术运算符`+ - * / % ++ --`完全等同于c
* Go的关系运算符`== != > < >= <=`完全等同于c
* Go的逻辑运算符`&& || !`完全等同于c
* Go的位运算符`& | ^ << >>`完全等同于c
* Go的赋值运算符`= += -= *= /= %= <<= >>= &= ^= |=`完全等同于c
* Go的指针运算符`& *`完全等同于c，指针声明的格式为`var ptr *int`，然后再调用ptr
* 无限循环用`for true {}`，等同于c的`while(1) {}`
* Go函数可以有多个返回值，如`func get_value(x,y string) (string,string){return x,y}`，调用时`a,b=get_value(x,y)`
* 空指针的值为nil，等同于C语言的NULL
* 变量的数据类型转换，如将浮点数a转化为整数：`int(a)`

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
3. 省略 var, 如：`intVal,intVal1 := 1,2`，注意 := 左侧必须是没有声明的新的变量，否则会编译错误（局部变量声明首选）

### 枚举
```
const(
    Unknow = 0
    Female = 1
    Male = 2
)
```

### 特殊常量iota
iota可以认为是一个可以被编译器修改的常量
iota在const出现时被重置为0（const内部第一行之前），const中每新增一行常量声明，将使iota计数一次（iota可理解为const语句块的行索引）

iota可以被用作枚举值，如：
```
const(
    a = iota //第一个iota=0
    b = iota //第二个iota=1
    c = iota //第三个iota=2
)
```
以上iota用作枚举值可简化为一下代码
```
const(
    a=iota //每当iota在新一行被使用时会自加1，a=0,b=1,c=2
    b
    c
)
```

### 声明数组的几种形式
| 形式 | 描述 |
| --- | --- |
| 声明不赋值（默认） | `var buf [10] int` |
| 声明赋值 | `var buf=[10]int{1,2,3,4}` |
| 不写数组大小直接赋值 |`var buf=[...]int{1,2,3,4}`|

### 结构体
定义结构体需要type和struct语句，声明结构体有两种形式，如
```
package main

import "fmt"

type Books struct {
    title string
    author string
    book_id int
}

func main(){
    var Book1,Book2 Books //声明Book1，Book2为Books类型

    /* 声明一个新结构体并赋值 */
    Book3 := Books{"书book3","tom",66}

    /* 声明一个新结构体并使用key-value格式赋值 */
    Book4 := Books{title:"书book4",author:"tom",book_id:66}

    /* 声明一个新结构体并使用key-value格式部分赋值 */
    Book5 := Books{title:"书book5",book_id:66}

    Book1.title = "书book1"
    Book1.author = "tom"
    Book1.book_id = 66

    Book2.title = "书book1"
    Book2.author = "tom"
    Book2.book_id = 66

    fmt.Printf("Book1=%s\n",Book1.title)
    fmt.Printf("Book2=%s\n",Book2.title)
    fmt.Printf("Book3=%s\n",Book3.title)
    fmt.Printf("Book4=%s\n",Book4.title)
    fmt.Printf("Book5=%s\n",Book5.title)
}
```
输出得到：
Book1=书book1
Book2=书book1
Book3=书book3
Book4=书book4
Book5=书book5

### 切片（slice）（从部分未深入研究，用到时再记录笔记）
* Go提供一种内置类型：切片，切片是对数组的抽象，等同于“动态数组”，长度不固定，可以追加元素，在追加时可能使切片容量增大
* 声明切片不需要说明长度，如`var slice1 []int`
* 也可以使用make()函数创建切片，如：`slice1 := make([]type,len,capacity)`，这里len是切片的初始长度 ,capacity为可选参数，作用是指定容量
* 初始化切片,如：`s := []int{1,2,3}`,[]表示切片类型，cap=len=3
* ...

### 范围（range）（从部分未深入研究，用到时再记录笔记）
* 关键字用于for循环中迭代数组（）、切片（）、通道（）、或集合（）的元素，在数字和切片中返回元素的索引和索引对应的值，在集合中返回key-value对

### 集合（map）（从部分未深入研究，用到时再记录笔记）
* Map 是一种无序的键值对的集合。Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值。
* Map 是一种集合，所以我们可以像迭代数组和切片那样迭代它。不过，Map 是无序的，我们无法决定它的返回顺序，这是因为 Map 是使用 hash 表来实现的。
* 定义map可以使用内建函数make，如：`var map1 map[string]string`,也可以使用map关键字定义map，如：`map1 := make(map[string]string)`
* 如果不初始化map，会创建nil map，nil map不能存放键值对

### 递归函数
递归就是在运行过程中调用自己，格式如下
```
func recursion(){
    recursion{}
}

func main(){
    recursion()
}
```
在使用递归时，必须设置退出条件，否则将陷入无限循环

### 数据类型：接口（interface）（从部分未深入研究，用到时再记录笔记）
Go提供另外一种数据类型即接口，它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口。

### 错误处理（从部分未深入研究，用到时再记录笔记）
Go 语言通过内置的错误接口提供了非常简单的错误处理机制。
error类型是一个接口类型，以下是它的定义：
```
type error interface {
    Error() string
}
```
我们可以在编码中通过实现 error 接口类型来生成错误信息。函数通常在最后的返回值中返回错误信息。使用errors.New 可返回一个错误信息：
```
func Sqrt(f float64) (float64, error) {
    if f < 0 {
        return 0, errors.New("math: square root of negative number")
    }
    // 实现
}
```
### 并发（goroutine）
* Go 语言支持并发，我们只需要通过 go 关键字来开启一个新的运行期线程goroutine,以一个不同的、新创建的goroutine来执行一个函数。
* 同一个程序中的所有 goroutine 共享同一个地址空间。
* goroutine 是轻量级线程，goroutine 的调度是由 Golang 运行时进行管理,格式如下

```
package main

import (
    "fmt"
    "time"
)

func say(s string){
    for i := 0;i < 5; i++{
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s)
    }
}

/* 开启4个线程，同时输出tom1,tom2,tom3,tom4,而tom5，tom6是等线程运行完再逐个输出 */
func main(){
    go say("tom1") //开启第1个say线程
    go say("tom2") //开启第2个say线程
    go say("tom3") //开启第3个say线程
    say("tom4")
    say("tom5")
    say("tom6")
}
```

### 通道（channel）
channel是用来传递数据的有一个数据结构
通道可用于两个goroutine之间通过一个指定类型的值来同步运行和通讯。操作符<-用于指向通道的方向，发送或接收，如果方向未指定，则为双向通道
```
ch <- v //把V发送到通道ch
v := <-ch // 从ch接收数据，并把值赋给v
```
定义一个通道使用chan关键字,如`ch := make(chan int)`
注意：默认情况下，通道是不带缓冲区的。发送端发送数据，同时必须又接收端相应的接收数据。

以下例子使用
```
package main

import "fmt"

func sum(s []int, c chan int) {
        sum := 0
        for _, v := range s {
                sum += v
                fmt.Printf("sum=%d\n",sum)
        }
        c <- sum // 把 sum 发送到通道 c
}

/* 通过两个 goroutine 来计算数字之和，在 goroutine 完成计算后，它会计算两个结果的和 */
func main() {
        s := []int{7, 2, 8, -9, 4, 0}

        c := make(chan int)
        go sum(s[:len(s)/2], c)
        go sum(s[len(s)/2:], c)
        x, y := <-c, <-c // 从通道 c 中接收

        fmt.Println(x, y, x+y)
}
```

## 下次学习
1. 理解通道channel在goroutine的作用及用法
2. 学习通道缓冲区
3. 遍历通道与关闭通道
4. 学习几个跳过的知识点，比如range等
5. 找其他资料学习Go的语法，确定是不是所有语法都包含
6. 用Go实现几个项目