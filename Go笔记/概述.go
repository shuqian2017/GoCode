/*
@project: GoLand
@author: fke
@file: 概述.go
@time: 2020-05-18 23:13:44
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

/*
// 变量
func main() {
    var x int32
    var s = "hello, world!"
    X := 100   // 在函数内部还可以省略var关键字， 使用:=

    println(x, s, X)
}

*/


//*********************************************************************************************


/*
//表达式    if / switch / for
//func main() {
//    x := 100
//
//    if x > 0 {
//        println("x")
//    } else if x < 0 {
//        println("-x")
//    } else {
//        println("0")
//    }
//}


//func main() {
//    x := 100
//
//    switch {
//    case x > 0:
//        println("x")
//    case x < 0:
//        println("-x")
//    default:
//        println("0")
//    }
//}



//func main() {
//   for i := 0; i < 5; i++ {
//       println("①", i)
//   }
//
//   for i := 4; i >= 0; i-- {
//       println("②", i)
//   }
//}



//func main() {
//    x := 0
//
//    for x < 5 {   // 相当于while (x < 5) {...}
//        println(x)
//        x++
//    }
//}


//func main() {
//    x := 4
//
//    for {       // 相当于while (true) {...}
//        println(x)
//        x--
//
//        if x < 0 {
//            break
//        }
//    }
//}


//func main() {
//    x := []int{100, 101, 102}
//
//    // 在迭代遍历时, for...range除元素外，还可返回索引
//    for i, n := range x {
//        println(i, ":", n)
//    }
//}
*/




//*********************************************************************************************




/*
//函数 函数可定义多个返回值，甚至对其命名
//import (
//    "errors"
//    "fmt"
//)
//
//func div(a, b int) (int, error) {
//    if b == 0 {
//        return 0, errors.New("division by zero")
//    }
//
//    return a / b, nil
//}
//
//func main()  {
//    a, b := 10, 2       // 定义多个变量
//    c, err := div(a, b)     // 接收多返回值
//
//    fmt.Println(c, err)
//}


// 函数是第一类型，可作为参数或者返回值
//func test(x int) func() {       // 返回函数类型
//    return func() {             // 匿名函数
//        println(x)              // 闭包
//    }
//}
//
//func main() {
//    x := 100
//
//    f := test(x)
//    f()
//}


// 用defer定义延迟调用, 无论函数是否出错，它都确保结束前被调用
//func test(a, b int) {
//    defer println("① dispose...")       // 常用来释放资源,解除锁定,或执行一些清理操作
//    defer println("② dispose...")       // 可以定义多个defer， 按FILO顺序执行
//
//    println(a / b)
//
//}
//
//func main() {
//    test(10, 0)
//}
*/





//*********************************************************************************************






/*
// 数据    切片slice, 字典map, 结构体struct
//切片slice可实现类似动态数组的功能
//import (
//    "fmt"
//)
//
//func main() {
//    x := make([]int, 0, 5)      // 创建一个容量为5的切片
//    fmt.Println("①",x)
//
//    for i := 0; i < 8; i++ {
//        x = append(x, i)        // 追加数据, 当超出切片容量限制后，自动分配更大的存储空间
//    }
//
//    fmt.Println("②", x)
//}



// 将字典map类型内置, 可直接从运行时层面获取性能优化
//import (
//    "fmt"
//)
//
//func main() {
//    m := make(map[string]int)       // 创建字典类型对象
//
//    m["a"] = 1                      // 添加或设置
//
//    x, ok := m["a"]                 // 使用 ok-idiom获取值, 可知key/value是否存在 (所谓ok-idiom模式，是指在多返回值中用一个名为ok的布尔值来标示操作是否成功，因为很多操作默认返回0值，所以需额外说明)
//    fmt.Println(x, ok)
//
//    delete(m, "a")              // 删除
//
//    X, ok := m["a"]                 // 使用 ok-idiom获取值, 可知key/value是否存在
//    fmt.Println(X, ok)
//}



// 结构体struct可匿名嵌入其他类型
//import (
//    "fmt"
//)
//
//type user struct {
//    name string             // 结构体类型
//    age byte
//}
//
//type manager struct {
//    user                    // 匿名嵌入其他类型
//    title string
//}
//
//func main() {
//    var m manager
//
//    m.name = "fke"          // 直接访问匿名字段的类型
//    m.age = 18
//    m.title = "CEO"
//
//    fmt.Println(m)
//}
*/








