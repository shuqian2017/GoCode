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



//import (
//  "fmt"
//)
//
//func main() {
//  // 数据类型为 切片，长度为5，容量为10
//  a := make([]int,5,10)
//  fmt.Println(a,cap(a),len(a))        // 运行结果 : [0 0 0 0 0] 10 5
//
//  // 切片追加元素，当超过原来的的容量的时候，会翻倍扩容，但不是一定翻倍，如果容量太大不会再翻倍
//  for i := 0; i < 10; i++  {
//     a = append(a, i)
//  }
//
//  // 再对值进行修改
//  for i:= 0; i < 10; i++  {
//     a[i] = i
//  }
//  fmt.Println(a,cap(a),len(a)) // [0 1 2 3 4 5 6 7 8 9 5 6 7 8 9] 20 15
//}  //cap 数据类型的容量; len 数据类型的实际长度。cap主要是为了让slice提供可变长度



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

//import (
//   "fmt"
//)
//
//func main() {
//   x, y := 10, 20
//   a := [...]*int{&x, &y}           // 元素为指针的指针数组
//   p := &a                          // 存储数组地址的指针
//
//   fmt.Printf("%T, %v\n", a, a)
//   fmt.Printf("%T, %v\n", p, p)
//}



//func main() {
//   a := [...]int{1, 2}
//   println(&a, &a[0], &a[1])
//} // 可获取任意元素地址


//func main() {
//   a := [...]int{1, 2}
//   p := &a
//
//   p[1] += 10         // 数组指针可直接用来操作元素
//   println(p[1])
//}







// 复制    与C数组变量隐式作为指针不同，Go数组是值类型，赋值和传参操作都会复制整个数组数据

//import (
//   "fmt"
//)
//
//func test(x [2]int) {
//   fmt.Printf("x: %p, %v\n", &x, x)
//}
//
//func main() {
//   a := [2]int{10, 20}
//   var b [2]int
//   b = a
//
//   fmt.Printf("a: %p, %v\n", &a, a)
//   fmt.Printf("b: %p, %v\n", &b, b)
//
//   test(a)
//}



//import (
//   "fmt"
//)
//
//func test(x *[2]int) {
//   fmt.Printf("x: %p, %v\n", x, *x)       // 可以使用指针或切片，以此避免数据复制
//   x[1] += 100
//}
//
//func main() {
//   a := [2]int{10, 20}
//   test(&a)
//
//   fmt.Printf("a: %p, %v\n", &a, a)
//}










//*********************************************************************************************














// 5.3 切片
// *切片(slice)本身并非动态数组或数组指针。它内部通过指针引用底层数组，设定相关属性将数组读写操作限定在指定区域内。 切片本身是个只读对象
// *属性cap表示切片所引用数组的真实长度，len用于限定可读的写元素数量。另外数组必须addressable否则会引发错误

//func main() {
//   m := map[string][2]int{
//      "a": {1, 2},
//   }
//
//   s := m["a"][:]    // 无效  m["a"][:] (slice of unaddressable value)
//   println(s)
//}



//func main() {
//   x := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
//   s := x[2:5]
//
//   for i := 0; i < len(s); i++ {
//      println(s[i])                 // 和数组一样，切片也使用索引号访问元素内容。起始值为0，而非使用数组本身的索引号
//   }
//}



//import (
//   "fmt"
//)
//
//func main() {
//   s1 := make([]int, 3, 5)                // 指定len, cap底层数组初始化为0
//   s2 := make([]int, 3)                   // 省略cap 和 len相等
//   s3 := []int{10, 20, 5: 30}             // 按初始化元素分配底层数组，并设置len, cap
//
//   fmt.Println(s2, len(s2), cap(s2))
//   fmt.Println(s1, len(s1), cap(s1))
//   fmt.Println(s3, len(s3), cap(s3))      //  结果： [10 20 0 0 0 30] 6 6
//}



