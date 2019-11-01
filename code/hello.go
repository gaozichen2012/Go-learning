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
