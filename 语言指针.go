/*
@project: GoLand
@author: fke
@file: 语言指针.go
@time: 2019-09-17 10:20:24
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
// Go语言取地址符是&，放到变量前使用就会返回相应的变量内存地址； 在指针类型前加上*号，来获取指针所指向的内容
func main() {
    var a int = 20   // 声明实际变量
    var ip *int   // 声明指针变量, 指向整型

    ip = &a     // 指针变量的存储地址
    fmt.Printf("a 变量的地址是：%x\n", &a)

    // 指针变量的存储地址
    fmt.Printf("ip 变量存储的指针地址：%x\n", ip)

    // 使用指针访问值
    fmt.Printf("*ip 变量的值： %d\n", *ip)
}
*/



/*
//空指针
func main()  {
    var ptr *int

    if(ptr !=nil) {
        fmt.Println("ptr 不是空指针")
    }else {
        fmt.Println("ptr 是空指针")
    }
    fmt.Printf("ptr 的值为：%x\n", ptr)
}
*/



/*
//语言指针数组
const MAX int = 3

func main()  {
    a := []int{10, 100, 200}
    var i int

    for i = 0; i < MAX; i++ {
        fmt.Printf("a[%d] = %d\n", i, a[i])
    }
}
*/


/*
//eg: 假如我们需要保存数组，这样我们就需要使用到指针
const MAX int = 3

func main()  {
    a := []int{10, 100, 200}
    var i int
    var ptr [MAX]*int       // ptr为整型指针数组，因此每个元素都指向一个值

    for i = 0; i < MAX; i++ {
        ptr[i] = &a[i]   // 整数地址赋值给指针数组
    }

    for i =0; i < MAX; i++ {
        fmt.Printf("a[%d] = %d\n", i, *ptr[i])
    }
}
*/




/*
//指向指针的指针  var ptr **int
func main()  {

    var a int
    var ptr *int
    var pptr **int

    a = 3000

    ptr = &a  // 指针ptr的地址
    pptr = &ptr   // 指向指针ptr的地址

    fmt.Printf("变量a = %d\n", a)
    fmt.Printf("指针变量 *ptr = %d\n", *ptr)
    fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)
}
*/




/*
//三重指针及其对应关系
func main()  {
    var a int = 5
    var ptr *int = &a  // 把ptr指针指向a所在地址
    var pts *int = ptr  // 开辟一个新的指针，指向ptr指针指向的地方
    var pto **int = &ptr  //二级指针，指向一个地址，这个地址存储的是一级指针的地址
    var pt3 ***int = &pto   //三级指针，指向一个地址，这个地址存储的是二级指针的地址，二级指针同上

    fmt.Println("a的地址:",&a,
                "\n 值", a, "\n\n",

                "ptr指针所在地址:",&ptr,
                "\n ptr指向的地址:",ptr,
                "\n ptr指针指向地址对应的值",*ptr,"\n\n",

                "pts指针所在地址:",&pts,
                "\n pts指向的地址:", pts,
                "\n pts指针指向地址对应的值:",*pts,"\n\n",

                "pto指针所在地址:",&pto,
                "\n pto指向的指针（ptr）的存储地址:",pto,
                "\n pto指向的指针（ptr）所指向的地址: " ,*pto,
                "\n pto最终指向的地址对应的值（a）",**pto,"\n\n",

                "pt3指针所在的地址:",&pt3,
                "\n pt3指向的指针（pto）的地址:",pt3,//等于&*pt3,
                "\n pt3指向的指针（pto）所指向的指针的（ptr）地址", *pt3, //等于&**pt3,
                "\n pt3指向的指针（pto）所指向的指针（ptr）所指向的地址（a）:",**pt3, //等于&***pt3,
                "\n pt3最终指向的地址对应的值（a）", ***pt3)
}
*/





//Go语言指针作为函数参数
func main()  {
    // 定义局部变量
    var a int = 100
    var b int = 200

    fmt.Printf("交换前a的值： %d\n", a)
    fmt.Printf("交换前b的值： %d\n", b)

    // 调用函数用于交换值
    swap(&a, &b)
    fmt.Printf("交换后a的值： %d\n", a)
    fmt.Printf("交换后b的值： %d\n", b)
}

func swap(x *int, y *int) {
    var temp int
    temp = *x
    *x = *y
    *y = temp
}
