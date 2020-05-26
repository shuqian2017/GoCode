/*
@project: GoLand
@author: fke
@file: 函数.go
@time: 2020-05-26 22:11:31
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



// 4.1      关键字func用于定义函数
//* 1. 无需前置声明
//* 2. 不支持命名嵌套定义(nested)
//* 3. 不支持同名函数重载(overload)
//* 4. 不支持默认参数
//* 5. 支持不定长变参
//* 6. 支持多返回值
//* 7. 支持命名返回值
//* 8. 支持匿名函数和闭包

//func test()
//{                                       // 错误： syntax error: unexpected semicolon or newline before {
//}
//
//func test() {                           // 错误3： test redeclared in this block
//}
//
//func main() {
//    func add(x, y int) int {            // 错误2： syntax error: unexpected add, expecting (
//        return x + y
//    }
//}


//func hello() {
//    println("hello, world!")
//}
//
//func exec(f func()) {
//    f()
//}
//
//func main() {
//    f := hello
//    exec(f)
//}   // 函数第一类对象(注释：指可在运行期创建，可用作函数参数或返回值，可存入变量的实体)，具备相同签名(参数及返回值列表)的视作同一类型


//func a() {}
//func b() {}
//
//func main() {
//    println(a == nil)
//    println(a == b)                 // 错误： invalid operation: a == b (func can only be compared to nil)
//}   // 函数只能判断其是否为nil，不支持其他比较操作


//func test() *int {
//    a := 0x100
//    return &a                   // 从函数返回局部变量指针是安全的，由编译器来决定是否在堆上分配内存
//}
//
//func main() {
//    var a *int = test()
//    println(a, *a)
//}


// 建议命名规则
//* 1. 通常是动词和介词加上名词,例如scanWords
//* 2. 避免不必要的缩写,printError要比printErr更好一些
//* 3. 避免使用类型关键字,比如buildUserStruct看上去会很别扭
//* 4. 避免歧义,不能有多种用途的解释造成误解
//* 5. 避免只能通过大小写区分的同名函数
//* 6. 避免与内置函数同名,这会导致误用
//* 7. 避免使用数字,除非是特定的专有名词,例如UTF8
//* 8. 避免添加作用域提示前缀
//* 9. 统一是使用camel/pascal case拼写风格
//* 10. 使用相同术语,保持一致性
//* 11. 使用习惯用语,比如init表示初始化,is/has返回布尔值结果
//* 12. 使用反义词组命名行为相反的函数,比如get/set、 min\max等





//*********************************************************************************************






// 4.2 参数       Go不支持有默认值的可选参数，不支持命名实参；调用时必须按照顺序传递指定类型和数量的实参；参数列表中，相邻的同类型参数可合并
//func test(x, y int, s string, _ bool) *int {
//    return nil
//}
//
//func main() {
//    test(1, 2, "abc")               // 错误: not enough arguments in call to test
//}



//func add(x, y int) int {
//    x := 100                          // 错误： no new variables on left side of :=
//    var y int                         // 错误： y redeclared in this block
//    return x + y
//}
//
//func main() {
//    var f int = add(100, 200)
//    println(f)
//}   // 形参指函数定义中的参数，实参则是函数调用时所传递的参数，形参类似函数局部变量，实参则是函数外部对象。因此不能在相同层次定义同名变量



//import (
//    "fmt"
//)
//
//func test(x *int) {
//    fmt.Printf("pointer: %p, target: %v\n", &x, x)      // 输出形参x的地址
//}
//
//func main() {
//    a := 0x100
//    p := &a
//    fmt.Printf("pointer: %p, target: %v\n", &p, p)      // 输出实参p的地址
//
//    test(p)
//}   // 函数调用前，会为形参和返回值分配内存空间,并将实参拷贝到形参内存；从运行结果来看，虽然形参和实参都指向同一目标，但传递指针时依然被复制



//func test(p *int) {
//    go func() {                 // 延长P生命周期
//        println(p)
//    } ()
//}
//
//func main() {
//    x := 100
//    p := &x
//    test(p)
//}



//func test(p **int) {
//    x := 100
//    *p = &x
//}
//
//func main() {
//    var p *int
//    test(&p)
//    println(p, &p, *p)
//}   // 要实现传出参数(out), 通常建议使用返回值。当然也可继续使用二级指针



//import (
//    "log"
//    "time"
//)

//type serverOption struct {
//    address string
//    port int
//    path string
//    timeout time.Duration
//    log *log.Logger
//}
//
//func newOption() *serverOption {
//    return &serverOption{                       // 默认参数
//        address: "0.0.0.0",
//        port: 8080,
//        path: "/var/test",
//        timeout: time.Second * 5,
//        log: nil,
//    }
//}
//
//func server(option *serverOption) {
//    println(option.port)
//}
//
//func main() {
//    opt := newOption()
//    opt.port = 8085                             // 命名参数设置
//
//    server(opt)
//}   // 函数参数过多，可以将其重构为一个复合结构类型，算是变相实现可选参数和命名实参功能




// 变参       变参本质上是一个切片。只能同时接收一到多个同类型参数，且必须放到列表尾部
//import (
//    "fmt"
//)
//
//func test(s string, a ...int) {
//    fmt.Printf("%T, %v\n", a, a)                // 显示类型和值
//}
//
//func main() {
//    test("abc", 1, 2, 3, 4)
//}


//import "fmt"
//func test(a ...int) {
//   fmt.Println(a)
//}
//
//func main() {
//   a := [3]int{10, 20, 30}
//   test(a[:]...)                               // 转换为slice后展开
//}   // 将切片作为变参时，需要进行展开操作。如果是数组，先将其转换为切片



//import "fmt"
//func test(a ...int) {
//    for i := range a {
//        a[i] += 100
//    }
//}
//
//func main() {
//    a := []int{10, 20, 30}
//    test(a...)
//
//    fmt.Println(a)
//}   // 变参是切片，那么参数复制的是切片自身，并不包括底层数组，也因此可以修改原数据。有需要可以使用内置函数copy复制底层数据






//*********************************************************************************************








// 4.3 返回值      有返回值的函数，必须有明确的return终止语句
//func test(x int) int {
//    if x > 0 {
//        return 1
//    } else if x < 0 {
//        return -1
//    }                           // 错误： missing return at end of function
//}
//
//func main() {
//    x := test(5)
//    println(x)
//}



//import (
//    "errors"
//    "fmt"
//)
//
//func div(x, y int) (int, error) {
//    if y == 0 {
//        return 0, errors.New("division by zero")
//    }
//    return x / y, nil
//}
//
//func main() {
//    if value, err := div(30, 0); err != nil {
//        fmt.Println(value, err)
//    } else {
//        println(value)
//    }
//}   // 借鉴动态语言多返回值模式, 函数可以返回更多状态, 尤其是error模式



//import (
//    "errors"
//    "fmt"
//)
//func div(x, y int) (int, error) {
//    if y == 0 {
//        return 0, errors.New("division by zero!")
//    }
//    return x / y, nil
//}
//
//func log(x int, err error) {
//    fmt.Println(x, err)
//}
//
//func test() (int, error) {
//    return div(5, 0)                            // 多返回值用作return结果
//}
//
//func main() {
//    log(test())                                      // 多返回值用作实参
//}




// 命名返回值        命名返回值和参数一样，可当作函数局部变量使用，最后由return隐式返回。
//import (
//    "errors"
//    "fmt"
//)
//
//func div(x, y int) (z int, err error) {
//    if y == 0 {
//        err = errors.New("division by zero!")
//        return
//    }
//
//    z = x / y
//    return                              // 相当于"return z, err"
//}
//
//func main() {
//    value, err := div(5, 0)
//    fmt.Println(value, err)
//}



//func add(x, y int) (z int) {
//    {
//        z := x + y                      // 新定义的同名局部变量，同名遮蔽
//        return                         // 错误： z is shadowed during return （改为return z即可）
//    }
//}
//
//func main() {
//    value := add(5, 3)
//    println(value)
//}   // 相当如果返回值能明确其含义，就可以不用对其命名 例如： func NewUser() (*User, error)







//*********************************************************************************************








// 匿名函数     没有定义名字符号的函数