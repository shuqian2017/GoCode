/*
@project: GoLand
@author: fke
@file: Go语言结构体.go
@time: 2019-09-17 15:31:55
 code is far away from bugs with the god animal protecting
    I love animals. They taste delicious.
              ┏┓      ┏┓
            ┏┛┻━━━┛┻┓
            ┃      ☃      ┃
            ┃  ┳┛  ┗┳  ┃
            ┃      ┻      ┃
            ┗━┓      ┏━┛
                ┃      ┗━━━┓
                ┃  神兽保佑    ┣┓
                ┃　永无BUG。   ┏┛
                ┗┓┓┏━┳┓┏┛
                  ┃┫┫  ┃┫┫
                  ┗┻┛  ┗┻┛
*/
package main

import(
    "fmt"
)


/*
//结构体定义需要使用 type和struct语句，struct语句定义一个新的数据类型，结构体中可以有一个或者多个成员
type Books struct {
    title string
    author string
    subject string
    book_id int
}
func main() {
    // 创建一个新的结构体
    fmt.Println(Books{"Go语言", "www.runoob.com", "Go语言教程", 12580})

    // 页1可以使用key => value格式
    fmt.Println(Books{title:"Go语言", author:"www.runoob.com", subject:"Go语言教程", book_id:12580})

    // 忽略的字段为0或空
    fmt.Println(Books{title:"Go语言", author:"www.runoob.com"})
}
*/



/*
//访问结构体成员; 需要使用点号(.)操作符
type Books struct {
    title string
    author string
    subject string
    book_id int
}

func main()  {
    var Book1 Books
    var Book2 Books

    // book1 描述
    Book1.title = "Go 语言"
    Book1.author = "www.runoob.com"
    Book1.subject = "Go 语言教程"
    Book1.book_id = 12580

    // book2 描述
    Book2.title = "Python 教程"
    Book2.author = "www.runoob.com"
    Book2.subject = "Python 语言教程"
    Book2.book_id = 321585

    // 打印book1 信息
    fmt.Println("Book 1 title: ", Book1.title,
                "\nBook 1 author: ", Book1.author,
                "\nBook 1 subject: ", Book1.subject,
                "\nBook 1 book_id: ", Book1.book_id,
                "\n")

    // 打印book1 信息
    fmt.Println("Book 2 title: ", Book2.title,
                "\nBook 2 author: ", Book2.author,
                "\nBook 2 subject: ", Book2.subject,
                "\nBook 2 book_id: ", Book2.book_id)
}
*/




/*
//结构体作为函数参数
type Books struct {
    title string
    author string
    subject string
    book_id int
}

func main()  {
    var Book1 Books
    var Book2 Books

    // book1 描述
    Book1.title = "Go 语言"
    Book1.author = "www.runoob.com"
    Book1.subject = "Go 语言教程"
    Book1.book_id = 12580

    // book2 描述
    Book2.title = "Python 教程"
    Book2.author = "www.runoob.com"
    Book2.subject = "Python 语言教程"
    Book2.book_id = 321585

    // 打印Book1 信息
    printBook(Book1)

    // 打印Book2 信息
    printBook(Book2)
}

func printBook( book Books) {
    fmt.Println("Book title:", book.title,
        "\nBook author:", book.author,
        "\nBook subject:", book.subject,
        "\nBook book_id:", book.book_id,
        "\n")
}
*/





//结构体指针
type Books struct {
    title string
    author string
    subject string
    book_id int
}

func main()  {
    var Book1 Books
    var Book2 Books

    // book1 描述
    Book1.title = "Go 语言"
    Book1.author = "www.runoob.com"
    Book1.subject = "Go 语言教程"
    Book1.book_id = 12580

    // book2 描述
    Book2.title = "Python 教程"
    Book2.author = "www.runoob.com"
    Book2.subject = "Python 语言教程"
    Book2.book_id = 321585

    // 打印Book1 信息
    printBook(&Book1)
    // 打印Book2 信息
    printBook(&Book2)
}

func printBook(book *Books) {
    fmt.Println("Book title:", book.title,
        "\nBook author:", book.author,
        "\nBook subject:", book.subject,
        "\nBook book_id", book.book_id,
        "\n")
}