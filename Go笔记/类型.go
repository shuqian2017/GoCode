/*
@project: GoLand
@author: fke
@file: 类型.go
@time: 2020-05-19 22:10:04
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
// 2.1 变量

// 简短模式     简短定义在函数多返回值,以及if/for/switch等语句中定义局部变量非常方便
//var x = 100
//
//func main() {
//    println(&x, x)              // 全局变量
//
//    x := "abc"                  // 重新定义和初始化同名局部变量
//    println(&x, x)
//    // 对比运行结果的内存地址， 可以看出是两个不同的变量
//}


//func main() {
//    x := 100
//    println(&x)
//
//    x, y := 200, "abc"              // 注意: x 退化为赋值操作,仅有y是变量定义
//
//    println(&x, x)
//    println(y)
//    // 对比变量的内存地址,可以确认x属于同一变量
//}

//func main() {
//    x := 100
//    println(&x)
//
//    x := 200                        // 这种就是错误的： no new variables on left side of :=
//    println(&x, x)
//}

// 退化赋值的前提条件是： 最少有一个新变量被定义，且必须是同一作用域
//func main() {
//    x := 100
//    println(&x, x)
//
//    {
//        x, y := 200, 300            // 这种作用域不同，全部都是新变量的定义
//        println(&x, x, y)
//    }
//}


// 在处理函数错误返回值时， 退化赋值允许我们重复使用err变量，这是相当有益的
//import (
//    "fmt"
//    "os"
//)
//
//func main() {
//    f, err := os.Open("/dev/random")
//
//    buf := make([]byte, 1024)
//    n, err := f.Read(buf)               // err退化赋值, n 新定义
//    // 伪代码
//}


// 多变量赋值     在进行多变量赋值操作时，首先计算出所有右值，然后再依次完成赋值操作
//func main() {
//    x, y := 1, 2
//    x, y = y+3, x+2                 // 先计算出右值y+3, x+2, 然后再对x、y变量赋值
//
//    println(x, y)
//}


// 未使用错误     编译器将未使用的局部变量当作错误
//var x int           // 全局变量未使用没有问题
//func main() {
//    y := 10         // 错误： y declared and not used
//}
*/






//*********************************************************************************************






/*
// 2.2 命名

//// 命名建议
//* 以字母或下划线开始, 由多个字母, 数字和下划线组合而成
//* 区分大小写
//* 使用驼峰(camel case)拼写格式
//* 局部变量优先使用短名
//* 不要使用保留的关键字
//* 不建议使用与预定义常量, 类型, 内置函数相同的名字
//* 专有名词通常会全部大写, 例如escapeHTML

// 空标识符     和python类似, 有个“_”特殊成员，通常作为忽略占位符使用，可做表达式左值，无法读取内容
//import (
//    "strconv"
//)
//
//func main() {
//    x, _ := strconv.Atoi("12")              // 忽略Atoi的err返回值
//    println(x)
//    // 空标识符可以用来临时规避编译器对未使用的变量和导入包的错误检查。它是预置成员，不能重新定义
//}
*/






//*********************************************************************************************







// 2.3 常量
//func main() {
//    const x = 123
//    println(x)
//
//    const y = 1.23              // 常量未使用，不会引发编译器报错
//
//    {
//        const x = "abc"         // 在不同作用域内的同名常量
//        println(x)
//    }
//}


// 常量值也可以是某些编译器能计算出结果的表达式，如unsafe, Sizeof, len, cap
//import (
//    "unsafe"
//)
//
//const (
//    ptrSize = unsafe.Sizeof(uintptr(0))
//    strSize = len("hello, world!")
//)
//
//func main() {
//    println(ptrSize, strSize)
//}


// 在常量组中如不指定类型和初始化值，则与上一行非空常量右值相同
//import (
//    "fmt"
//)
//
//func main() {
//    const (
//        x uint16 = 120
//        y
//        s = "abc"
//        z
//    )
//
//    fmt.Printf("%T, %v\n", y, y)
//    fmt.Printf("%T, %v\n", z, z)
//}




// 枚举     Go并没有明确意义上的enum定义，不过可以借助iota标识符实现一组自增常量值来实现枚举类型
//const (
//    x = iota        // 0
//    y               // 1
//    z               // 2
//)
//
//const (
//    _ = iota        // 0
//    KB = 1 << (10 * iota)       // 1 << (10 * 1)
//    MB                          // 1 << (10 * 2)
//    GB                          // 1 << (10 * 3)
//)
//
//const (                         // 多常量定义中可以使用多个iota，它们各自单独计数
//    _, _ = iota, iota * 10      // 0, 0 * 10
//    A, B                        // 1, 1 * 10
//    C, D                        // 2, 2 * 10
//)
//
//const (                         // 如中断iota,则必须显式恢复。且后续自增值按行序递增， 而非enum那般按上一行取值递增
//    a = iota                    // 0
//    b                           // 1
//    c = 100                     // 100
//    d                           // 100  (与上一行常量右值表达式相同)
//    e = iota                    // 4    (恢复iota自增， 计数包括c, d)
//    f                           // 5
//)
//
//const (                         // 自增默认数据类型为int, 可显式指定数据类型
//    a1 = iota                   // int
//    b1 float32 = iota           // float32
//    c1 = iota                   // int (如不显式指定iota, 则与b数据类型相同)
//)



//type color byte                 // 自定义类型
//
//const (
//    black color = iota          // 指定常量类型
//    red
//    blue
//)
//
//func test(c color) {
//    println(c)
//}
//
//func main() {
//    test(red)
//    test(100)                   // 100 并没有超出 color/byte 类型的取值范围
//
//    x := 2
//    test(x)                     // cannot use x (type int) as type color in argument to test
//}


