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








// 4.4 匿名函数     ①没有定义名字符号的函数；②匿名函数可以直接调用，保存到变量，作为参数或返回值； ③不曾使用的匿名函数会被编译器当作错误；
//             ④相比语句块，匿名函数的作用域被隔离(不使用闭包)，更加灵活，不会引发外部污染；⑤没有定义顺序限制，必要时可以抽离
//func main() {
//    func (s string) {
//        println(s)
//    } ("hello world!")                    // 直接执行
//}


//func main() {
//    add := func(x, y int) int {           // 赋值给变量
//        return x + y
//    }
//
//    println(add(1, 3))
//}


//func test(f func()) {
//    f()
//}
//
//func main() {
//    test(func() {                         // 作为参数
//        println("hello, world!")
//    })
//}


//func test() func(int, int) int {
//    return func(x int, y int) int {         // 作为返回值
//        return x + y
//    }
//}
//
//func main() {
//    add := test()
//    println(add(2, 3))
//}



//func testStruct() {
//    type calc struct {                              // 定义结构体类型
//        mul func(x, y int) int                      // 函数类型字段
//    }
//
//    x := calc{
//        mul: func(x, y int) int {
//            return x * y
//        },
//    }
//
//    println(x.mul(3, 4))
//}
//
//func testChannel() {
//    c := make(chan func(int, int) int, 2)
//
//    c <- func(x int, y int) int {
//        return x + y
//    }
//
//    println((<-c)(1, 2))
//}
//
//func main() {
//    a := testStruct
//    b := testChannel
//
//    a()
//    b()
//}   // 普通函数和匿名函数都可以作为结构体字段，或经通道传递




// 闭包       函数和其引用的环境组合体
//func test(x int) func() {           // test返回的匿名函数会引用上下文环境变量x
//    return func() {
//        println(x)
//    }
//}
//
//func main() {
//    f := test(123)              // test在main函数中执行时，它依然可正确读取x的值，这种现象就称作闭包
//    f()
//}



//func test(x int) func() {
//    println(&x)
//
//    return func() {                 // 如果分析汇编代码，其实可以发现返回的不仅仅是匿名函数，还有其引用的环境变量指针，所以说闭包是函数和其引用的环境组合体
//        println(&x, x)
//    }
//}
//
//func main() {
//    f := test(0x100)
//    f()
//}



//func test() []func() {
//    var s []func()
//
//    for i := 0; i < 3; i++ {                    // for循环复用局部变量i，添加操作只是将匿名函数放入列表，并未执行；所以main执行这些函数时，读取的是环境变量最后一次循环的值3
//        s = append(s, func() {                  // 将多个匿名函数添加到列表
//            println(&i, i)
//        })
//    }
//
//    return s                                    // 返回匿名函数列表
//}
//
//func main() {
//    for _, f := range test() {                  // 迭代执行所有的匿名函数
//        f()
//    }
//}



//func test() []func() {
//    var s []func()
//
//    for i := 0; i < 3; i++ {
//        x := i                              // x每次循环都重新定义了
//        s = append(s, func() {
//            println(&x, x)
//        })
//    }
//
//    return s
//}
//
//func main() {
//    for _, f := range test() {
//        f()
//    }
//}



//func test(x int) (func(), func()) {                 //返回2个匿名函数
//    return func() {
//        println(x)
//        x += 10                                     // 修改环境变量
//    }, func() {
//        println(x)                                  // 打印环境变量
//    }
//}
//
//func main() {
//    a, b := test(100)
//    b()
//    a()
//    b()
//}   // 多个匿名函数引用同一个环境变量，任何修改行为都会影响到其他函数取值，在并发模式下需要做同步处理








//*********************************************************************************************








// 4.5 延迟调用     语句defer向当前函数注册后执行的函数调用，被称为延迟调用，因为它们直到当前函数执行结束前才被执行，常用于资源释放，以及错误处理等操作
//import (
//    "log"
//    "os"
//)
//
//func main() {
//    f, err := os.Open("表达式.go")
//    if err != nil {
//        log.Fatalln(err)
//    }
//
//    defer  f.Close()                // 仅注册，直到main退出前才执行
//}  //  伪代码