//import (
//   "fmt"
//   "reflect"
//   "unsafe"
//)
//
//func main() {
//   var a []int             // 此操作仅定义一个 []int 类型变量，并未执行初始化操作
//   b := []int{}            // 使用初始化表达式完成了全部创建过程
//
//   // a == nil 仅表示它是一个未初始化的切片对象，但切片本身依然会分配内存。对nil切片执行slice[:]操作，同样返回nil
//   println(a == nil, b == nil)         // 结果 ： true false
//
//   fmt.Printf("a: %#v\n", (*reflect.SliceHeader)(unsafe.Pointer(&a)))
//   fmt.Printf("b: %#v\n", (*reflect.SliceHeader)(unsafe.Pointer(&b)))
//   fmt.Printf("a size: %d\n", unsafe.Sizeof(a))
//}



//func main() {
//   a := make([]int, 1)
//   b := make([]int, 1)
//
//   println(a == b)  // 无效 a == b (slice can only be compared to nil)
//   println(&a, &b)
//}  // 切片不支持比较操作，就算元素类型支持也不行。仅能判断是否为nil



// 可以获取元素地址，但不能向数组那样直接用指针访问元素内容
//import (
//   "fmt"
//)
//
//func main() {
//   s := []int{0, 1, 2, 3, 4}
//
//   p := &s
//   p0 := &s[0]
//   p1 := &s[1]
//
//   println(p, p0, p1)
//   (*p)[0] += 100
//   *p1 += 100
//
//   fmt.Println(s)
//}




//import (
//  "fmt"
//)

//func main() {
//  x := [][]int{
//     {1, 2},
//     {10, 20, 30},
//     {100},
//  }
//
//  fmt.Println(x[1])              // 结果： [10 20 30]
//
//  x[2] = append(x[2], 200, 300)
//  fmt.Println(x[2], x)           // 结果： [100 200 300] [[1 2] [10 20 30] [100 200 300]]
//}




// 切片是很小的结构体对象，用来代替数组传参可避免复制开销；make函数允许在运行期动态指定数组长度，绕开了数组类型编译期常量的限制。
// 并非所有时候都适合用切片来代替数组。因为切片底层数组可能会在堆上分配内存。而小数组在栈上拷贝的消耗未必比make代价大
//import (
//  "testing"
//)
//
//func array() [1024]int {
//  var x [1024]int
//  for i := 0; i < len(x); i++ {
//     x[i] = i
//  }
//
//  return x
//}
//
//func slice() []int {
//  x := make([]int, 1024)
//  for i := 0; i < len(x); i++ {
//     x[i] = i
//  }
//
//  return x
//}
//
//func BenchmarkArray(b *testing.B) {                   //  500000	      3106 ns/op
//  for i := 0; i < b.N; i++{
//     array()
//  }
//}
//
//func BenchmarkSlice(b *testing.B) {                   //  300000	      5144 ns/op
//  for i:= 0; i < b.N; i++{
//     slice()
//  }
//}  // 复制此段代码到 main_test.go文件运行即可






// reslice   将切片视作[cap]slice数据源，据此创建新切片对象。不能超出cap，但不受len限制
//import (
//    "fmt"
//)
//
//func main() {
//    d := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
//    s1 := d[3:7]
//    s2 := s1[1:3]
//
//    for i := range s2 {
//        s2[i] += 100        // 新建切片对象依旧指向原底层数组，也就是说修改对所有关联切片可见
//    }
//
//    fmt.Println(d)          // 结果： [0 1 2 3 104 105 6 7 8 9]
//    fmt.Println(s1)         // 结果： [3 104 105 6]
//    fmt.Println(s2)         // 结果： [104 105]
//}




//import (
//    "fmt"
//    "errors"
//)
//
//func main() {
//    // 栈最大容量5
//    stack := make([]int, 0, 5)
//
//    // 入栈
//    push := func(x int) error{
//        n := len(stack)
//        if n == cap(stack) {
//            return errors.New("stack is full")
//        }
//
//        stack = stack[:n+1]           // 利用reslice操作，很容易就能实现一个栈式数据结构
//        stack[n] = x
//
//        return nil
//    }
//
//    // 出栈
//    pop := func() (int, error) {
//        n := len(stack)
//        if n == 0 {
//            return 0, errors.New("stack is empty")
//        }
//
//        x := stack[n-1]
//        stack = stack[:n-1]
//
//        return x, nil
//    }
//
//    // 入栈测试
//    for i := 0; i < 7; i++ {
//        fmt.Printf("push %d: %v, %v\n", i, push(i), stack)
//    }
//
//    // 出栈测试
//    for i := 0; i < 7; i++ {
//        x, err := pop()
//        fmt.Printf("pop: %d, %v, %v\n", x, err, stack)
//    }
//}






