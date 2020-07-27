/*
@project: GoLand
@author: fke
@file: 数据.go
@time: 2020-06-15 23:39:16
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


// 5.1 字符串     字符串是不可变字节(byte)序列，其本身是一个复合结构

//import (
//    "fmt"
//)
//
//func main() {
//    s := "菜徐坤\x61\142\u0041"
//
//    fmt.Printf("%s\n", s)
//    fmt.Printf("% x, len: %d\n", s, len(s))
//}   // 内置函数len返回字节数组长度, cap不接受字符串类型参数。


//func main() {
//    var s string
//
//    println(s == "")                // true
//    println(s == nil)               // 错误：s == nil (mismatched types string and nil)
//}   // 字符串默认值是"",而不是nil


//func main() {
//    s := `line\r\n
//    line 2`
//
//    println(s)
//}   // 编译器不会解析原始字符串内的注释语句，且前置缩进空格也属于字符串内容



//func main() {
//    s := "ab" +                     // 跨行时, 加法操作符必须在上一行结尾
//        "cd"
//
//    println(s == "abcd")
//    println(s > "abc")
//}   // 支持 "!=、 ==、 <、 >、 +、 +=、"操作符



//func main() {
//    s := "abc"
//
//    println(s[1])
//    println(&s[1])              // 错误: cannot take the address of s[1]
//}   // 允许以索引号访问字节数组(非字符),但不能获取元素地址



//import (
//    "fmt"
//    "reflect"
//    "unsafe"
//)
//
//func main() {
//    s := "abcdefg"
//
//    s1 := s[:3]
//    s2 := s[1:4]
//    s3 := s[2:]
//
//    println(s1, s2, s3)
//    // 提示：
//    // reflect.StringHeader 和 string 头部结构相同
//    // unsafe.Pointer 用于指针类型转换
//
//    fmt.Printf("%#v\n", (*reflect.StringHeader) (unsafe.Pointer(&s)))
//    fmt.Printf("%#v\n", (*reflect.StringHeader) (unsafe.Pointer(&s1)))
//}   // 以切片语法(起始和结束索引号)返回子串时，其内部依旧指向原字节数组



//import (
//    "fmt"
//)
//
//func main() {
//    s := "书神"
//
//    fmt.Printf("字符串s的长度为：%d\n", len(s))
//    for i := 0; i < len(s); i++ {                           // byte
//        fmt.Printf("%d: [%c]\n", i, s[i])
//    }
//
//    for i, c := range s {                                   // rune： 返回数组索引号，以及Unicode字符
//        fmt.Printf("%d: [%c]\n", i, c)
//    }
//}   // 使用for遍历字符串时，分byte和rune两种方式





// 转换   要修改字符串须将其转换为可变类型([]rune 或者 []byte),待完成后再转换回来。过程中一定会重新分配内存，并复制数据
//import (
//    "fmt"
//    "reflect"
//    "unsafe"
//)
//
//func pp(format string, ptr interface{}) {
//    p := reflect.ValueOf(ptr).Pointer()
//    h := (*uintptr) (unsafe.Pointer(p))
//    fmt.Printf(format, *h)
//}
//
//func main() {
//    s := "hello, world!"
//    fmt.Printf("%p\n", &s)
//    pp("s: %x\n", &s)
//
//    bs := []byte(s)
//    s2 := string(bs)
//
//    pp("string to []byte, bs: %x\n", &bs)
//    pp("[]byte to string, s2: %x\n", &s2)
//
//    rs := []rune(s)
//    s3 := string(rs)
//
//    pp("string to []byte, bs: %x\n", &rs)
//    pp("[]byte to string, s2: %x\n", &s3)
//}



//import (
//    "unsafe"
//)
//
//func toString(bs []byte) string {
//    return *(*string) (unsafe.Pointer(&bs))
//}
//
//func main() {
//    bs := []byte("hello, world!")
//    s := toString(bs)
//
//    printDataPointer("bs: %x\n", &bs)
//    printDataPointer("s: %x\n", &s)
//}     // 伪代码



//import (
//    "fmt"
//)
//
//func main() {
//    var bs []byte
//    bs = append(bs, "abc"...)           // 使用append函数，可将string直接追加到[]byte内
//
//    // 说明：
//    //  1. 将[]byte转换为string key, 去map[string]查询的时候
//    //  2. 将string转换为[]byte，进行for range迭代时，直接取字节赋值给局部变量。
//    fmt.Println(bs)
//}



//import (
//   "fmt"
//)
//
//func main() {
//   m := map[string]int {
//       "abc": 123,
//   }
//
//   key := []byte("abc")
//   x, ok := m[string(key)]
//
//   fmt.Printf("%T, %c\n", key, key)
//   println(x, ok)
//}





// 性能   除类型转换之外，动态构建字符串也容易造成性能问题
//import (
//   "testing"
//)
//
//func test() string {
//   var s string
//   for i := 0; i < 1000; i++ {
//      s += "a"                          // 用加法操作字符串拼接，每次都需要重新分配内存。如此构建超大字符串时，性能就会极差
//   }
//
//   return s
//}
//
//func BenchmarkTest(b *testing.B) {
//   for i := 0; i < b.N; i++ {
//      test()                                  // BenchmarkTest-4   	    3000	    442980 ns/op
//   }
//}   // 运行需要复制该部分内容，到demo_test.go中


// *** 上面例子优化后的版本1
//import (
//   "testing"
//)
//
//func Join(a []string, sep string) string {
//   // 统计分隔符长度
//   n := len(sep) * (len(a) - 1)
//
//   // 统计所有拼接字符串长度
//   for i := 0; i < len(a); i++ {
//      n += len(a[i])
//   }
//
//   // 一次分配所需长度数组空间
//   b := make([]byte, n)
//
//   // 拷贝数据
//   bp := copy(b, a[0])
//   for _, s := range a[1:] {
//      bp += copy(b[bp:], sep)
//      bp += copy(b[bp:], s)
//   }
//
//   return string(b)
//}
//
//func test() string {
//   s := make([]string, 1000)
//   for i := 0; i < 1000; i++ {
//      s[i] = "a"
//   }
//
//   return Join(s, "")
//}
//
//func BenchmarkTest(b *testing.B) {
//   for i := 0; i < b.N; i++ {
//      test()                                  // BenchmarkTest-4   	   50000	     30231 ns/op
//   }
//}   // 运行需要复制该部分内容，到demo_test.go中


//*** 上面例子优化后的版本2
//import (
//   "bytes"
//   "testing"
//)
//
//func test() string {
//   var b bytes.Buffer                           // bytes.Buffer也能完成优化1中的操作，且性能相当
//   b.Grow(1000)
//
//   for i := 0; i < 1000; i++ {
//      b.WriteString("a")
//   }
//
//   return b.String()
//}
//
//func BenchmarkTest(b *testing.B) {
//   for i := 0; i < b.N; i++ {
//      test()                                   // BenchmarkTest-4   	   50000	     30039 ns/op
//   }
//}   // 对于数量较少的字符串格式化拼接。可使用fmt.Sprintf、text/template等方法。






// Unicode    类型rune专门用来存储Unicode，它是int32的别名。使用单引号的字面量，其默认类型就是rune
//import (
//   "fmt"
//)
//
//func main() {
//   r := '书'
//   fmt.Printf("%T %c\n", r, r)
//}



//import (
//   "fmt"
//)
//
//func main() {
//   r := '书'
//
//   s := string(r)
//   b := byte(r)
//
//   s2 := string(b)
//   r2 := rune(b)
//
//   fmt.Println(s, b, s2, r2)
//}  // 除[]rune外，还可以直接在rune、byte、string间进行转换。



//import (
//  "fmt"
//  "unicode/utf8"
//)
//
//func main() {
//  s := "书神"
//  println(len(s), utf8.RuneCountInString(s))    // 标准库Unicode提供丰富的操作函数。还可以用RuneCountInString代替len来返回准确的字符数
//
//  s = string(s[0:1] + s[3:4])            // 获取并拼接一个"不合法"的字符串； 字符串存储的字节数组，不一定就是合法的UTF-8文本
//  fmt.Println(s, utf8.ValidString(s))
//}










//*********************************************************************************************












//  5.2  数组    定义数组类型时，数组长度必须是非负整型常量表达式。元素类型相同，但是长度不同的数组不属于同一类型。
//func main() {
//   var d1 [3]int
//   var d2 [2]int
//
//   d1 = d2                    // 错误： cannot use d2 (type [2]int) as type [3]int in assignment
//}


//import (
//  "fmt"
//)
//
//func main() {
//  var a [4]int                  // 元素自动初始化为0
//
//  b := [4]int{2, 5}             // 未提供初始化值的元素自动初始化为0
//  c := [4]int{5, 3: 10}         // 可指定索引位置初始化
//
//  d := [...]int{1, 2, 3}        // 编译器按初始化值数量确定数组长度
//  e := [...]int{10, 3: 100}     // 支持索引初始化，但注意数组长度与此有关
//
//  fmt.Println(a, b, c, d, e)
//}



//import (
//   "fmt"
//)
//
//func main() {
//   type user struct {
//      name string
//      age byte
//   }
//
//   d := [...]user{
//      {"Tom", 20},      // 省略了类型标签
//      {"fke", 18},
//   }
//
//   fmt.Printf("%#v\n", d)
//} // 对于结构等复合类型，可省略元素初始化类型标签



//import (
//   "fmt"
//)
//
//func main() {
//   a := [2][2]int{
//      {1, 2},
//      {3, 4},
//   }
//
//   b := [...][2]int{       // 在定义多维数组时，仅第一维度允许使用"..."
//      {10, 20},
//      {30, 40},
//   }
//
//   c := [...][2][2]int{    // 三维数组
//      {
//         {1, 2},
//         {3, 4},
//      },
//      {
//         {10, 20},
//         {30, 40},
//      },
//   }
//
//   fmt.Println(a)
//   fmt.Println(b)
//   fmt.Println(c)
//}



//func main() {
//   a := [2]int{}
//   b := [...][2]int{
//      {10, 20},
//      {30, 40},
//      {50, 60},
//   }
//
//   println(len(a), cap(a))             // len和cap都返回第一维度长度
//   println(len(b), cap(b))
//   println(len(b[1]), cap(b[1]))
//}



//func main() {
//    var a, b [2]int
//    println(a ==b)
//
//    c := [2]int{1, 2}
//    d := [2]int{0 ,1}
//    println(c != d)
//
//    var e, f [2]map[string]int
//    println(e == f)              // 无效操作： e == f ([2]map[string]int cannot be compared)
//}






// 指针   要分清指针数组和数组指针的区别。
//'''
//指针数组： 指元素为指针类型的数组
//数组指针： 指获取数组变量的地址
//'''

import (
   "fmt"
)

func main() {
   x, y := 10, 20
   a := [...]*int{&x, &y}           // 元素为指针的指针数组
   p := &a                          // 存储数组地址的指针

   fmt.Printf("%T, %v\n", a, a)
   fmt.Printf("%T, %v\n", p, p)
}
