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
func main() {
    x := 10
    var p *int = &x             // 获取地址，保存到指针变量
    *p += 20                    // 用指针间接引用，并更新对象

    println(p, *p)              // 输出指针所存储的地址，以及目标对象
}




