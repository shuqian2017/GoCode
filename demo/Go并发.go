/*
@project: GoLand
@author: fke
@file: Go并发.go
@time: 2019-09-19 15:56:38
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
import "fmt"


/*
// Go允许使用go语句开启一个新的运行期线程，即goroutine来执行一个函数。同一个程序中所有的goroutine共享一个地址空间
import "time"
func say(s string)  {
    for i := 0; i < 5; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s)
    }
}

func main()  {
    go say("world")
    say("hello")
}
// 执行上面的代码，"hello","world"没有固定的先后顺序，因为他们是两个goroutine在执行
*/




/*
// 通道（channel） 关键字chan 使用前必须先创建
func sum(s []int, c chan int) {
    sum := 0
    for _, v := range s {
        sum += v
    }
    c <- sum  // 把sum发送到通道c
}

func main()  {
    s := []int{7,2,8,-9,4,0}

    c := make(chan int)
    go sum(s[:len(s)/2], c)
    go sum(s[len(s)/2:], c)
    x, y := <-c, <-c   // 从通道中接收
    fmt.Println(x, y, x+y)
}
*/




/*
//通道可以设置缓冲区，通过make的第二个参数指定缓冲区大小
func main()  {
    // 设置缓冲去区的大小为10
    ch := make(chan int, 10)

    // 因为ch是带缓冲的通道，所以我们可以同时发送多个数据
    ch <- 1
    ch <- 2
    ch <- 3
    ch <- 4

    // 获取这几个数据
    fmt.Println(<-ch)
    fmt.Println(<-ch)
    fmt.Println(<-ch)
    fmt.Println(<-ch)
}
*/





//Go遍历通道与关闭通道； 通过range关键字来实现遍历读取到的数据
func fibonacci(n int, c chan int) {
    x, y := 0, 1
    for i := 0; i < n; i++ {
        c <- x
        x, y = y, x+y
    }
    close(c)
}

func main()  {
    c := make(chan int, 10)
    go fibonacci(cap(c), c)

        // range进行遍历每个从通道接收的数据，因为c 发送完10个后就关闭了通道，所以range也在接收到10个后就结束了，因此在接收第11个数据的时候会出现阻塞
    for i := range c {
        fmt.Println(i)
    }
}

