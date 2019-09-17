/*
@project: GoLand
@author: fke
@file: 语言函数.go
@time: 2019-09-12 17:31:50
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
func main() {
    // 定义局部变量
    var a int = 300
    var b int = 200
    var ret int

    // 调用函数并返回最大值
    ret = max(a, b)
    fmt.Printf("最大值为 ：%d\n", ret)
    fmt.Println("ret的值类型为:", reflect.TypeOf(ret))
}

func max(num1, num2 int) int {
    // 定义局部变量
    var result int

    if (num1 > num2) {
        result = num1
    } else {
        result = num2
    }
    return result
}
*/




/*
// 函数返回多个值
func swap(x, y string) (string, string) {
    return y, x
}

func main() {
    a, b := swap("Google", "fke")
    fmt.Println(a, b)
}
*/




/*
// 语言函数值传递值
func main()  {
    // 定义局部变量
    var a int = 100
    var b int = 200

    fmt.Printf("交换前 a 得值为：%d\n", a)
    fmt.Printf("交换前 b 得值为：%d\n", b)

    // 通过调用函数来交换
    swap(a, b)

    fmt.Printf("交换后 a 的值：%d\n", a)
    fmt.Printf("交换后 b 的值：%d\n", b)
}

// 定义相互交换值的函数
func swap(x, y int) int {
    var temp int

    temp = x // 先保存x的值
    x = y   // 将y的值赋给x
    y = temp    // 将temp值赋给y

    return temp;
}
*/




/*
// 语言函数引用传递值
// 通过引用传递来调用swap()函数
func main()  {
    // 定义局部变量
    var a int = 100
    var b int = 200

    fmt.Printf("交换前 a 的值：%d\n", a)
    fmt.Printf("交换前 b 的值：%d\n", b)

    // 调用swap()函数;  &a 指向a指针, a变量的地址;   &b 指向b指针, b变量的地址
    swap(&a, &b)
    fmt.Printf("交换后 a 的值：%d\n", a)
    fmt.Printf("交换后 b 的值：%d\n", b)
}

func swap(x *int, y *int)  {
    var temp int
    temp = *x
    *x = *y
    *y = temp
}
*/





/*
// 语言函数作为实参
import "math"

func main() {
    // 声明函数变量
    getSquareRoot := func(x float64) float64 {
        return math.Sqrt(x)
    }

    // 使用函数
    fmt.Println(getSquareRoot(9))
}
*/





/*
// 语言函数闭包
func getSequence() func() int {
    i:= 0
    return func() int {
        i += 1
        return i
    }
}

func main() {
    // nextNumber 为一个函数，函数i为0
    nextNumber := getSequence()

    // 调用nextNumber 函数，i变量自增1并返回
    fmt.Println(nextNumber())
    fmt.Println(nextNumber())
    fmt.Println(nextNumber())

    // 创建新的函数nextNumber1, 并查看结果
    nextNumber1 := getSequence()
    fmt.Println(nextNumber1())
    fmt.Println(nextNumber1())
}
*/




/*
// eg：带参数的闭包函数调用
func main() {
    add_func := add(1, 2)
    fmt.Println(add_func(1, 1))
    fmt.Println(add_func(0, 0))
    fmt.Println(add_func(2, 2))
}

func add(x1, x2 int) func(x3 int, x4 int)(int, int, int) {
    i := 0
    return func(x3 int, x4 int) (int, int, int) {
        i++
        return i, x1+x2, x3+x4
    }
}
*/





// 语言函数方法
// 定义结构体
type Circle struct {
    radius float64
}

func main() {
    var c1 Circle
    c1.radius = 10.00
    fmt.Println("圆的面积=", c1.getArea())
}

// 该method属于Circle 类型对象中的方法
func (c Circle) getArea() float64 {
    // c.radius 即为 Circle类型对象中的属性
    return 3.14 * c.radius * c.radius
}
