package	main

import "fmt"

/*func main() {
	var i int
	var f float64
	var b bool
	var s string
	k := 15
	fmt.Printf("%v, %v %v %q, %v\n", i, f, b, s, k)
}*/

var x, y int
var ( // 这种因式分解关键字的写法一般用于申明全局变量
	a int
	b bool
)


// 这种不带申明格式的只能在函数体中出现
// g, h := 123, "hello"

func main() {
	g, h := 123, "hello"
	println(x, y, a, b, g, h)

	_, numb, strs := numbers()	// 只获取函数返回值的后2个
	fmt.Println(numb, strs)
}

func numbers()(int, int, string) {
	c, d ,e := 1, 2, "string"
	return c, d, e
}