//func main() {
//    x, y := 1, 2
//
//    defer func(a int) {
//        println("defer x, y = ", a, y)     // y为闭包引用
//    } (x)                                  // 注册时复制调用函数
//
//    x += 100                               // 对x的修改不会影响延迟调用
//    y += 200
//    println(x, y)
//}   // 延迟调用注册的是调用，必须提供执行所需的参数(哪怕为空）。参数值在注册时被复制并缓存起来。如对状态敏感可以改用指针或者闭包



//func main() {
//    defer println("a")
//    defer println("b")
//}   // 输出结果b a;     多个延迟注册按FILO次序执行



//func test() (z int) {
//    defer func() {
//        println("defer:", z)
//        println("2:", z)
//        z += 100                            // 修改命名返回值
//        println("3:", z)
//    }()
//
//    println("1:", z)
//    return 100                          // 实际执行次序 z = 100, call defer, ret
//}
//
//func main() {
//    println("test:", test())
//}   // return和panic语句会终止当前函数流程，引发延迟调用。另外return语句，它会先更新返回值




// 误用   重点：延迟调用是在函数结束时才被执行。不合理的使用方式会浪费很多资源，甚至造成逻辑错误
//eg:    循环处理多个日志文件，不恰当的defer导致文件关闭时间延长
//import (
//    "fmt"
//    "log"
//    "os"
//)
//
//func main() {
//    for i := 0; i < 10; i++ {
//        path := fmt.Sprintf("./log/%d.txt", i)
//
//        f ,err := os.Open(path)
//        if err != nil {
//            log.Println(err)
//            continue
//        } else {
//            fmt.Printf("f detail : %v\n", f)
//        }
//
//     // 这个关闭操作在main函数结束时才会执行，而不是当前循环中执行
//     // 这无端延长了逻辑结束时间和f的生命周期，平白多消耗量内存等资源
//     defer f.Close()
//    }
//}


// 上面eg例子 应该直接调用，或者重构为函数，将循环个处理算法分离
//import (
//   "fmt"
//   "os"
//   "log"
//)
//
//func main() {
//   // 日志处理算法
//   do := func(n int) {
//       path := fmt.Sprintf("./log/%d.txt", n)
//       f ,err := os.Open(path)
//       if err != nil {
//           log.Println(err)
//       }
//       // 该延迟调用在此匿名函数结束时才执行，而非main
//       defer f.Close()
//   }
//   for i := 0; i < 10; i++ {
//       do(i)
//   }
//}





// 性能   延迟调用花费更大代价。其中包括注册、调用等操作，还有额外的缓存开销
//import (
//   "sync"
//   "testing"
//)
//
//var m sync.Mutex
//
//func call()  {
//   m.Lock()
//   m.Unlock()
//}
//
//func deferCall()  {
//   m.Lock()
//   defer m.Unlock()
//}
//
//func BenchmarkCall(b *testing.B)  {
//   for i := 0; i < b.N; i++ {
//       call()                                       // BenchmarkCall-4    	20000000	        66.8 ns/op
//   }
//}
//
//func BenchmarkDefer(b *testing.B)  {
//   for i := 0; i < b.N; i++ {                       // BenchmarkDefer-4   	10000000	       195 ns/op
//       deferCall()
//   }
//}  // 将例子单独复制到一个文件中运行，例如（main_test.go）








//*********************************************************************************************









// 4.6 错误处理         官方推荐的标准做法是返回error状态

//type error interface {
//    Error() string              // 标准库将error定义为接口类型，以便实现自定义错误类型
//}


//import (
//    "errors"
//    "log"
//)
//
//var errDivByZero = errors.New("division by zero")
//
//func div(x, y int) (int, error) {
//    if y == 0 {
//        return 0, errDivByZero
//    }
//
//    return x / y, nil
//}
//
//func main() {
//    z, err := div(5, 0)
//    if err == errDivByZero {        // 应该通过错误变量，而非文本内容来判断错误类别
//       log.Fatalln(err)
//    }
//
//    println(z)
//}