// append   向切片尾部slice[len] 添加数据，返回新的切片对象
//import (
//    "fmt"
//)
//
//func main() {
//    s := make([]int, 0, 5)
//
//    s1 := append(s, 10)
//    s2 := append(s1, 20 ,30)        // 追加多个数据
//
//    fmt.Println(s, len(s), cap(s))         // 不会修改原slice属性
//    fmt.Println(s1, len(s1), cap(s1))
//    fmt.Println(s2, len(s2), cap(s2))
//}




//import (
//    "fmt"
//)
//
//func main() {
//    s := make([]int, 0, 100)
//    fmt.Println(s)
//    s1 := s[:2:4]
//    s2 := append(s1, 1, 2, 3, 4, 5, 6)      // 超出新cap：s1 cap=4限制，分配新的底层数组
//
//    fmt.Printf("s1: %p: %v\n", &s1[0], s1)
//    fmt.Printf("s2: %p: %v\n", &s2[0], s2)
//
//    fmt.Printf("s data: %v\n", s[:10])
//    fmt.Printf("s1 cap: %d , s2 cap: %d\n", cap(s1), cap(s2))       // 运行结果 ： s1 cap: 4 , s2 cap: 8
//}   // 每次扩容并非都是上次的2倍。如果切片较大，会尝试扩容1/4， 以节约内存



//import (
//    "fmt"
//)
//
//func main() {
//    var s []int
//    s = append(s, 1, 2, 3, 4)       // 向nil切片追加数据时，会为其分配底层数组内存
//    fmt.Println(s)      // 运行结果： [1 2 3 4]
//}   // 因为存在重新分配底层数组的缘故，所以在某些时候建议预留足够多的空间，避免中途内存分配和数据复制开销







// copy    在2个切片对象间进行复制数据，允许指向同一底层数组，允许目标区间重叠。最终复制长度以较短的切片长度为准
//import (
//    "fmt"
//)
//
//func main() {
//    s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
//
//    s1 := s[5:8]
//    n := copy(s[3:], s1)
//    fmt.Println(n, s, s1)           // 运行结果： 3 [0 1 2 5 6 7 6 7 8 9] [7 6 7]
//
//    s2 := make([]int, 6)
//    fmt.Println(s2)                 // 运行结果： [0 0 0 0 0 0]
//    n = copy(s, s2)                 // s2 >> s
//    fmt.Println(n, s, s1, s2)       // 6 [0 0 0 0 0 0 6 7 8 9] [0 6 7] [0 0 0 0 0 0]
//}



//import (
//    "fmt"
//)
//
//func main() {
//    b := make([]byte, 3)
//    n := copy(b, "abcdef")          // 直接从字符串中复制数据到[]byte
//    fmt.Println(n, b)
//}   // 如果切片长时间引用大数组中的小片段，建议新建独立切片，复制出所需要的数据，以便原数组内存可被及时回收










//*********************************************************************************************












// 字典       一种使用频率极高的数据结构。将其作为内置类型，从运行时层面进行优化，可以获得更高效的性能
// *字典是引用类型，使用make函数或初始化表达语句来创建
//import (
//    "fmt"
//)
//
//func main() {
//    m := make(map[string]int)
//    m["a"] = 1
//    m["b"] = 2
//
//    m2 := map[int]struct{       // 值为匿名结构类型
//        x int
//    }{
//        1 : {x: 100},           // 可省略key、value类型标签
//        2 : {x: 200},
//    }
//    fmt.Println(m, m2)          // 运行结果： map[a:1 b:2] map[1:{100} 2:{200}]
//
//    m["a"] = 10                 // 修改
//    m["c"] = 30                 // 新增
//
//    if v, ok := m["d"]; ok {    // 使用ok-idiom判断key是否存在，返回值
//        fmt.Println(v)
//    } else {
//        fmt.Println("键值对不存在 m['d']")
//    }
//
//    delete(m, "d")          // 删除键值对。不存在时，不会出错
//}

