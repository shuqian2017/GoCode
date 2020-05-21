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






/*
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




// 展开
//Q： 常量除"只读"外, 和变量究竟有什么不同？

//var x = 0x100
//const y = 0x200
//
//func main() {
//   println(&x, x)
//   println(y)
//   //println(&y, y)                 // cannot take the address of y
//   // 数字常量不会分配存储空间，无需像变量那样通过内存寻址来取值，因此无法获取地址
//}

//A：不同于变量在运行期分配存储内存, 常量通常会被编译器在预处理阶段直接展开，作为指令数据使用

//import (
//   "fmt"
//)
//
//const x = 100                   // 无类型声明的常量
//const y byte = x                // 直接展开x, 相当于const y byte = 100
//
//const a int = 100               // 显示指定常量类型,编译器会做强类型检查
////const b byte = a                // 错误：cannot use a (type int) as type byte in const initializer
//
//func main() {
//   fmt.Printf("%T %v\n", x, x)
//   fmt.Printf("%T %d\n", y, y)
//   fmt.Printf("%T %v\n", a, a)
//}
*/






//*********************************************************************************************









// 2.4 基本类型     array(数组)，struct(结构体)，function(函数)，interface(接口)，map(字典)，slice(切片)，channel(通道)
//import (
//    "fmt"
//    "math"
//)
//
//func main() {
//    a, b, c := 100, 0144, 0x64
//
//    fmt.Println(a, b, c)
//    fmt.Printf("0b%b, %#o, %#x\n", a, b, c)
//
//    fmt.Println(math.MinInt8, math.MaxInt8)
//    // 10 -> 2 除2取余
//    // 2 -> 10  乘2的n次方，再相加
//    // 2 -> 8   按三位来进行划分
//    // 2 -> 16  按4位来进行划分
//}



//import (
//    "strconv"
//)
//
//func main() {
//    a, _ := strconv.ParseInt("1100100", 2, 32)
//    b, _ := strconv.ParseInt("0144", 8, 32)
//    c, _ := strconv.ParseInt("64", 16, 32)
//
//    println(a, b, c)
//
//    println("0b" + strconv.FormatInt(a, 2))
//    println("0" + strconv.FormatInt(a, 8))
//    println("0x" + strconv.FormatInt(a, 16))
//    // 标准库strconv 可在不同进制(字符串)间转换
//}



//func main() {
//    var a float32 = 1.1234567899        // 注意： 默认浮点类型是float64
//    var b float32 = 1.12345678
//    var c float32 = 1.123456781
//
//    println(a, b, c)
//    println(a == b, a == c)
//    fmt.Printf("%v %v, %v\n", a, b, c)
//    // 使用浮点数时，须注意小数位的有效精度，相关细节可参考IEEE-754标准
//}




// 别名   byte -> alias for uint8;    rune -> alias for int32
//func test(x byte) {
//    println(x)
//}
//
//func main() {
//    var a byte = 0x11       // 16进制0x11转化为10进制的a = 17
//    var b uint8 = a
//    var c uint8 = a + b
//
//    test(c)
//}



//func add(x, y int) int {
//    return x + y
//}
//
//func main() {
//    var x int = 100
//    var y int64 = x     //错误： cannot use x (type int) as type int64 in assignment
//
//    add(x, y)           //错误： cannot use y (type int64) as type int in argument to add
//}






//*********************************************************************************************








// 2.5 引用类型     所谓引用类型(reference type)特指slice map channel ; 引用类型必须使用make函数创建
//func mkslice() []int {
//    s := make([]int, 0, 10)
//    s = append(s, 100)
//    return s
//}
//
//func mkmap() map[string]int {
//    m := make(map[string]int)
//    m["a"] = 1
//    return m
//}
//
//func main() {
//    m := mkmap()
//    println(m["a"])
//
//    s := mkslice()
//    println(s[0])
//    println(s[0:3:5])      // 从下标0开始取值，取到3(不包含3)，切片最大容量为5
//}




//import(
//    "fmt"
//)
//
//func main() {
//    p := new(map[string]int)        // 函数new返回指针
//    m := *p
//    m["a"] = 1                      // 错误： assignment to entry in nil map
//    fmt.Println(m)
//    // new函数可以为引用类型分配内存，但不完整以map为例子，它没有分配键值存储内存，也没有初始化内部属性，所以无法正常工行工作
//}






