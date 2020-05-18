/*
@project: GoLand
@author: fke
@file: 语言变量作用域.go
@time: 2019-09-16 15:23:06
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

import (
    "fmt"
)

/*
// 局部变量
func main() {
    // 声明局部变量
    var a, b, c int

    // 初始化参数
    a = 10
    b = 20
    c = a+ b
    fmt.Printf("结果：a = %d, b = %d and  c = %d\n", a, b, c)
}
*/




/*
// 全局变量
var g int

func main()  {
    // 声明局部变量
    var a, b int

    // 初始化参数
    a = 10
    b = 20
    g = a + b
    fmt.Printf("结果：a = %d, b = %d and  g = %d\n", a, b, g)
}
// Go语言中全局变量和局部变量名称可以相同，但是函数内的局部变量会被优先考虑
*/




/*
// 形式参数
// 先声明全局变量
var a int = 20;

func main() {
    // main 函数中声明局部变量
    var a int =10
    var b int =20
    var c int =0

    fmt.Printf("main()函数中 a = %d\n", a);
    c = sum(a, b);
    fmt.Printf("main()函数中 c = %d\n", c);
}

// 函数定义-两数相加
func sum(a, b int) int {
    fmt.Printf("sum()函数中 a = %d\n", a);
    fmt.Printf("sum()函数中 b = %d\n", b);
    return a + b;
}
*/



//eg: 形参使用，比较sum函数中的a和main函数中的a，sum函数中虽然加了1，但是main中还是原来值10
// 声明全局变量
var a int = 20

func main() {
    // main 函数中声明的局部变量
    var a int = 10
    var b int = 20
    var c int = 0

    fmt.Printf("main()函数中 a = %d\n", a)
    c = sum(a, b)
    fmt.Printf("main()函数中 a = %d\n", a)
    fmt.Printf("main()函数中 c = %d\n", c)
}

// 函数定义-两数相加
func sum(a, b int) int {
    a += 1
    fmt.Printf("sum()函数中 a = %d\n", a)
    fmt.Printf("sum()函数中 b = %d\n", b)
    return a + b
}