//import (
//    "fmt"
//    "log"
//)
//
//type DivError struct {                          // 自定义错误类型
//    x, y int
//}
//
//func (DivError) Error() string {                // 实现error接口方法
//    return "division by zero."
//}
//
//func div(x, y int) (int, error) {
//    if y == 0 {
//        return 0, DivError{x, y}           // 返回自定义错误类型
//    }
//
//    return x / y, nil
//}
//
//func main() {
//    z, err := div(5, 0)
//
//    if err != nil {
//        switch e := err.(type) {                // 根据类型匹配
//        case DivError:                          // 自定义错误类型通常以Error为名称后缀。switch匹配case时，应将自定义类型放在前面
//            fmt.Println(e, e.x, e.y)
//        default:
//            fmt.Println(e)
//        }
//
//        log.Fatalln(err)
//    }
//    println(z)
//}   // 在调用多返回值函数时，除error外，其他返回值同样需要关注


//```go
//大量函数和方法返回error, 使得调用代码变得很难看，一堆堆的检查语句充斥在代码行间，解决思路有：
//①：使用专门的检查函数处理错误逻辑(比如记录日志)，简化检查代码
//②：在不影响逻辑的情况下, 使用defer延后处理错误状态(err退化赋值)
//③：在不中断逻辑的情况下, 将错误作为内部状态保存，等最终"提交"时再处理。
//```




// panic, recover    与error相比, panic/recover在使用方法上更接近try/catch结构化异常
//import (
//    "log"
//)
//
//func main() {
//    defer func() {
//        if err := recover(); err != nil {           // 捕获错误
//            log.Fatalln(err)
//        }
//    }()
//
//    panic("i am dead")                          // 引发错误
//    println("exit.")                        // 永远不会执行
//}   // panic/recover 是内置函数而非语句。panic会中断当前函数流程。执行延迟调用，而当前延迟调用函数recover可捕获并返回panic提交的错误对象



//import (
//    "log"
//)
//
//func test() {
//    defer println("test.1")
//    defer println("test.2")
//
//    panic("i am dead")                  // panic参数是空接口类型，因此可以使用任何对象作为错误状态
//}
//
//func main() {
//    defer func() {
//        log.Println(recover())            // 而recover返回结果同样需要做转型才能获得具体信息
//    }()
//
//    test()
//}   // 无论是否执行recover,所有延迟调用都会被执行。但中断性错误会沿调用堆栈向外传递，要么被外层捕获，要么导致进程崩溃



//import (
//    "log"
//)
//
//func main() {
//    defer func() {
//        for {
//            if err := recover(); err != nil {
//                log.Println(err)
//            } else {
//                log.Fatalln("fatal")
//            }
//        }
//    }()
//
//    defer func() {
//        panic("you are dead")               // 类似重新抛出异常(rethrow)
//    }()                                        // 可先recover捕获, 包装后重新抛出
//
//    panic("i am dead")                      // 连续调用panic, 仅最后一个会被recover捕获，因此此处其实并没有被捕获
//}



//import (
//    "log"
//)
//
//func catch() {
//    log.Println("catch:", recover())
//}
//
//func main() {
//    defer catch()                           // 捕获
//    defer log.Println(recover())            // 失败
//    defer recover()                         // 失败
//
//    panic("i am dead")
//}   // 在延迟函数中panic,不会影响后续延迟调用执行，而recover之后panic,可被再次捕获。recover必须在延迟调用函数中执行才能正常工作



//func test(x, y int) {
//    z := 0
//
//    func() {
//        defer func() {
//            if recover() != nil {
//                z = 0
//            }
//        }()
//
//        z = x / y
//    } ()
//
//    println("x / y =", z)
//}
//
//func main() {
//    test(5, 0)
//}   // 考虑到recover特性，如果要保护代码片段，那么只能将其重构为函数调用


// 调试阶段，可使用runtime/debug.PrintStack函数输出完整调用堆栈信息
import (
    "runtime/debug"
)

func test() {
    panic("i am dead")
}

func main() {
    defer func() {
        if err := recover(); err != nil {
            debug.PrintStack()
        }
    }()

    test()
}


//```Go
//建议： 除非是不可恢复性、导致系统无法正常工作的错误，否则不建议使用panic.
//```