//*********************************************************************************************







/*
// 方法   可以为当前包内的任意类型定义方法
//type X int
//func (x *X) inc() {         // 名称前的参数称作receiver, 作用类似python self
//    *X++
//}
//
//func main() {
//    var x X
//    x.inc()
//    println(x)
//}



// 还可以直接调用匿名字段的方法，这种方式可实现与继承类似的功能
//import (
//    "fmt"
//)
//
//type user struct {
//    name string
//    age byte
//}
//
//func (u user) ToString() string {
//    return fmt.Sprintf("%+v", u)          // %+v在 %v 基础上，对结构体字段名和值进行展开
//}
//
//type manager struct {
//    user
//    title string
//}
//
//func main() {
//    var m manager
//    m.name = "fke"
//    m.age = 18
//
//    println(m.ToString())           // 调用user.ToString()
//}


动 词	功 能
%v	按值的本来值输出
%+v	在 %v 基础上，对结构体字段名和值进行展开
%#v	输出 Go 语言语法格式的值
%T	输出 Go 语言语法格式的类型和值
%%	输出 % 本体
%b	整型以二进制方式显示
%o	整型以八进制方式显示
%d	整型以十进制方式显示
%x	整型以十六进制方式显示
%X	整型以十六进制、字母大写方式显示
%U	Unicode 字符
%c  输出单个字符
%f	浮点数
%p	指针，十六进制方式显示
*/





//*********************************************************************************************





/*
// 接口    接口采用duck type方式，也就是说无须在实现类型上添加显式声明
//import (
//    "fmt"
//)
//
//type user struct {
//    name string
//    age byte
//}
//
//func (u user) Print() {
//    fmt.Printf("%+v\n", u)
//}
//
//type Printer interface {            // 接口类型
//    Print()
//}
//
//func main() {
//    var u user
//    u.name = "fke"
//    u.age = 18
//
//    var p Printer = u               // 只要包含接口所需的全部方法，即表示实现了该接口
//    p.Print()
//    // 另有空接口类型interface{}, 用途类似OOP里的system.Object, 可接收任意类型对象
//}
*/






//*********************************************************************************************







// 并发   goroutine方式，这是一种比普通协程或线程更加高效的并发设计，轻松创建和运行成千上万的并发任务
//import (
//    "fmt"
//    "time"
//)
//
//func task(id int) {
//    for i := 0; i < 5; i++ {
//        fmt.Printf("%d: %d\n", id, i)
//        time.Sleep(time.Second)
//    }
//}
//
//func main() {
//    go task(1)              // 创建goroutine
//    go task(2)
//
//    time.Sleep(time.Second * 6)
//}



// 通道channel 与goroutine搭配，实现用通信代替内存共享的CSP模型
// 消费者
func consumer(data chan int, done chan bool)  {
    for x := range data {                       // 接收数据，直到通道被关闭
        println("recv:", x)
    }

    done <- true                                // 通知main， 消费结束
}

// 生产者
func producer(data chan int)  {
    for i := 0; i < 8; i++ {
        data <- i                               // 发送数据
    }

    close(data)                                 // 生产结束，关闭通道
}

func main() {
    done := make(chan bool)                     // 用于接收消费结束信号
    data := make(chan int)                      // 数据管道

    go consumer(data, done)                     // 启动消费者
    go producer(data)                           // 启动生产者

    <- done                                     // 阻塞，直到消费者发回结束信号
}

