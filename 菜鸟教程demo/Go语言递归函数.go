/*
@project: GoLand
@author: fke
@file: Go语言递归函数.go
@time: 2019-09-19 14:26:11
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
// 递归就是在运行过程中，自己调用自己
// 递归函数实现阶乘
func main() {
    var i int = 15
    fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(uint64(i)))
}

func Factorial(n uint64)(result uint64) {
    if (n > 0) {
        result = n * Factorial(n-1)
        return result
    }
    return 1
}
*/




//斐波那契数列
func fibonacci(n int) int {
    if n < 2 {
        return n
    }
    return fibonacci(n-2) + fibonacci(n-1)
}

func main()  {
    var i int
    for i = 0; i < 10; i++ {
        fmt.Printf("%d\t", fibonacci(i))
    }
}
