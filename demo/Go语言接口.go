/*
@project: GoLand
@author: fke
@file: Go语言接口.go
@time: 2019-09-19 14:41:29
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
import "fmt"

/*
//定义接口
type interface_name interface {
    method_name1 [return_type]
    method_name2 [return_type]
    method_name3 [return_type]
    ...
}

//定义结构体
type struct_name struct {
    //
}

//实现接口方法
func (struct_name_variable struct_name) method_name1() [return_type] {
    // 实现方法
}
*/


//本例中，我们定义一个接口Phone,接口里面有一个方法call()。然后我们在main函数里面定义一个Phone类型变量，并分别为其赋值。然后调用call()方法，并输出结果
type Phone interface {
    call()
}

type NokiaPhone struct {
}

type IPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
    fmt.Println("I am Nokia, I can call you!")
}

func (iphone IPhone) call() {
    fmt.Println("I am iphone, I can call you!")
}

func main() {
    var phone Phone

    phone = new(NokiaPhone)
    phone.call()

    phone = new(IPhone)
    phone.call()
}