//*********************************************************************************************







// 2.6 类型转换     隐式转换带来的问题远大于它带来的好处; 所以Go强制要求使用显式类型转换
//import (
//    "fmt"
//)
//
//func main() {
//    //x := 100
//    //
//    //var y bool = x      // 错误： cannot use x (type int) as type bool in assignment
//    //if x {              // 错误： non-bool x (type int) used as if condition
//    //}                   // 不能将非bool类型结果当作true/false使用
//
//    a := 10
//    b := byte(a)
//    c := a + int(b)       // 混合类型表达式必须确保类型一致
//
//    fmt.Printf("%T %v\n", c, c)
//}


// 语法歧义     如果转换目标是指针、单向通道、或没有返回值的函数，那么必须使用括号避免造成语法分解错误
//func main() {
//    x := 100
//    p := *int(&x)       // 错误： cannot convert &x (type *int) to type int
//                        // 错误： invalid indirect of int(&x) (type int)
//    println(p)
//}







//*********************************************************************************************







// 2.7 自定义类型       使用关键字type定义用户自定义类型，包括基于现有类型创建,或者是结构体、函数类型
//import (
//    "fmt"
//)
//type flags byte
//
//const (
//    read flags = 1 << iota
//    write
//    exec
//)
//
//func main() {
//    f := read | exec
//    fmt.Println(write, exec)
//    fmt.Printf("%b\n", f)       // 输出二进制标记位
//}



//func main() {
//    type (
//        user struct {
//            name string
//            age uint8
//        }
//
//        event func(string) bool
//    )
//
//    u := user{"Tom", 20}
//    fmt.Println(u)
//
//    var f event = func(s string) bool {
//        println(s)
//        return s != ""      // s不为""就返回true，反之返回false
//    }
//
//    println(f("abc"))
//    println(f(""))
//    // 和var、const类似，多个type定义可以合并组成，可在函数内定义局部类型
//}


//func main() {
//    type data int
//    var d data = 10
//
//    var x int = d       // 错误： cannot use d (type data) as type int in assignment
//    println(x)
//    println(d == x)     // 错误： invalid operation: d == x (mismatched types data and int)
//}



// 未命名的类型       与bool int string等类型相比，array slice map channel等它跟具体的元素类型或长度属性有关，故它这类称作未命名类型

//* 具有相同声明的未命名类型被视作同一类型
//* 1. 具有相同基类型的指针
//* 2. 具有相同元素类型和长度的数组array
//* 3. 具有相同元素类型的切片slice
//* 4. 具有相同键值类型的字典map
//* 5. 具有相同数据类型及操作方向的通道channel
//* 6. 具有相同字段序列(字段名、字段类型、标签、以及字段顺序)的结构体struct
//* 7. 具有相同签名(参数和返回值列表，不包括参数名)的函数func
//* 8. 具有相同方法集(方法名、方法签名、不包括顺序)的接口interface


//import (
//    "fmt"
//)
//
//func main() {
//    var a struct {
//        x int `x`
//        s string `s`
//    }
//
//    var b struct {
//        x int
//        s string
//    }
//
//    b = a
//
//    fmt.Println(b)      // 错误： cannot use a (type struct { x int "x"; s string "s" }) as type struct { x int; s string } in assignment
//}


//func main() {
//    var a func(int, string)
//    var b func(string, int)
//
//    b = a               // 错误： cannot use a (type func(int, string)) as type func(string, int) in assignment
//    b("s", 1)
//    // 函数的参数顺序也是属于签名组成部分
//}


import (
    "fmt"
)

func main() {
    type data [2]int
    var d data = [2]int{1, 2}           // 基础类型相同，右值为未命名类型

    fmt.Println(d)

    a := make(chan int, 2)
    var b chan <- int = a               // 双向通道转换为单向通道,其中b为未命名类型
    b <- 2
}

//* 未命名类型转换规则：
//* 1. 所属类型相同
//* 2. 基础类型相同，且其中一个是未命名类型
//* 3. 数据类型相同，将双向通道赋值给单向通道，且其中一个为未命名类型
//* 4. 将默认值nil赋值给切片、字典、通道、指针、函数或接口
//* 5. 对象实现了目标接口