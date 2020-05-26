/*
@project: GoLand
@author: fke
@file: 表达式.go
@time: 2020-05-22 00:53:57
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



// 3.1 保留字   Go语言共25个保留关键字如下:
// break case chan const continue default defer else fallthrough for func
// go goto if import interface map package range return select struct switch type var





//*********************************************************************************************





// 3.2  运算符 (没有乘幂、绝对值运算符; 可以通过标准库math里的Pow、Abs函数实现)
//  +   -   *   /   %   &   |   ^   <<  >>  &^  +=  -=  *=  /=  %=  &=  |=  ^=  <<=  >>=  &^=   &&  ||  <-  ++  --  ==
//  <   >   =   !   !=  <=  >=  :=  ...  (  [   {   ,   .   )   ]   }   ;   :


// 二元运算符
//import (
//    "fmt"
//)
//
//func main() {
//    const v = 20                // 无显式类型声明的常量
//
//    var a byte = 10
//    b := v + a                  // v 自动转换为byte/uint8 类型
//    fmt.Printf("%T, %v\n", b, b)
//
//    const c float32 = 1.2
//    d := c + v                  // 自动转换为float32 类型
//    fmt.Printf("%T, %v\n", d, d)
//}


//func main() {
//    b := 23             // b 是有符号int类型变量
//    x := 1 << b         // 错误： invalid operation: 1 << b (shift count type int, must be unsigned integer)
//    println(x)
//    // 位移右操作数必须是无符号整数，或可以转换为无显式类型常量
//}


//import (
//   "fmt"
//)
//
//func main() {
//   a := 1.0 << 4                                // 常量表达式(包括常量展开)
//   fmt.Printf("%T, %v\n", a, a)          // int, 8
//
//   var s uint = 3
//   b := 1.0 << s                                // 错误：invalid operation: 1 << s (shift of type float64)
//   fmt.Printf("%T, %v\n", b, b)          // 因为b没有提供类型，编译器通过1.0推断其为float64，显然无法对浮点数做位运算
//
//   var c int32 = 1.0 << s                       // 自动将1.0 转换为int32类型
//   fmt.Printf("%T, %v\n", c, c)          // int32, 8
//}



// 位运算符     (AND NOT)和(XOR) 是不同的
//* AND               按位与: 都为1              a & b           0101 & 0011 = 0001
//* OR                按位或: 至少一个1           a | b           0101 | 0011 = 0111
//* XOR               按位异或: 只有一个1         a ^ b           0101 ^ 0011 = 0110
//* NOT               按位取反: (一元)           ^a              ^0111 = 1000
//* AND NOT           按位清除: (bit clear)     a &^ b           0110 &^ 1011 = 0100
//* LEFT SHIFT        位左移                    a << 2           0001 << 3 = 1000
//* RIGHT SHIFT       位右移                    a >> 2           1010 >> 2 = 0010


//import (
//    "fmt"
//)
//
//const (
//    read byte = 1 << iota
//    write
//    exec
//    freeze
//)
//
//func main() {
//    a := read | write | freeze
//    b := read | freeze | exec
//    c := a &^ b
//
//    fmt.Printf("%04b &^ %04b = %04b\n", a, b, c)
//}



// 自增       自增自减不再是运算符，只能作为独立语句，不能用于表达式
//func main() {
//    a := 1
//    ++a                             // 错误： unexpected ++  (不能前置)
//
//    if (a++) > 1 {                  // 错误：syntax error: unexpected ++, expecting } （语句不能作为表达式使用）
//    }
//
//    p := &a
//    *p++                            // 相当于 (*p)++
//    println(a)
//    // 表达式可以作为语句，但是语句不能作为表达式
//}




// 指针      指针和内存地址不能混为一谈  ("&": 取址运算符； "*": 指针运算符)
//func main() {
//    x := 10
//    var p *int = &x             // 获取地址，保存到指针变量
//    *p += 20                    // 用指针间接引用，并更新对象
//
//    println(p, *p)              // 输出指针所存储的地址，以及目标对象
//}


//func main() {
//    m := map[string]int {"a": 1}
//    println(&m["a"])            // 错误： cannot take the address of m["a"]
//}


//func main() {
//    x := 10
//    p := &x
//
//    p++                             // 无效操作： p++ (non-numeric type *int)
//    var p2 *int = p + 1             // 无效操作： p + 1 (mismatched types *int and int)
//
//    p2 = &x
//    println(p == p2)
//    // 指针类型支持相等运算符，但是不能做加减法运算和类型转换。
//}


//func main() {
//    a := struct {
//        x int
//    } {}
//
//    a.x = 100
//
//    p := &a
//    p.x += 100                      // 相当于p ->x += 100
//    println(p, p.x)
//    // 指针没有专门指向成员的“->”运算符，统一使用“.”选择表达式
//}


//func main() {
//    var a, b struct{}
//
//    println(&a, &b)
//    println(&a == &b, &a == nil)
//    // 零长度(zero-size) 对象的地址是否相等和具体实现版本有关，不过肯定不等于nil
//}







//*********************************************************************************************






// 3.3 初始化
//* 对复合类型(数组、切片、字典、结构体)变量初始化时，有一些语法限制:
//* 1. 初始化表达式必须包含类标签
//* 2. 左花括号必须在类型尾部，不能另起一行
//* 3. 多个成员初始值以逗号分隔
//* 4. 允许多行，但每行须以逗号或右花括号结束


//type data struct {
//    x int
//    s string
//}
//
//var a data =data{1, "fke"}        // 伪代码
//
//b := data{1,
//    "fke",
//}
//
//c := []int{
//    1,
//    2}
//
//d := []int{1, 2,
//    3, 4,
//    5,
//}






//*********************************************************************************************







// 3.4 流控制      Go精简(合并)了流控制语句，虽然不够便捷但是够用
//func main() {
//    x := 3
//
//    if x > 5 {                      // 条件表达式值必须是布尔类型，可省略括号，且左边的花括号不能另起一行
//        println("a")
//    } else if x < 5 && x > 0 {
//        println("b")
//    } else {
//        println("z")
//    }
//}


//func main() {
//    x := 10
//
//    //if xinit(); x == 0 {            // 优先执行xinit函数
//    //    println("a")
//    //}
//
//    if a, b := x+1, x+10; a < b {       // 定义一个或多个局部变量(也可以是函数返回值)
//        println(a)
//    } else {
//        println(b)
//    }
//    // 比较特别的是对初始化语句的支持，可定义块局部变量或执行初始化函数
//}   // 局部变量的有效范围包括整个if/else块



//import (
//   "errors"
//   "log"
//)
//
//func check (x int) error {
//   if x <= 0 {
//       return errors.New("x <= 0")
//   }
//
//   return nil
//}
//
//
////func main() {
////  x := 10
////
////  if err := check(x); err == nil {
////      x++
////      println(x)
////  } else {
////      log.Fatalln(err)
////  }
////}   // 尽可能减少代码块嵌套，让正常逻辑处于相同层次，所以优化后如下
//
//func main() {
//    x := 10
//
//    if err := check(x); err != nil {
//        log.Fatalln(err)
//    }
//
//    x++
//    println(x)
//}



//import (
//    "log"
//    "strconv"
//)
//
//func main() {
//    s := "9"
//    n, err := strconv.ParseInt(s, 10, 64)       // 使用外部变量
//
//    if err != nil {
//        log.Fatalln(err)
//    } else if n < 0 || n >10 {
//        log.Fatalln("invalid number")
//    }
//
//    println(n)                      // 避免if局部变量将该逻辑放到else块
//}



//import(
//    "log"
//    "strconv"
//)
//
//func main() {
//    s := "9"
//
//    if n, err := strconv.ParseInt(s, 10, 64); err != nil || n < 0 || n > 10 || n %2 != 0 {
//        log.Fatalln("invalid number")
//    }
//
//    println("ok")
//}    // 对于某些过于复杂的组合条件，建议将其重构为函数，虽然可能会有一些性能损失，但是更易于测试，也提高代码可维护性; 重构如下



//import (
//    "errors"
//    "log"
//    "strconv"
//)
//
//func check(s string) error {
//    n, err := strconv.ParseInt(s, 10, 64)
//    if err != nil || n < 0 || n > 10 || n%2 != 0 {
//        return errors.New("invalid number")
//    }           // check函数仅被if块调用，也可以将其称为局部函数，避免扩大作用域
//
//    return nil
//}
//
//func main() {
//    s := "9"
//
//    if err := check(s); err != nil {
//        log.Fatalln(err)
//    }
//
//    println("ok")
//}



// switch      与if类似，switch语句也用于选择执行，但使用场景会有所不同
//func main() {
//    a, b, c, x := 1, 2, 3, 2
//
//    switch x {                          // 将x与case条件匹配
//    case a, b:                          // 多个匹配条件命中其一即可(OR),变量
//        println("a | b")
//    case c:                             // 单个匹配条件
//        println("c")
//    case 4:                             // 常量
//        println("d")
//    default:
//        println("z")
//    }
//}


//func main() {
//    switch x := 5; x {
//    default:                            // 编译器确保不会先执行default块
//        x += 100
//        println(x)
//    case 5:
//        x += 50
//        println(x)
//    }
//}   // 考虑到default作用类似else,建议将其放置在switch末尾


//func main() {
//    switch x := 5; x {
//    case 5:
//        println("a")
//    case 6, 5:                          // 错误： duplicate case 5 in switch
//        println("b")
//    }
//}   // 不能出现重复的case值


//func main() {
//    switch x := 5; x {
//    default:
//        println(x)
//    case 5:
//        x += 10
//        println(x)
//
//        fallthrough                     // 继续执行下一case, 但不再匹配条件表达式
//    case 6:
//        x += 20
//        println(x)
//
//        //fallthrough                   // 错误： cannot fallthrough final case in switch
//    }
//}


//func main() {
//    switch x := 5; x {
//    case 5:
//        x += 10
//        println(x)
//
//        if x >= 15 {
//            break                          // 终止, 不再执行后续语句
//        }
//
//        fallthrough                        // 必须是case块的最后一条语句
//    case 6:
//        x += 20
//        println(x)
//    }
//}   // fallthrough必须放在case块的结尾, 可使用break语句终止


//func main() {
//    switch x := 5; {                        // 相当于"switch x := 5; true {...}"
//    case x > 5:
//        println("a")
//    case x > 0 && x <= 5:                   // 不能写成"case x > 0, x <= 5" 因为多条件是or关系
//        println("b")
//    default:
//        println("z")
//    }
//}



// for     仅有for一种循环语句
//func count() int {
//    println("count.")
//    return 3
//}
//
//func main() {
//    for i, c := 0, count(); i < c; i++ {
//        println("a", i)
//    }
//
//    c := 0
//    for c < count() {
//        println("b", c)
//        c++
//    }
//}   // 规避方法就是在初始化表达式中定义局部变量保存count结果


//func main() {
//    data := [3]string{"a", "b", "c"}
//
//    for i, s := range data {
//        println(i, s)                       // 返回下标和值
//    }
//}


//func main() {
//    data := [3]string{"a", "b", "c"}
//
//    for i := range data {
//        println(i, data[i])
//    }
//
//    for _, s := range data{         // 可以使用“_”，这样可以返回单值
//        println(s)
//    }
//
//    for range data {                // 仅迭代，不返回. 可以用来执行清空channel等操作
//    }
//}


//func main() {
//    data := [3]string{"a", "b", "c"}
//
//    for i, s := range data {
//        println(&i, &s)
//    }
//}   // 这种写法对闭包有一定的影响; 无论是for循环，还是range迭代，其定义的局部变量都会重复使用。


//import (
//    "fmt"
//)
//
//func main() {
//    data := [3]int{10, 20, 30}
//
//    for i, x := range data {                            // 从data复制品中取值
//        if i == 0 {
//            data[0] += 100
//            data[1] += 200
//            data[2] += 300
//        }
//
//        fmt.Printf("x: %d, data: %d\n", x, data[i])
//    }
//
//    for i, x := range data[:] {                         // 仅复制slice, 不包括底层的array
//        if i == 0 {
//            data[0] += 100
//            data[1] += 200
//            data[2] += 300
//        }
//
//        fmt.Printf("x: %d, data: %d\n", x, data[i])
//    }
//}   // range 会复制目标数据。受直接影响的是数组，可改用数组指针或切片类型


//func data() []int {
//    println("origin data.")
//    return []int{10, 20, 30}
//}
//
//func main() {
//    for i, x := range data() {              // 如果range目标表达式是函数调用，也仅被执行一次。
//        println(i, x)
//    }
//}




// goto 、 continue、 break
//func main() {
//start:                                      // 错误： label start defined and not used
//    for i := 0; i < 3; i++ {
//        println(i)
//
//        if i > 1 {
//            goto exit                       // 使用goto前需要先定义好标签。标签区分大小写，未使用的标签会引发编译器错误
//        }
//    }
//exit:
//    println("exit.")
//}


//func test() {
//test:
//    println("test")
//    println("test exit.")
//}
//
//func main() {
//    for i := 0; i < 3; i++ {
//        loop:
//            println(i)
//    }
//
//    goto test                           // 错误： label test not defined
//    goto loop                           // 错误： goto loop jumps into block
//}
////* 1. goto定点跳转
////* 2. break 用于switch、for、select语句，终止整个语句块执行
////* 3. continue 仅用于for循环，终止后续逻辑，立即进入下一轮循环


//func main() {
//    for i := 0; i < 10; i++ {
//        if i%2 == 0 {
//            continue                // 立即进入下一轮循环
//        }
//
//        if i > 5 {
//            break                   // 终止整个for循环
//        }
//
//        println(i)
//    }
//}


func main() {
outer:
    for x := 0; x < 5; x++ {
        for y := 0; y <= 10; y++ {
            if y > 2 {
                println()
                continue outer
            }

            if x > 2 {
                break outer
            }

            println(x, ":", y, " ")
        }
    }
}   // 配合标签break和continue可以在多层嵌套中指定目标层级。
