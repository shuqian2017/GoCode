/**
@project: GoCode
@author: fke
@file: 常量.go
@time: 2019-05-28 19:23:29
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

package main

/*
func main() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str"   // 多重赋值

	area = LENGTH * WIDTH
	fmt.Printf("面积为 ：%d", area)
	println()
	println(a, b, c)
}
*/



/*
import "unsafe"
const (
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(a)
)

func main() {
	println(a, b, c)
}

 */


/*
import "fmt"
func main() {
	const (
		a = iota	// 0
		b			// 1
		c			// 2
		d = "ha"	// 独立值, iota += 1
		e			// "ha" iota += 1
		f = 100		// iota += 1
		g			// 100 iota += 1
		h = iota	// 7, 恢复计数
		i			// 8
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)
}

*/


/*
import "fmt"
const (
	i = 1 << iota
	j = 3 << iota
	k
	l
)

func main() {
	fmt.Println("i = ", i)
	fmt.Println("j = ", j)
	fmt.Println("k = ", k)
	fmt.Println("l = ", l)
}
*/


import "fmt"
const (
	a = 1
	b
	c
	d
)

func main() {
	fmt.Println(a)
	// b, c, d 没有初始化， 使用上一行(即a)的值
	fmt.Println(b)   // 输出1
	fmt.Println(c)   // 输出1
	fmt.Println(d)   // 输出1
}
