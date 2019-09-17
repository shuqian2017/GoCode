/**
@project: GoCode
@author: fke
@file: test.go
@time: 2019-05-28 12:39:40
# code is far away from bugs with the god animal protecting
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
**/




/*package main
import "fmt"
var age int

func main() {

	//Author by 菜鸟教程

	fmt.Println("菜鸟教程, runoob.com")
	fmt.Println("Google" + " Runoob")
	fmt.Println(age)
}*/

package main

import "fmt"
func main()  {
	var A string = "Runoob"
	fmt.Println(A)

	var B, C int = 1, 2
	fmt.Println(B, C)

	// 声明一个变量并初始化
	var a = "RUNOOB"
	fmt.Println(a)

	// 没有初始化就为零值
	var b int
	fmt.Println(b)

	// bool 零值为false
	var c bool
	fmt.Println(c)

	// : 也可用用来申明新的变量
	f := 15
	fmt.Println(f)
}


func ceshi() {
	f, n := 15, 30
	fmt.Println(f, n)
}

